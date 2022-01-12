package models

import "time"

type Story struct {
	Name string
}

type StoryAssignment struct {
	Story        *Story
	Users        []*User
	Roles        []*Role
	Automatable  bool
	AcceptPeriod time.Duration
}
