package model

import "time"

type Category struct {
	ID          int64      `gorm:"id"`
	Title       string     `gorm:"title"`
	Slug        string     `gorm:"slug"`
	CreatedByID int64      `gorm:"created_by_id"`
	User        User       `gorm:"foreignKey:CreatedByID;references:ID"` // Foreign key relationship to User
	CreatedAt   time.Time  `gorm:"created_at"`
	UpdatedAt   *time.Time `gorm:"updated_at"` // Pointer to time.Time to allow null values
}
