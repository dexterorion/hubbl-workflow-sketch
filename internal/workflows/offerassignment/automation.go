package offerassignment

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"go.temporal.io/sdk/workflow"
)

func AutomationWorkflow(ctx workflow.Context) (automationResponse string, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting AutomationWorkflow...")
	defer logger.Debug("Ending AutomationWorkflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	automationResponse = "succeded"

	canAutomate := true
	if err = workflow.ExecuteActivity(ctx, activities.EvalAutomation).Get(ctx, &canAutomate); err != nil {
		automationResponse = "failed"
		return
	}

	if !canAutomate {
		automationResponse = "no"
		return
	}

	willAutomate := true
	if err = workflow.ExecuteActivity(ctx, activities.DecidesAutomation).Get(ctx, &willAutomate); err != nil {
		automationResponse = "failed"
		return
	}

	if !willAutomate {
		automationResponse = "no"
		return
	}

	automated := true
	if err = workflow.ExecuteActivity(ctx, activities.AutomateAutomation).Get(ctx, &automated); err != nil {
		automationResponse = "failed"
		return
	}

	if !willAutomate {
		automationResponse = "failed"
		return
	}

	return
}
