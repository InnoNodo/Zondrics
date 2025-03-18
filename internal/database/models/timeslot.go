package models

import "time"

type Timeslot struct {
	ID                 uint             `gorm:"primary_key"`
	UserID             *uint            `gorm:"user_id"`
	User               User             `gorm:"foreignKey:UserID"`
	OrganizationID     uint             `gorm:"organization_id"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID *uint            `gorm:"organization_user_id"`
	OrganizationUser   OrganizationUser `gorm:"organization_user"`
	StartTime          time.Time        `gorm:"not null"`
	EndTime            time.Time        `gorm:"not null"`
}

type SlotStatus struct {
	StartTime time.Time
	EndTime   time.Time
	IsBooked  bool
}
