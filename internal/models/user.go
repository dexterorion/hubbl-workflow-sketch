package models

type User struct {
	Email    string
	Username string
	Roles    []*Role
}

type Role struct {
	Name        string
	Permissions []string
}
