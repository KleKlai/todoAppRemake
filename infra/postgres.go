package infra

import (
	"errors"
	"fmt"

	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {

	var (
		dbUser     = "postgres"
		dbPassword = "postgres"
		dbHost     = "localhost"
		dbPort     = "5432"
		dbDatabase = "todo"
		dbSsl      = "disable"
	)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s", dbUser, dbPassword, dbHost, dbPort, dbDatabase, dbSsl, "Asia/Singapore")

	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.New("Failed to connect to database")
	}

	return db, nil
}
