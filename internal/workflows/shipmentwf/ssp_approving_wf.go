package shipmentwf

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func SspApprovingWf(ctx workflow.Context) (nextevt string, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting SspApprovingWf...")
	defer logger.Debug("Ending SspApprovingWf...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	if err = workflow.ExecuteChildWorkflow(ctx, ValidateRequestWf).Get(ctx, nil); err != nil {
		nextevt = "PBJ"
		return
	}

	nextevt = "PBA"

	return
}
