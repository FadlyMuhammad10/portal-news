package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rs/zerolog/log"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectPostgres() (*Postgres, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Psql.User, cfg.Psql.Password, cfg.Psql.Host, cfg.Psql.Port, cfg.Psql.DBName)

	db, err := gorm.Open(postgres.Open(dbConn), &gorm.Config{})

	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-1] failed to connect database" + cfg.Psql.Host)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-2] failed to get database connection" + cfg.Psql.Host)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.Psql.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.Psql.DBMaxIdle)

	return &Postgres{DB: db}, nil
}
