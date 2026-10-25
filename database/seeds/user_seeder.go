package seeds

import (
	"portalnews/internal/core/domain/model"
	"portalnews/lib/conv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating hash")
	}
	admin := model.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "admin@gmail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error user seeding")
	} else {
		log.Info().Msg("Admin user seeding successfully")
	}
}
