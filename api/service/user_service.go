package service

import (
	"endorseview/api/models"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	*models.User
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	log.Print("delete user")
	db = db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
