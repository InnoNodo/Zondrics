package models

type Resource struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	OrganizationID uint         `gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	Name           string       `gorm:"not null"`
	Type           string       `gorm:"type:varchar(50);not null"`
	Location       string       `gorm:"not null"`
	Capacity       uint         `gorm:"not null"`
	Status         string       `gorm:"type:varchar(20);default:'available'"`
	CreatedAt      string       `gorm:"autoCreateTime"`
	UpdatedAt      string       `gorm:"autoUpdateTime"`
}

type Booking struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	UserID         uint         `gorm:"not null"`
	User           User         `gorm:"foreignKey:UserID"`
	OrganizationID uint         `gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	ResourceID     uint         `gorm:"not null"`
	Resource       Resource     `gorm:"foreignKey:ResourceID"`
	EventDate      string       `gorm:"not null"`
	EventTime      string       `gorm:"not null"`
	Duration       uint         `gorm:"not null"`
	Payment        PaymentInfo  `gorm:"embedded"`
	Status         string       `gorm:"type:varchar(20);default:'active'"`
	CreatedAt      string       `gorm:"autoCreateTime"`
}

type Training struct {
	ID              uint         `gorm:"primaryKey;autoIncrement"`
	OrganizationID  uint         `gorm:"not null"`
	Organization    Organization `gorm:"foreignKey:OrganizationID"`
	ResourceID      uint         `gorm:"not null"`
	Resource        Resource     `gorm:"foreignKey:ResourceID"`
	TrainerID       uint         `gorm:"not null"`
	Trainer         User         `gorm:"foreignKey:TrainerID"`
	Title           string       `gorm:"not null"`
	Description     string       `gorm:"type:text"`
	StartTime       string       `gorm:"not null"`
	Duration        uint         `gorm:"not null"`
	Payment         PaymentInfo  `gorm:"embedded"`
	MaxParticipants uint         `gorm:"not null"`
	CreatedAt       string       `gorm:"autoCreateTime"`
	UpdatedAt       string       `gorm:"autoUpdateTime"`
}

type PaymentInfo struct {
	Amount      float64 `gorm:"not null"`
	Currency    string  `gorm:"type:varchar(10);not null"`
	Status      string  `gorm:"type:varchar(20);default:'pending'"`
	PaymentDate string  `gorm:"type:datetime"`
}
