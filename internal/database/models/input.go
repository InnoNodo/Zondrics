package models

type UserRegistrationInput struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type OrganizationUserRegistrationInput struct {
	Login          string `json:"login"`
	OrganizationID uint   `json:"organization_id"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Password       string `json:"password"`
	Phone          string `json:"phone"`
	Type           string `json:"type"`
}

type OrganizationRegistrationInput struct {
	Name     string `json:"name"`
	City     string `json:"city"`
	Activity string `json:"activity"`
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
