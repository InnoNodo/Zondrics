package database

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"regexp"
)

var DB *gorm.DB

func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^\+\d{1,3}\d{10}$`)
	return re.MatchString(phone)
}

func Hash(login, password string) string {
	secret := os.Getenv("SECRET")

	data := login + ":" + password + ":" + secret

	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
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
