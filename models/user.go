package models

import (
	"final-project/helpers"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id              uint      `gorm:"primaryKey;type:bigint" json:"id"`
	Email           string    `gorm:"unique;not null;type:varchar(150)" json:"email"`
	Username        string    `gorm:"unique;not null;type:varchar(50)" json:"username"`
	Password        string    `gorm:"not null;type:text" json:"password"`
	Age             int       `gorm:"not null;type:int" json:"age"`
	ProfileImageUrl string    `gorm:"type:text" json:"profile_image_url"`
	CreatedAt       time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt       time.Time `gorm:"type:timestamp" json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// enkripsi password
	hash, err := helpers.HashPass(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	return
}