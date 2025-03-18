package models

type Event struct {
	OrganizationID     uint             `gorm:"primaryKey;autoIncrement"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	ResourceID         uint             `gorm:"not null"`
	Resource           Resource         `gorm:"foreignKey:ResourceID"`
	TimeSlotID         uint             `gorm:"not null"`
	Timeslot           Timeslot         `gorm:"foreignKey:TimeSlotID"`
	Capacity           int              `gorm:"not null"`
}
