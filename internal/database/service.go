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

func CreateTokenForUser(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"role": "user",
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your-secret-key"))
}

func CreateTokenForOrganizationUser(user models.User, organizationID uint) (string, error) {
	claims := jwt.MapClaims{
		"sub":             user.ID,
		"role":            "organization_user",
		"organization_id": organizationID,
		"exp":             time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your-secret-key"))
}

func InitDatabase() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.Organization{},
		&models.OrganizationUser{},
		&models.TrainingBooking{},
		&models.HaircutBooking{},
		&models.RestaurantBooking{},
		&models.Resource{},
		&models.Payment{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
