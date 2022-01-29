package offerassignment

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"go.temporal.io/sdk/workflow"
)

func ExternalSystemWorkflow(ctx workflow.Context) (evalExtern string, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting ExternalSystemWorkflow...")
	defer logger.Debug("Ending ExternalSystemWorkflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	evalExtern = "yes"

	evaluateExternal := true
	if err = workflow.ExecuteActivity(ctx, activities.EvaluateExternalSystem).Get(ctx, &evaluateExternal); err != nil {
		evalExtern = "no"
		return
	}

	if evaluateExternal {
		decideExternal := true
		if err = workflow.ExecuteActivity(ctx, activities.DecideExternalSystem).Get(ctx, &decideExternal); err != nil {
			evalExtern = "no"
			return
		}

		if !decideExternal {
			evalExtern = "no"
		}
	}

	return
}
