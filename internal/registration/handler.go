package registration

import (
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	return c.SendString("You registered!")
}

func NewUser(username, password string) *User {
	return &User{
		Username: username,
		Password: password,
		Role:     "user",
	}
}

func NewAdmin(username, password string, permissions []string) *Admin {
	return &Admin{
		User: User{
			Username: username,
			Password: password,
			Role:     "admin",
		},
		Permissions: permissions,
	}
}

func NewOrganizationUser(username, password string, orgID int) *OrganizationUser {
	return &OrganizationUser{
		User: User{
			Username: username,
			Password: password,
			Role:     "organization_user",
		},
		OrganizationID: orgID,
	}
}
