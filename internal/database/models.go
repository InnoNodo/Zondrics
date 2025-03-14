package database

type User struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Login   string `gorm:"unique"`
	Hash    string `gorm:"not null"`
	Name    string `gorm:"not null"`
	Surname string `gorm:"not null"`
	Phone   string `gorm:"not null;unique"`
}

type Admin struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`
}

type Organization struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
	City string `gorm:"not null"`
}

type OrganizationUser struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	UserID         uint         `gorm:"not null"`
	User           User         `gorm:"foreignKey:UserID"`
	OrganizationID uint         `gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:Organization"`
}

type RegistrationInput struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
