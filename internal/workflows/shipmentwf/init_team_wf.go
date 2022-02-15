package shipmentwf

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func InitTeamWf(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting InitTeamWf...")
	defer logger.Debug("Ending InitTeamWf...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	return
}
