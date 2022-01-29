package offerassignment

import (
	"errors"
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"go.temporal.io/sdk/workflow"
)

const (
	EzxecutedSignal = "ExecutedSignal"
)

func CompletionWorkflow(ctx workflow.Context, automationResponse string) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting CompletionWorkflow...")
	defer logger.Debug("Ending CompletionWorkflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	if automationResponse != "succeded" {
		// wait execution or run deadline

		wg := workflow.NewWaitGroup(ctx)
		wg.Add(1)

		workflow.Go(ctx, func(ctx workflow.Context) {
			s := workflow.NewSelector(ctx)
			executedChan := workflow.GetSignalChannel(ctx, EzxecutedSignal)

			s.AddReceive(executedChan, func(c workflow.ReceiveChannel, more bool) {
				logger.Debug("User executed")

				wg.Done()
			})

			s.Select(ctx)
		})

		workflow.Go(ctx, func(ctx workflow.Context) {
			{
				workflow.Sleep(ctx, 1*time.Hour) // waits needed time

				logger.Debug("User denied with timeout")
				err = errors.New("timeout")

				wg.Done()
			}
		})

		wg.Wait(ctx)

		if err != nil {
			if errNotify := workflow.ExecuteActivity(ctx, activities.NotifyFail).Get(ctx, nil); errNotify != nil {
				err = errNotify
			}
			return
		}
	}

	if err = workflow.ExecuteActivity(ctx, activities.NotifySuccess).Get(ctx, nil); err != nil {
		return
	}

	return
}
