package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	TaskId uuid.UUID
	Name   string
}

type TaskAssignment struct {
	Task *Task
	User *User
}

type TaskPlanTemplate struct {
	TaskPlanId       uuid.UUID
	OwnerRole        string
	Fields           []string
	CompletionPeriod time.Duration
	Stages           []*TaskStage
}

type TaskStage struct {
	TaskStageId    uuid.UUID
	AcceptedPeriod time.Duration
	Roles          []*Role
	Automatable    bool
}

type TaskPlan struct {
	CreatedAt        time.Time
	StoryAssignments []*StoryAssignment
}
