package workflows

import (
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

const (
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

	notificationsMemory = []*models.Notice{}

	for _, storyAssignment := range taskPlan.StoryAssignments {
		if storyAssignment.Automatable {
			if err = workflow.ExecuteChildWorkflow(ctx, DispatchAutomatableTask).Get(ctx, nil); err != nil {
				return
			}
		} else {
			// should be done by some user
			notifications := []*models.Notice{}

			if err = workflow.ExecuteChildWorkflow(ctx, NotifyAndWaitAcceptance, noticeSet, storyAssignment).Get(ctx, &notifications); err != nil {
				return
			}

			notificationsMemory = append(notificationsMemory, notifications...)
		}
	}

	return
}
