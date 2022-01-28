package taskassignment

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

type WaitAcceptanceResult struct {
	Acceptance             *models.Notice
	GeneratedNotifications []*models.Notice
}

type UserSignalNotice struct {
	Accepted bool
	Notice   *models.Notice
}

func NotifyAndWaitAcceptance(ctx workflow.Context, noticeSet *models.NoticeSet, storyAssignment *models.StoryAssignment) (result *WaitAcceptanceResult, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting NotifyAndWaitExecution workflow...")
	defer logger.Debug("Finishing NotifyAndWaitExecution workflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	result = &WaitAcceptanceResult{}
	result.GeneratedNotifications = make([]*models.Notice, 0)

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

	parentWfID := workflow.GetInfo(ctx).WorkflowExecution.ID
	parentRunID := workflow.GetInfo(ctx).WorkflowExecution.RunID

	for _, user := range storyAssignment.Users {

		notice := &models.Notice{}
		// create notice
		if err = workflow.ExecuteActivity(ctx, activities.CreateNotice, noticeSet, user, storyAssignment.AcceptPeriod).Get(ctx, notice); err != nil {
			return
		}

		result.GeneratedNotifications = append(result.GeneratedNotifications, notice)

		// async call
		workflow.Go(ctx, func(ctx workflow.Context) {
			logger.Debug("Starting golang routine for WaitOrDeadline", "notice", notice)
			if err = workflow.ExecuteChildWorkflow(ctx, WaitOrDeadline, notice, parentWfID, parentRunID).Get(ctx, nil); err != nil {
				return
			}
		})
	}

	var signalVal *UserSignalNotice
	userNoticeChan := workflow.GetSignalChannel(ctx, UserNoticeSignal)

	s.AddReceive(userNoticeChan, func(c workflow.ReceiveChannel, more bool) {
		for {
			c.Receive(ctx, &signalVal)
			logger.Debug("User updated task notice...", "email", signalVal.Notice.User.Email)

			if !signalVal.Accepted {
				receivedDeny++
				if receivedDeny == total {
					// all waiting is done. can proceed to next stage
					return
				}
			} else {
				result.Acceptance = signalVal.Notice
				return
			}

		}
	})

	s.Select(ctx) // waits here

	return
}
