package workflows

import (
	"github.com/dexterorion/hubbl-workflow-sketch/helpers"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func StartTaskAssignmentWorkflow(ctx workflow.Context, story *models.Story, template *models.TaskPlanTemplate) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug(("Starting StartTaskAssignmentWorkflow..."))

	taskPlan := &models.TaskPlan{}
	if err = workflow.ExecuteChildWorkflow(ctx, ConstructTaskPlanWorkflow, template).Get(ctx, taskPlan); err != nil {
		return
	}

	queue := helpers.StartEmptyQueue()

	logger.Debug("Ending StartTaskAssignmentWorkflow...", "queue", queue)

	return
}
