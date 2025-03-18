package models

import "time"

type HaircutBooking struct {
	ID                 uint             `gorm:"primaryKey;autoIncrement"`
	UserID             uint             `gorm:"not null"`
	OrganizationID     uint             `gorm:"not null"`
	ResourceID         uint             `gorm:"not null"`
	Resource           Resource         `gorm:"foreignKey:ResourceID"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	EventTime          time.Time        `json:"event_time" validate:"required"`
	Price              float64          `gorm:"not null"`
}

type RestaurantBooking struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	UserID         uint         `gorm:"not null"`
	OrganizationID uint         `gorm:"not null"`
	ResourceID     uint         `gorm:"not null"`
	Resource       Resource     `gorm:"foreignKey:ResourceID"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	TableNumber    int          `gorm:"not null"`
	EventTime      time.Time    `json:"event_time" validate:"required"`
	Guests         int          `gorm:"not null"`
}

type TrainingBooking struct {
	ID                 uint             `gorm:"primaryKey;autoIncrement"`
	UserID             uint             `gorm:"not null"`
	OrganizationID     uint             `gorm:"not null"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	ResourceID         uint             `gorm:"not null"`
	Resource           Resource         `gorm:"foreignKey:ResourceID"`
	Duration           int              `gorm:"not null"`
	EventTime          time.Time        `json:"event_time"`
	Age                int              `gorm:""`
}

type Event struct {
	OrganizationID     uint             `gorm:"primaryKey;autoIncrement"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	ResourceID         uint             `gorm:"not null"`
	Resource           Resource         `gorm:"foreignKey:ResourceID"`
	EventTime          time.Time        `gorm:"not null"`
	Capacity           int              `gorm:"not null"`
}

type Payment struct {
	Amount      float64 `gorm:"not null"`
	Currency    string  `gorm:"type:varchar(10);not null"`
	Status      string  `gorm:"type:varchar(20);default:'pending'"`
	PaymentDate string  `gorm:"type:datetime"`
}
