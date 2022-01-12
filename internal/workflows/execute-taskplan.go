package workflows

import (
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func ExecuteTaskPlan(ctx workflow.Context, taskPlan *models.TaskPlan) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting ExecuteTaskPlan workflow...")
	defer logger.Debug("Finishing ExecuteTaskPlan workflow...")

	for _, storyAssignment := range taskPlan.StoryAssignments {
		if storyAssignment.Automatable {
			if err = workflow.ExecuteChildWorkflow(ctx, DispatchAutomatableTask).Get(ctx, nil); err != nil {
				return
			}
		} else {
			// should be done by some user
		}
	}

	return
}
