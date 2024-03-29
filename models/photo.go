package models

import "time"

type Photo struct {
	Id        uint      `gorm:"primaryKey;type:bigint" json:"id"`
	Title     string    `gorm:"not null;type:varchar(100)" json:"title"`
	Caption   string    `gorm:"type:varchar(200)" json:"caption"`
	PhotoUrl  string    `gorm:"not null;type:text" json:"photo_url"`
	UserId    uint      `gorm:"not null;type:bigint" json:"user_id"`
	Comments  []Comment `gorm:"foreignKey:PhotoId" json:"-"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"-"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"-"`
}