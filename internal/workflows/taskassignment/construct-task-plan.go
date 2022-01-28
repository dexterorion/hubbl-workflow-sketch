package taskassignment

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/workflow"
)

func ConstructTaskPlanWorkflow(ctx workflow.Context, template *models.TaskPlanTemplate, story *models.Story) (result *models.TaskPlan, err error) {
	logger := workflow.GetLogger(ctx)
	logger.Debug("Starting ConstructTaskPlanWorkflow workflow...")
	defer logger.Debug("Finishing ConstructTaskPlanWorkflow workflow...")

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 1 * time.Hour,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	users := []*models.User{}

	result = &models.TaskPlan{
		StoryAssignments: make([]*models.StoryAssignment, 0),
		CreatedAt:        time.Now(),
	}

	for _, stage := range template.Stages {
		storyAssignment := &models.StoryAssignment{}

		if err = workflow.ExecuteActivity(ctx, activities.FindUsersMatchingRoles, stage.Roles).Get(ctx, &users); err != nil {
			logger.Error("Error or FindUsersMatchingRoles...")
			return
		}

		if err = workflow.ExecuteActivity(ctx, activities.IsTaskAutomatable, template.TaskPlanId, stage.TaskStageId).Get(ctx, &storyAssignment.Automatable); err != nil {
			logger.Error("Error or IsTaskAutomatable...")
			return
		}

		storyAssignment.AcceptPeriod = stage.AcceptedPeriod
		storyAssignment.Roles = stage.Roles
		storyAssignment.Story = story
		storyAssignment.Users = users

		result.StoryAssignments = append(result.StoryAssignments, storyAssignment)
	}

	// check possible(s) owner(s)

	ownerRole := []*models.Role{{Name: template.OwnerRole}}
	if err = workflow.ExecuteActivity(ctx, activities.FindUsersMatchingRoles, ownerRole).Get(ctx, &users); err != nil {
		logger.Error("Error or FindUsersMatchingRoles for owner...")
		return
	}

	result.StoryAssignments = append(result.StoryAssignments, &models.StoryAssignment{
		Story:        story,
		Users:        users,
		Roles:        ownerRole,
		Automatable:  false,
		AcceptPeriod: time.Hour * 1, // needs to be validated if there is any info to set here
	})

	return
}
