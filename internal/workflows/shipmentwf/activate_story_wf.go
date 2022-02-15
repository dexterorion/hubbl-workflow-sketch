package shipmentwf

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func ActivateStoryWf(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting ActivateStoryWf...")
	defer logger.Debug("Ending ActivateStoryWf...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	return
}
