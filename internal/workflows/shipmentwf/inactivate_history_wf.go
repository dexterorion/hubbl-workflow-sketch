package shipmentwf

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func InactivateHistoryWf(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting InactivateHistoryWf...")
	defer logger.Debug("Ending InactivateHistoryWf...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	return
}
