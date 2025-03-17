package models

type RegistrationInput struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type BookingInput struct {
	OrganizationID     uint    `json:"organization_id" validate:"required"`
	ResourceID         uint    `json:"resource_id" validate:"required"`
	EventDate          string  `json:"event_date" validate:"required"`
	EventTime          string  `json:"event_time" validate:"required"`
	PaymentID          uint    `json:"payment_id" validate:"required"`
	Duration           int     `json:"duration"`
	Age                int     `json:"age"`
	Price              float64 `json:"price"`
	OrganizationUserID uint    `gorm:"not null"`
	TableNumber        int     `json:"table_number"`
	Guests             int     `json:"guests"`
}
