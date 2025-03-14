package registration

import (
	"Zondrics/internal/database"
)

func NewUser(name, password, phone string) *database.User {
	return &database.User{
		Name:  name,
		Hash:  password,
		Phone: phone,
	}
}

func NewAdmin(name, password, phone string) *database.Admin {
	return &database.Admin{
		User: database.User{
			Name:  name,
			Hash:  password,
			Phone: phone,
		},
	}
}

func NewOrganizationUser(name, password, phone string, orgID uint) *database.OrganizationUser {
	return &database.OrganizationUser{
		User: database.User{
			Name: name,
			Hash: password,
		},
		OrganizationID: orgID,
	}
}
