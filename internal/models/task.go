package models

import "time"

type Task struct {
	Name string
}

type TaskAssignment struct {
	Task Task
}

type TaskPlanTemplate struct {
	OwnerRole        string
	Fields           []string
	CompletionPeriod time.Duration
	Stages           []TaskStage
}

type TaskStage struct {
	AcceptedPeriod time.Duration
	Roles          []Role
	Automatable    bool
}

type TaskPlan struct {
	CreatedAt        time.Time
	StoryAssignments []StoryAssignment
}
