package workflows

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func NotifyAndWaitAcceptance(ctx workflow.Context, noticeSet *models.NoticeSet, storyAssignment *models.StoryAssignment) (notifications []*models.Notice, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting NotifyAndWaitExecution workflow...")
	defer logger.Debug("Finishing NotifyAndWaitExecution workflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	notifications = make([]*models.Notice, 0)

	wg := workflow.NewWaitGroup(ctx)

	for _, user := range storyAssignment.Users {

		notice := &models.Notice{}
		// create notice
		if err = workflow.ExecuteActivity(ctx, activities.CreateNotice, noticeSet, user, storyAssignment.AcceptPeriod).Get(ctx, notice); err != nil {
			return
		}

		notifications = append(notifications, notice)

		wg.Add(1)
		// async call
		workflow.Go(ctx, func(ctx workflow.Context) {
			wfID := workflow.GetInfo(ctx).WorkflowExecution.ID
			runID := workflow.GetInfo(ctx).WorkflowExecution.RunID
			// send message to user and wait for deadline
			if err = workflow.ExecuteActivity(ctx, activities.NotifyTaskAssignment, notice, wfID, runID).Get(ctx, nil); err != nil {
				return
			}

			// or wait for signal
			{
				// build signal
				s := workflow.NewSelector(ctx)

				var signalVal string
				acceptedChan := workflow.GetSignalChannel(ctx, DispatchUserAcceptedSignal)
				declinedChan := workflow.GetSignalChannel(ctx, DispatchUserRefusedSignal)

				s.AddReceive(acceptedChan, func(c workflow.ReceiveChannel, more bool) {
					c.Receive(ctx, &signalVal)

					logger.Debug("User accepted task...", "email", user.Email)
				})
				s.AddReceive(declinedChan, func(c workflow.ReceiveChannel, more bool) {
					c.Receive(ctx, &signalVal)

					logger.Debug("User refused task...", "email", user.Email)
				})
				s.Select(ctx) // waits here
			}
			// or wait for duration
			{
				now := time.Now()
				duration := notice.Deadline.Sub(now)

				<-time.After(duration)
			}
			wg.Done()
		})
	}

	// wait for signal or all above deadline are done
	wg.Wait(ctx)

	return
}
