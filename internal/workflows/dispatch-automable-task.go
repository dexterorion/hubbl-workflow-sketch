package workflows

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

const (
	DispatchAutomatableSignal = "DispatchAutomatableTaskCompletion"
)

func DispatchAutomatableTask(ctx workflow.Context, storyAssignment *models.StoryAssignment) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting DispatchAutomatableTask workflow...")
	defer logger.Debug("Finishing DispatchAutomatableTask workflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// automated task will probably force us to wait for a signal
	// because it might be an external task which will signal the
	// flow when it is done

	internalCh := make(chan error)

	workflow.Go(ctx, func(ctx workflow.Context) {
		wfID := workflow.GetInfo(ctx).WorkflowExecution.ID
		runID := workflow.GetInfo(ctx).WorkflowExecution.RunID

		message := struct {
			WorkflowID string
			RunID      string
			Info       []interface{} // any info here
		}{
			WorkflowID: wfID,
			RunID:      runID,
			Info:       []interface{}{storyAssignment},
		}

		if err := workflow.ExecuteActivity(ctx, activities.SendMessageToTopic, "externaltopic", message).Get(ctx, nil); err != nil {
			// if it fails, signal with error
			internalCh <- err
		} else {
			internalCh <- nil
		}
	})

	// waits to see if workflow.Go worked fine
	err = <-internalCh
	if err != nil {
		return
	}

	// starts waiting for outside world
	s := workflow.NewSelector(ctx)

	var signalVal string
	signalChan := workflow.GetSignalChannel(ctx, DispatchAutomatableSignal)

	s.AddReceive(signalChan, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &signalVal)

		logger.Debug("Automated task done. Proceeding...")
	})
	s.Select(ctx) // waits here

	return
}
