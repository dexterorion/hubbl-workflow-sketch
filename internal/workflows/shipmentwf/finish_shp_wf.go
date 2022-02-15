package shipmentwf

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func FinishShpWf(ctx workflow.Context, endevent string) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting FinishShpWf...")
	defer logger.Debug("Ending FinishShpWf...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	if err = workflow.ExecuteChildWorkflow(ctx, DeactivateTeamWf).Get(ctx, nil); err != nil {
		return
	}

	if err = workflow.ExecuteChildWorkflow(ctx, InactivateHistoryWf).Get(ctx, nil); err != nil {
		return
	}

	if endevent == "CDC" {
		// means success

	} else {
		// fails, probably PBJ, BKJ
	}

	return
}
