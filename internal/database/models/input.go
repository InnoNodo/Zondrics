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
	Name        string `json:"name"`
	City        string `json:"city"`
	Activity    string `json:"activity"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

type AuthInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type BookingInput struct {
	UserID             uint    `json:"user_id"`
	OrganizationID     uint    `json:"organization_id" validate:"required"`
	ResourceID         uint    `json:"resource_id" validate:"required"`
	EventTime          string  `json:"event_time" validate:"required"`
	Duration           int     `json:"duration"`
	Age                int     `json:"age"`
	Price              float64 `json:"price"`
	OrganizationUserID uint    `json:"organization_user_id" validate:"required"`
	TableNumber        int     `json:"table_number"`
	Guests             int     `json:"guests"`
	//PaymentID          uint    `json:"payment_id" validate:"required"`
}

type ResourceStatusInput struct {
	ResourceID uint   `json:"resource_id"`
	Status     string `json:"status"`
}

type EventInput struct {
	OrganizationID     uint   `json:"organization_id"`
	OrganizationUserID uint   `json:"organization_user_id"`
	Duration           int    `json:"duration"`
	ResourceID         uint   `json:"resource_id"`
	EventTime          string `json:"event_time"`
	Capacity           int    `json:"capacity"`
}

type ResourceInput struct {
	OrganizationID uint   `json:"organization_id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Location       string `json:"location"`
	Capacity       uint   `json:"capacity"`
	Status         string `json:"status"`
}

type CalendarInput struct {
	OrganizationID     uint   `json:"organization_id"`
	OrganizationUserID uint   `json:"organization_user_id"`
	Day                string `json:"day"`
}
