package models

import "time"

type Resource struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	OrganizationID uint         `gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	Name           string       `gorm:"not null"`
	Type           string       `gorm:"type:varchar(50);not null"`
	Location       string       `gorm:"not null"`
	Capacity       uint         `gorm:"not null"`
	Status         string       `gorm:"type:varchar(20);default:'available'"`
	CreatedAt      time.Time    `gorm:"autoCreateTime"`
	UpdatedAt      time.Time    `gorm:"autoUpdateTime"`
}
