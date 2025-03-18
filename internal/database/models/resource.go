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
