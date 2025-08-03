package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github/Babe-piya/book-collection/appconfig"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection(cfg appconfig.Database) (*gorm.DB, *sql.DB, error) {
	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			cfg.Hostname,
			cfg.Username,
			cfg.Password,
			cfg.DatabaseName,
			cfg.Port,
			cfg.SSLMode,
			cfg.Timezone,
		)),
		&gorm.Config{
			Logger: sqlLogger,
		})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	return db, sqlDB, nil
}
