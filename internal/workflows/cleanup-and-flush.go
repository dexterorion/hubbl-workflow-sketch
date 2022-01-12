package workflows

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"go.temporal.io/sdk/workflow"
)

func CleanupAndFlush(ctx workflow.Context) (err error) {
	// this workflow is just for demonstration.
	// let's think by the end of the task assignment workflow, we need to cleanup and flush stuff
	// from outside componentes

	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting CleanupAndFlush workflow...")
	defer logger.Debug("Finishing CleanupAndFlush workflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	if err = workflow.ExecuteActivity(ctx, activities.CleanDatabase).Get(ctx, nil); err != nil {
		return
	}

	if err = workflow.ExecuteActivity(ctx, activities.FLushLogs).Get(ctx, nil); err != nil {
		return
	}

	if err = workflow.ExecuteActivity(ctx, activities.ReleaseConn).Get(ctx, nil); err != nil {
		return
	}

	return
}
