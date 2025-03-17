package database

import (
	"Zondrics/internal/database/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"regexp"
	"time"
)

var DB *gorm.DB
var secret = os.Getenv("SECRET")
var jwtSecret = os.Getenv("JWT_SECRET")

func ValidatePhone(phone string) bool {
	re := regexp.MustCompile(`^\+\d{1,3}\d{10}$`)
	return re.MatchString(phone)
}

func Hash(login, password string) string {

	data := login + ":" + password + ":" + secret

	hash := sha256.Sum256([]byte(data))

	return hex.EncodeToString(hash[:])
}

func CreateJWTToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"id":       user.ID,
		"login":    user.Login,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"issuedAt": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}

func InitDatabase() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
