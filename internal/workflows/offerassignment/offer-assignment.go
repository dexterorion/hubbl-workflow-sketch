package offerassignment

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func StartOfferAssignmentWorkflow(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting StartOfferAssignmentWorkflow...")
	defer logger.Debug("Ending StartOfferAssignmentWorkflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	automationResponse := "succeded"
	if err = workflow.ExecuteChildWorkflow(ctx, AutomationWorkflow).Get(ctx, &automationResponse); err != nil {
		return
	}

	if automationResponse != "succeded" {
		// starts external system workflow
		externalWillDoTheJob := "yes"
		if err = workflow.ExecuteChildWorkflow(ctx, ExternalSystemWorkflow).Get(ctx, &externalWillDoTheJob); err != nil {
			return
		}

		if externalWillDoTheJob != "yes" {
			if err = workflow.ExecuteChildWorkflow(ctx, AssignmentWorkflow).Get(ctx, &externalWillDoTheJob); err != nil {
				return
			}
		}
	}

	if err = workflow.ExecuteChildWorkflow(ctx, CompletionWorkflow, automationResponse).Get(ctx, nil); err != nil {
		return
	}

	return
}
