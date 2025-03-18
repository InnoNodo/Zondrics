package models

import "time"

type EventBooking struct {
	ID                 uint             `gorm:"primaryKey;autoIncrement"`
	UserID             uint             `gorm:"not null"`
	User               User             `gorm:"foreignKey:UserID"`
	OrganizationID     uint             `gorm:"not null"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	ResourceID         uint             `gorm:"not null"`
	Resource           Resource         `gorm:"foreignKey:ResourceID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	//PaymentID          uint             `gorm:"not null"`
	//Payment            Payment         `gorm:"foreignKey:PaymentID"`
	Duration  int       `gorm:"not null"`
	EventTime time.Time `json:"event_time" validate:"required"`
	Age       int       `gorm:""`
}

type HaircutBooking struct {
	ID                 uint             `gorm:"primaryKey;autoIncrement"`
	UserID             uint             `gorm:"not null"`
	User               User             `gorm:"foreignKey:UserID"`
	OrganizationID     uint             `gorm:"not null"`
	ResourceID         uint             `gorm:"not null"`
	Resource           Resource         `gorm:"foreignKey:ResourceID"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	EventTime          time.Time        `json:"event_time" validate:"required"`
	Price              float64          `gorm:"not null"`
	//PaymentID          uint             `gorm:"not null"`
	//Payment            Payment         `gorm:"foreignKey:PaymentID"`
}

type RestaurantBooking struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	UserID         uint         `gorm:"not null"`
	User           User         `gorm:"foreignKey:UserID"`
	OrganizationID uint         `gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	ResourceID     uint         `gorm:"not null"`
	Resource       Resource     `gorm:"foreignKey:ResourceID"`
	TableNumber    int          `gorm:"not null"`
	EventTime      time.Time    `json:"event_time" validate:"required"`
	Guests         int          `gorm:"not null"`
	//PaymentID          uint             `gorm:"not null"`
	//Payment            Payment         `gorm:"foreignKey:PaymentID"`
}

type TrainingBooking struct {
	ID                 uint             `gorm:"primaryKey;autoIncrement"`
	UserID             uint             `gorm:"not null"`
	User               User             `gorm:"foreignKey:UserID"`
	OrganizationID     uint             `gorm:"not null"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	ResourceID         uint             `gorm:"not null"`
	Resource           Resource         `gorm:"foreignKey:ResourceID"`
	EventTime          time.Time        `json:"event_time"`
	Age                int              `gorm:""`
	//PaymentID          uint             `gorm:"not null"`
	//Payment            Payment         `gorm:"foreignKey:PaymentID"`
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
	UserID      uint    `gorm:"primaryKey;autoIncrement"`
	User        User    `gorm:"foreignKey:UserID"`
	Amount      float64 `gorm:"not null"`
	Currency    string  `gorm:"type:varchar(10);not null"`
	Status      string  `gorm:"type:varchar(20);default:'pending'"`
	PaymentDate string  `gorm:"type:datetime"`
}
