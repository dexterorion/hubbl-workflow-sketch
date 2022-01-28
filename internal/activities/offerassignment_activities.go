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
