package registration

import (
	"Zondrics/internal/database/models"
)

func NewUser(name, password, phone string) *models.User {
	return &models.User{
		Name:  name,
		Hash:  password,
		Phone: phone,
	}
}

func NewAdmin(name, password, phone string) *models.Admin {
	return &models.Admin{
		User: models.User{
			Name:  name,
			Hash:  password,
			Phone: phone,
		},
	}
}

func NewOrganizationUser(name, password, phone string, orgID uint) *models.OrganizationUser {
	return &models.OrganizationUser{
		User: models.User{
			Name: name,
			Hash: password,
		},
		OrganizationID: orgID,
	}
}
