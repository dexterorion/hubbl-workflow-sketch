package shipmentwf

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func InitShpWf(ctx workflow.Context) (nextevt string, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting InitShpWf...")
	defer logger.Debug("Ending InitShpWf...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	if err = workflow.ExecuteChildWorkflow(ctx, ActivateStoryWf).Get(ctx, nil); err != nil {
		return
	}

	if err = workflow.ExecuteActivity(ctx, InitTeamWf).Get(ctx, nil); err != nil {
		return
	}

	nextevt = "PBQ"

	return
}
