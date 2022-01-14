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

	total := len(storyAssignment.Users)
	receivedDeny := 0

	s := workflow.NewSelector(ctx)

	/*

		workflow principal
		    - pra cada user, dispara um workflow filho

		    - espera por
		        para cada channel recebido
		            - se alguem aceitou:
		                - retorna usuario que aceitou
		            - se alguem na aceitou:
		                total de signais == ao maximo esperado?
		                    - sim: retorna que ninguem aceitou
		                    - nao: continua

		workflow filho
		    - envia mensagem pro usuario

		    - routine 1:
		        - espera por
		            sinal recebido de aceitacao?
		                - envia channel de aceitacao
		            sinal recebido de declinio?
		                - envia channel de declinio

		    - routine 2:
		        - espera por deadline
		            - envia channel de declinio

	*/

	var signalVal *models.Notice
	acceptedChan := workflow.GetSignalChannel(ctx, AcceptedSignal)
	declinedChan := workflow.GetSignalChannel(ctx, RefusedSignal)

	result := make(chan *models.Notice, 1)

	s.AddReceive(acceptedChan, func(c workflow.ReceiveChannel, more bool) {
		c.Receive(ctx, &signalVal)

		logger.Debug("User accepted task...", "email", signalVal.User.Email)

		// received ack, continues and kills children workflows
		result <- signalVal

	})

	s.AddReceive(declinedChan, func(c workflow.ReceiveChannel, more bool) {
		for {
			c.Receive(ctx, &signalVal)
			logger.Debug("User refused task...", "email", signalVal.User.Email)

			receivedDeny++
			if receivedDeny == total {
				// all waiting is done. can proceed to next stage
				result <- nil
			}
		}
	})

	for _, user := range storyAssignment.Users {

		notice := &models.Notice{}
		// create notice
		if err = workflow.ExecuteActivity(ctx, activities.CreateNotice, noticeSet, user, storyAssignment.AcceptPeriod).Get(ctx, notice); err != nil {
			return
		}

		notifications = append(notifications, notice)

		// async call
		workflow.Go(ctx, func(ctx workflow.Context) {
			if err = workflow.ExecuteChildWorkflow(ctx, WaitOrDeadline, notice, acceptedChan, declinedChan).Get(ctx, nil); err != nil {
				return
			}
		})
	}

	s.Select(ctx) // waits here

	return
}
