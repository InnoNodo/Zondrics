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

type HaircutBooking struct {
	ID                 uint             `gorm:"primaryKey;autoIncrement"`
	UserID             uint             `gorm:"not null"`
	OrganizationID     uint             `gorm:"not null"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	EventDate          string           `json:"event_date" validate:"required"`
	EventTime          string           `json:"event_time" validate:"required"`
	Price              float64          `gorm:"not null"`
}

type RestaurantBooking struct {
	ID             uint         `gorm:"primaryKey;autoIncrement"`
	UserID         uint         `gorm:"not null"`
	OrganizationID uint         `gorm:"not null"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
	TableNumber    int          `gorm:"not null"`
	EventDate      string       `json:"event_date" validate:"required"`
	EventTime      string       `json:"event_time" validate:"required"`
	Guests         int          `gorm:"not null"`
}

type TrainingBooking struct {
	ID                 uint             `gorm:"primaryKey;autoIncrement"`
	UserID             uint             `gorm:"not null"`
	OrganizationID     uint             `gorm:"not null"`
	Organization       Organization     `gorm:"foreignKey:OrganizationID"`
	OrganizationUserID uint             `gorm:"not null"`
	OrganizationUser   OrganizationUser `gorm:"foreignKey:OrganizationUserID"`
	Duration           int              `gorm:"not null"`
	EventDate          string           `json:"event_date"`
	EventTime          string           `json:"event_time"`
	Age                int              `gorm:""`
}

type Payment struct {
	Amount      float64 `gorm:"not null"`
	Currency    string  `gorm:"type:varchar(10);not null"`
	Status      string  `gorm:"type:varchar(20);default:'pending'"`
	PaymentDate string  `gorm:"type:datetime"`
}
