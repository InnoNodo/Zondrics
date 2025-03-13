package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"regexp"
)

var DB *gorm.DB

func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^\+\d{1,3}\d{10}$`)
	return re.MatchString(phone)
}

func InitDatabase() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
