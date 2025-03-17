package models

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
