package models

import "time"

type SocialMedia struct {
	Id             uint      `gorm:"primaryKey;type:bigint" json:"id"`
	Name           string    `gorm:"not null;type:varchar(50)" json:"name"`
	SocialMediaUrl string    `gorm:"not null;type:text" json:"social_media_url"`
	UserId         uint      `gorm:"not null;type:bigint" json:"user_id"`
	CreatedAt      time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp" json:"updated_at"`
}