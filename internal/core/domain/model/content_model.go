package model

import "time"

type Content struct {
	ID          int64      `gorm:"id"`
	Title       string     `gorm:"title"`
	Excerpt     string     `gorm:"excerpt"`
	Description string     `gorm:"description"`
	Image       string     `gorm:"image"`
	Status      string     `gorm:"status"`
	Tags        string     `gorm:"tags"`
	CategoryID  int64      `gorm:"category_id"`
	Category    Category   `gorm:"foreignKey:CategoryID;references:ID"` // Foreign key relationship to Category
	CreatedByID int64      `gorm:"created_by_id"`
	User        User       `gorm:"foreignKey:CreatedByID;references:ID"` // Foreign key relationship to User
	CreatedAt   time.Time  `gorm:"created_at"`
	UpdatedAt   *time.Time `gorm:"updated_at"` // Pointer to time.Time to allow null values
}
