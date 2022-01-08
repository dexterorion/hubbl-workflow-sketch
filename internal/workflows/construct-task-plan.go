package workflows

import (
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func ConstructTaskPlanWorkflow(ctx workflow.Context, template *models.TaskPlanTemplate, assignments []*models.StoryAssignment) (result *models.TaskPlan, err error) {
	result = &models.TaskPlan{}

	return
}
