package workflows

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func WaitOrDeadline(ctx workflow.Context, notice *models.Notice, chAccept, chDeny workflow.Channel) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting WaitOrDeadline workflow...", "user", notice.User.Email)
	defer logger.Debug("Finishing WaitOrDeadline workflow...", "user", notice.User.Email)

	workflow.Go(ctx, func(ctx workflow.Context) {
		s := workflow.NewSelector(ctx)
		acceptedChan := workflow.GetSignalChannel(ctx, DispatchUserAcceptedSignal)
		declinedChan := workflow.GetSignalChannel(ctx, DispatchUserRefusedSignal)

		s.AddReceive(acceptedChan, func(c workflow.ReceiveChannel, more bool) {
			logger.Debug("User accepted", "notice", notice)

			chAccept.Send(ctx, notice)
		})

		s.AddReceive(declinedChan, func(c workflow.ReceiveChannel, more bool) {
			logger.Debug("User denied", "notice", notice)

			chDeny.Send(ctx, notice)
		})
	})

	workflow.Go(ctx, func(ctx workflow.Context) {
		{
			now := time.Now()
			duration := notice.Deadline.Sub(now)

			<-time.After(duration)

			chDeny.Send(ctx, notice)
		}
	})

	wfID := workflow.GetInfo(ctx).WorkflowExecution.ID
	runID := workflow.GetInfo(ctx).WorkflowExecution.RunID
	// send message to user and wait for deadline
	if err = workflow.ExecuteActivity(ctx, activities.NotifyTaskAssignment, notice, wfID, runID).Get(ctx, nil); err != nil {
		return
	}

	return
}
