package workflows

import (
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func StartTaskAssignmentWorkflow(ctx workflow.Context, story *models.Story, template *models.TaskPlanTemplate, assignments []*models.StoryAssignment) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug(("Starting StartTaskAssignmentWorkflow..."))

	taskPlan := &models.TaskPlan{}
	if err = workflow.ExecuteChildWorkflow(ctx, ConstructTaskPlanWorkflow, template, assignments).Get(ctx, taskPlan); err != nil {
		return
	}

	logger.Debug("Ending StartTaskAssignmentWorkflow...")

	return
}
