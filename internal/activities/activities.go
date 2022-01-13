package activities

import (
	"context"
	"math/rand"
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"github.com/google/uuid"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/worker"
)

func RegisterActivities(w worker.Worker) {
	w.RegisterActivity(FindUsersMatchingRoles)
	w.RegisterActivity(IsTaskAutomatable)
	w.RegisterActivity(CleanDatabase)
	w.RegisterActivity(FLushLogs)
	w.RegisterActivity(ReleaseConn)
	w.RegisterActivity(SendMessageToTopic)
	w.RegisterActivity(CreateNotice)
	w.RegisterActivity(NotifyTaskAssignment)
}

func FindUsersMatchingRoles(ctx context.Context, roles []*models.Role) (users []*models.User, err error) {
	activity.GetLogger(ctx).Debug("Finding users matching roles", "roles", roles)

	users = []*models.User{
		{
			Email:    "fake01@hubbl.ai",
			Username: "fake01",
		},
		{
			Email:    "fake02@hubbl.ai",
			Username: "fake02",
		},
	}

	return
}

func IsTaskAutomatable(ctx context.Context, taskPlanId, taskStageId uuid.UUID) (automatable bool, err error) {
	activity.GetLogger(ctx).Debug("Checking if task is automatable by this stage", "taskPlanId", taskPlanId, "taskStageId", taskStageId)

	// randomizing
	rand.Seed(int64(time.Now().UnixNano()))
	automatable = rand.Intn(2) == 1

	if automatable {
		activity.GetLogger(ctx).Debug("Task automable", "taskPlanId", taskPlanId, "taskStageId", taskStageId)
	} else {
		activity.GetLogger(ctx).Debug("Non-automable", "taskPlanId", taskPlanId, "taskStageId", taskStageId)
	}

	return
}

func CleanDatabase(ctx context.Context) (err error) {
	// implement me
	return
}

func FLushLogs(ctx context.Context) (err error) {
	// implement me
	return
}

func ReleaseConn(ctx context.Context) (err error) {
	// implement me
	return
}

func SendMessageToTopic(ctx context.Context, topic string, message interface{}) (err error) {
	// some message queue or api call
	return
}

func CreateNotice(ctx context.Context, noticeSet *models.NoticeSet, user *models.User, acceptPeriod time.Duration) (notice *models.Notice, err error) {
	// it might do way more than that. call external api to check user's deadline

	deadline := time.Now().Add(acceptPeriod)

	notice = &models.Notice{
		User:      user,
		NoticeSet: noticeSet,
		Deadline:  deadline,
	}

	return
}

func NotifyTaskAssignment(ctx context.Context, notice *models.Notice, workflowID, runID string) (err error) {
	// send email, whatsapp, notification to phone etc

	return
}
