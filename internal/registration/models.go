package registration

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Admin struct {
	User
	Permissions []string `json:"permissions"`
}

type OrganizationUser struct {
	User
	OrganizationID int `json:"organization_id"`
}
