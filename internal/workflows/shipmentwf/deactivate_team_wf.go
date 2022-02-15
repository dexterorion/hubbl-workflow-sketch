package shipmentwf

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func DeactivateTeamWf(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting DeactivateTeamWf...")
	defer logger.Debug("Ending DeactivateTeamWf...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	return
}
