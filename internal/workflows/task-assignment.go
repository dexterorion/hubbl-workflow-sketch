package workflows

import (
	"log"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func StartTaskAssignment(ctx workflow.Context, story models.Story, taskPlan models.TaskPlanTemplate, assignments []models.StoryAssignment) (err error) {
	log.Println("Starting StartTaskAssignment...")

	return
}
