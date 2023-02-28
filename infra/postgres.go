package infra

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {

	var err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var (
		dbUser     = "postgres"
		dbPassword = "postgres"
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbDatabase = "todo"
		dbSsl      = "disable"
	)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s", dbUser, dbPassword, dbHost, dbPort, dbDatabase, dbSsl, "Asia/Singapore")

	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})

	if err != nil {
		// return nil, errors.New("Failed to connect to database")
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	return db, nil
}
