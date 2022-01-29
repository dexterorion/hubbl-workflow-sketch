package offerassignment

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func AssignmentWorkflow(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting AssignmentWorkflow...")
	defer logger.Debug("Ending AssignmentWorkflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	return
}
