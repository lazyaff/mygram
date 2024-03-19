package models

import "time"

type Comment struct {
	Id        uint      `gorm:"primaryKey;type:bigint" json:"id"`
	UserId    uint      `gorm:"not null;type:bigint" json:"user_id"`
	PhotoId   uint      `gorm:"not null;type:bigint" json:"photo_id"`
	Message   string    `gorm:"not null;type:varchar(200)" json:"message"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}