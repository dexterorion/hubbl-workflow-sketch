package workflows

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func WaitOrDeadline(ctx workflow.Context, notice *models.Notice, parentWfID, parentRunID string) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting WaitOrDeadline workflow...", "user", notice.User.Email, "parentWfID", parentWfID, "parentRunID", parentRunID)
	defer logger.Debug("Finishing WaitOrDeadline workflow...", "user", notice.User.Email)

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	wg := workflow.NewWaitGroup(ctx)
	wg.Add(1)

	signalArg := &UserSignalNotice{
		Notice: notice,
	}

	workflow.Go(ctx, func(ctx workflow.Context) {
		s := workflow.NewSelector(ctx)
		acceptedChan := workflow.GetSignalChannel(ctx, DispatchUserAcceptedSignal)
		declinedChan := workflow.GetSignalChannel(ctx, DispatchUserRefusedSignal)

		s.AddReceive(acceptedChan, func(c workflow.ReceiveChannel, more bool) {
			logger.Debug("User accepted", "notice", notice)

			signalArg.Accepted = true
			err = workflow.SignalExternalWorkflow(ctx, parentWfID, parentRunID, UserNoticeSignal, signalArg).Get(ctx, nil)

			wg.Done()
		})

		s.AddReceive(declinedChan, func(c workflow.ReceiveChannel, more bool) {
			logger.Debug("User denied", "notice", notice)

			err = workflow.SignalExternalWorkflow(ctx, parentWfID, parentRunID, UserNoticeSignal, signalArg).Get(ctx, nil)

			wg.Done()
		})

		s.Select(ctx)
	})

	workflow.Go(ctx, func(ctx workflow.Context) {
		{
			now := time.Now()
			duration := notice.Deadline.Sub(now)

			workflow.Sleep(ctx, duration) // waits needed time

			logger.Debug("User denied with timeout", "notice", notice)

			err = workflow.SignalExternalWorkflow(ctx, parentWfID, parentRunID, UserNoticeSignal, signalArg).Get(ctx, nil)

			wg.Done()
		}
	})

	wfID := workflow.GetInfo(ctx).WorkflowExecution.ID
	runID := workflow.GetInfo(ctx).WorkflowExecution.RunID
	// send message to user and wait for deadline
	if err = workflow.ExecuteActivity(ctx, activities.NotifyTaskAssignment, notice, wfID, runID).Get(ctx, nil); err != nil {
		return
	}

	wg.Wait(ctx)

	return
}
