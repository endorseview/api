package models

type User struct {
	Base      BaseEntity
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Firstname string `gorm:"size:255;not null;unique" json:"firstname"`
	Lastname  string `gorm:"size:255;not null;unique" json:"lastname"`
	Nickname  string `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
}
