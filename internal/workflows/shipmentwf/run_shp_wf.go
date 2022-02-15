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

	currentState := "init"
	stateHistory := []string{currentState}

	eventHistory := []string{}
	currentEvent := ""

	if err = workflow.ExecuteChildWorkflow(ctx, ActivateStoryWf).Get(ctx, nil); err != nil {
		return
	}

	if err = workflow.ExecuteActivity(ctx, InitTeamWf).Get(ctx, nil); err != nil {
		return
	}

	currentState = "ssp.approving"
	stateHistory = append(stateHistory, currentState)

	currentEvent = "PBQ"
	eventHistory = append(eventHistory, currentEvent)

	if currentEvent != "PBQ" {
		err = fmt.Errorf("something went wrong running workflow. expected PBQ event")
		return
	}

	if err = workflow.ExecuteChildWorkflow(ctx, ValidateRequestWf).Get(ctx, nil); err != nil {
		currentEvent = "PBJ"
		currentState = "fail"

	} else {
		currentEvent = "PBA"
		currentState = "carrier.booking.send"
	}

	stateHistory = append(stateHistory, currentState)
	eventHistory = append(eventHistory, currentEvent)

	if currentEvent == "PBJ" {
		if err = workflow.ExecuteChildWorkflow(ctx, FinishShpWf, currentEvent).Get(ctx, nil); err != nil {
			return
		}
	}

	if err = workflow.ExecuteChildWorkflow(ctx, SendBookingWf).Get(ctx, nil); err != nil {
		currentEvent = "BKJ"
		currentState = "fail"

		stateHistory = append(stateHistory, currentState)
		eventHistory = append(eventHistory, currentEvent)
	} else {
		currentEvent = "BKQ"
		currentState = "carrier.booking.sent"

		stateHistory = append(stateHistory, currentState)
		eventHistory = append(eventHistory, currentEvent)

		currentEvent = "BKA"
		currentState = "carrier.booking.approved"

		stateHistory = append(stateHistory, currentState)
		eventHistory = append(eventHistory, currentEvent)
	}

	if currentEvent == "BKJ" {
		if err = workflow.ExecuteChildWorkflow(ctx, FinishShpWf, currentEvent).Get(ctx, nil); err != nil {
			return
		}
	}

	return
}
