package shipmentwf

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func RunShpWf(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting StartShpWf...")
	defer logger.Debug("Ending StartShpWf...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	initShpWfEvent := ""
	if err = workflow.ExecuteChildWorkflow(ctx, InitShpWf).Get(ctx, &initShpWfEvent); err != nil {
		return
	}

	if initShpWfEvent != "PBQ" {
		err = fmt.Errorf("something went wrong running workflow. expected PBQ event")
		return
	}

	sspApprovingWfEvent := ""
	if err = workflow.ExecuteChildWorkflow(ctx, SspApprovingWf).Get(ctx, &sspApprovingWfEvent); err != nil {
		return
	}

	if sspApprovingWfEvent == "PBJ" {
		if err = workflow.ExecuteChildWorkflow(ctx, FinishShpWf, sspApprovingWfEvent).Get(ctx, nil); err != nil {
			return
		}
	}

	return
}
