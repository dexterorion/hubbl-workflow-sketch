package models

type User struct {
	Email    string
	Username string
}

type Role struct {
	Name        string
	Permissions []string
}
