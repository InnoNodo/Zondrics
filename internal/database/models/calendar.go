package models

import "time"

type Calendar struct {
	ID                 uint             `gorm:"primary_key"`
	OrganizationID     uint             `gorm:"not null"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	Day                time.Time        `gorm:"not null"`
}
