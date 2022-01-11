package workflows

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func ConstructTaskPlanWorkflow(ctx workflow.Context, template *models.TaskPlanTemplate) (result *models.TaskPlan, err error) {
	result = &models.TaskPlan{
		StoryAssignments: make([]*models.StoryAssignment, 0),
		CreatedAt:        time.Now(),
	}

	for _, stage := range template.Stages {
		users := []*models.User{}

		if err = workflow.ExecuteActivity(ctx, activities.FindUsersMatchingRoles, stage.Roles).Get(ctx, &users); err != nil {
			return
		}

	}

	return
}
