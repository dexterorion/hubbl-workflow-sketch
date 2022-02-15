package shipmentwf

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func ValidateRequestWf(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting ValidateRequestWf...")
	defer logger.Debug("Ending ValidateRequestWf...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	return
}
