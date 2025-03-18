package models

import "time"

type Timeslot struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `gorm:"user_id"`
	User      User      `gorm:"foreignKey:UserID"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
}

type SlotStatus struct {
	StartTime time.Time
	EndTime   time.Time
	IsBooked  bool
}
