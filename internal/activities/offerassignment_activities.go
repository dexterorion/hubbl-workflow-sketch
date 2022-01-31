package activities

import (
	"context"
	"math/rand"
	"time"

	"go.temporal.io/sdk/activity"
)

func EvalAutomation(ctx context.Context) (result bool, err error) {
	activity.GetLogger(ctx).Debug("Evaluation automation")

	// randomizing
	rand.Seed(int64(time.Now().UnixNano()))
	result = rand.Intn(2) == 1

	return
}

func DecidesAutomation(ctx context.Context) (result bool, err error) {
	activity.GetLogger(ctx).Debug("Decides automation")

	// randomizing
	rand.Seed(int64(time.Now().UnixNano()))
	result = rand.Intn(2) == 1

	return
}

func AutomateAutomation(ctx context.Context) (result bool, err error) {
	activity.GetLogger(ctx).Debug("Automate automation")

	// randomizing
	rand.Seed(int64(time.Now().UnixNano()))
	result = rand.Intn(2) == 1

	return
}

func EvaluateExternalSystem(ctx context.Context) (result bool, err error) {
	activity.GetLogger(ctx).Debug("EvaluateExternalSystem automation")

	// randomizing
	rand.Seed(int64(time.Now().UnixNano()))
	result = rand.Intn(2) == 1

	return
}

func DecideExternalSystem(ctx context.Context) (result bool, err error) {
	activity.GetLogger(ctx).Debug("DecideExternalSystem automation")

	// randomizing
	rand.Seed(int64(time.Now().UnixNano()))
	result = rand.Intn(2) == 1

	return
}

func NotifySuccess(ctx context.Context) (err error) {
	activity.GetLogger(ctx).Debug("Notify success")

	return
}

func NotifyFail(ctx context.Context) (err error) {
	activity.GetLogger(ctx).Debug("Notify fail")

	return
}

func SendOfferToUsers(ctx context.Context) (result bool, err error) {
	activity.GetLogger(ctx).Debug("Sending offer to users")

	// randomizing
	rand.Seed(int64(time.Now().UnixNano()))
	result = rand.Intn(2) == 1

	return
}

func OwnerWillDoTheTask(ctx context.Context) (result bool, err error) {
	activity.GetLogger(ctx).Debug("OwnerWillDoTheTask")

	// randomizing
	rand.Seed(int64(time.Now().UnixNano()))
	result = rand.Intn(2) == 1

	return
}

func OwnerAssignToSomeone(ctx context.Context) (err error) {
	activity.GetLogger(ctx).Debug("OwnerAssignToSomeone")

	return
}
