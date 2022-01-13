package workflows

import (
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func StartTaskAssignmentWorkflow(ctx workflow.Context, story *models.Story, template *models.TaskPlanTemplate) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting StartTaskAssignmentWorkflow...")
	defer logger.Debug("Ending StartTaskAssignmentWorkflow...")

	// first we create the task plan, associating users/roles with stages
	// we create all story assignments. next workflow will handle all the processes
	// of create notify and so on
	taskPlan := &models.TaskPlan{}
	if err = workflow.ExecuteChildWorkflow(ctx, ConstructTaskPlanWorkflow, template, story).Get(ctx, taskPlan); err != nil {
		return
	}

	// most probably this child workflow will do the heavylifting
	// it should assign story, or trigger the automation and wait
	// it will also handle notifications
	// i am doing this in order to make more readable and state safer

	notificationsSent := []*models.Notice{}
	if err = workflow.ExecuteChildWorkflow(ctx, ExecuteTaskPlan, taskPlan).Get(ctx, &notificationsSent); err != nil {
		return
	}

	// flusing and releasing stuff
	if err = workflow.ExecuteChildWorkflow(ctx, CleanupAndFlush).Get(ctx, nil); err != nil {
		return
	}

	return
}
