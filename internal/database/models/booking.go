package models

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
	TimeSlotID         uint             `gorm:"not null"`
	Timeslot           Timeslot         `gorm:"foreignKey:TimeSlotID"`
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
	TimeSlotID     uint         `gorm:"not null"`
	Timeslot       Timeslot     `gorm:"foreignKey:TimeSlotID"`
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
	TimeSlotID         uint             `gorm:"not null"`
	Timeslot           Timeslot         `gorm:"foreignKey:TimeSlotID"`
	Age                int              `gorm:""`
	//PaymentID          uint             `gorm:"not null"`
	//Payment            Payment         `gorm:"foreignKey:PaymentID"`
}

type EventBooking struct {
	ID      uint  `gorm:"primaryKey;autoIncrement"`
	UserID  uint  `gorm:"not null"`
	User    User  `gorm:"foreignKey:UserID"`
	EventID uint  `gorm:"not null"`
	Event   Event `gorm:"foreignKey:EventID"`
	//PaymentID          uint             `gorm:"not null"`
	//Payment            Payment         `gorm:"foreignKey:PaymentID"`
}

type Payment struct {
	UserID      uint    `gorm:"primaryKey;autoIncrement"`
	User        User    `gorm:"foreignKey:UserID"`
	Amount      float64 `gorm:"not null"`
	Currency    string  `gorm:"type:varchar(10);not null"`
	Status      string  `gorm:"type:varchar(20);default:'pending'"`
	PaymentDate string  `gorm:"type:datetime"`
}
