package service

import (
	"endorseview/api/models"
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"log"
	"strings"
	"time"
)

func DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	log.Print("delete user")
	db = db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BeforeSave(u models.User) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func Prepare(u models.User) {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Base.CreatedAt = time.Now()
	u.Base.UpdatedAt = time.Now()

}

func Validate(u models.User, action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func SaveUser(u models.User, db *gorm.DB) (*models.User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &models.User{}, err
	}
	return &u, nil
}

func FindAllUsers(db *gorm.DB) (*[]models.User, error) {
	var err error
	users := []models.User{}
	err = db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
	}
	return &users, err
}

func FindUserByID(u models.User, db *gorm.DB, uid uint32) (*models.User, error) {
	var err error
	err = db.Debug().Model(models.User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &models.User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.User{}, errors.New("User Not Found")
	}
	return &u, err
}

func UpdateAUser(u models.User, db *gorm.DB, uid uint32) (*models.User, error) {

	// To hash the password
	err := BeforeSave(u)
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &models.User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &models.User{}, err
	}
	return &u, nil
}
