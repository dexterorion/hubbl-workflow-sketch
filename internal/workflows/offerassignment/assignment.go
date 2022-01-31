package offerassignment

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"go.temporal.io/sdk/workflow"
)

const (
	AcceptOffer  = "AcceptOffer"
	DeclineOffer = "DeclineOffer"
)

func AssignmentWorkflow(ctx workflow.Context) (err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting AssignmentWorkflow...")
	defer logger.Debug("Ending AssignmentWorkflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	for {
		var offersSent bool
		var allDeclined bool
		var timedOut bool

		// send offer
		if err = workflow.ExecuteActivity(ctx, activities.SendOfferToUsers).Get(ctx, &offersSent); err != nil {
			return
		}

		if !offersSent {
			// indicates whether there are no users to receive offers

			ownerWillDo := true
			err = workflow.ExecuteActivity(ctx, activities.OwnerWillDoTheTask).Get(ctx, &ownerWillDo)

			if err == nil && !ownerWillDo {
				err = workflow.ExecuteActivity(ctx, activities.OwnerAssignToSomeone).Get(ctx, nil)
			}

			return
		}

		// wait for response or timeout
		wg := workflow.NewWaitGroup(ctx)
		wg.Add(1)

		workflow.Go(ctx, func(ctx workflow.Context) {
			s := workflow.NewSelector(ctx)
			acceptChan := workflow.GetSignalChannel(ctx, AcceptOffer)
			declineChan := workflow.GetSignalChannel(ctx, DeclineOffer)

			s.AddReceive(acceptChan, func(c workflow.ReceiveChannel, more bool) {
				logger.Debug("User accepted")

				wg.Done()
			})

			s.AddReceive(declineChan, func(c workflow.ReceiveChannel, more bool) {
				logger.Debug("User declined")

				// we need to do some check to see how many users received the offer
				// and if all of them decline, we continue
				// as for now we are just keeping it simple

				allDeclined = true

				wg.Done()
			})

			s.Select(ctx)
		})

		workflow.Go(ctx, func(ctx workflow.Context) {
			{
				workflow.Sleep(ctx, 1*time.Hour) // waits needed time

				logger.Debug("User denied with timeout")
				timedOut = true

				wg.Done()
			}
		})

		wg.Wait(ctx)

		// timed out
		if timedOut || allDeclined {
			continue
		}

		// someone has accepted
		return
	}
}
