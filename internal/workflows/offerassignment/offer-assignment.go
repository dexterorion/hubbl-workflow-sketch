package offerassignment

import "go.temporal.io/sdk/workflow"

func StartOfferAssignmentWorkflow(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting StartOfferAssignmentWorkflow...")
	defer logger.Debug("Ending StartOfferAssignmentWorkflow...")

	automationResponse := "succeded"
	if err = workflow.ExecuteChildWorkflow(ctx, AutomationWorkflow).Get(ctx, &automationResponse); err != nil {
		return
	}

	if automationResponse != "succeded" {
		// starts external system workflow
		if err = workflow.ExecuteChildWorkflow(ctx, ExternalSystemWorkflow).Get(ctx, &automationResponse); err != nil {
			return
		}
	} else {
		logger.Debug("OfferAssignmentWorkflow was finished succeded...")
	}

	return
}
