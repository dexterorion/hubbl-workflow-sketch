package models

type Story struct {
	Name string
}

type StoryAssignment struct {
	Story *Story
	User  []*User
	Roles []*Role
}
