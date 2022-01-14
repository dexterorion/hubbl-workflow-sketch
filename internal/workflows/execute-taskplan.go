package workflows

import (
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

const (
	UserNoticeSignal = "UserNoticeSignal"

	DispatchUserAcceptedSignal      = "DispatchUserAcceptedSignal"
	DispatchUserRefusedSignal       = "DispatchUserRefusedSignal"
	DispatchUserCompletedTaskSignal = "DispatchUserCompletedTaskSignal"
)

func ExecuteTaskPlan(ctx workflow.Context, taskPlan *models.TaskPlan) (notificationsMemory []*models.Notice, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting ExecuteTaskPlan workflow...")
	defer logger.Debug("Finishing ExecuteTaskPlan workflow...")

	noticeSet := &models.NoticeSet{
		Actions: []*models.Action{}, // bla bla a bunch of actions
	}

	waitacceptance := &WaitAcceptanceResult{}
	notificationsMemory = []*models.Notice{}

	for _, storyAssignment := range taskPlan.StoryAssignments {
		if storyAssignment.Automatable {
			err = workflow.ExecuteChildWorkflow(ctx, DispatchAutomatableTask).Get(ctx, nil)
			return
		} else {
			// should be done by some user

			if err = workflow.ExecuteChildWorkflow(ctx, NotifyAndWaitAcceptance, noticeSet, storyAssignment).Get(ctx, waitacceptance); err != nil {
				return
			}

			notificationsMemory = append(notificationsMemory, waitacceptance.GeneratedNotifications...)
		}
	}

	return
}
