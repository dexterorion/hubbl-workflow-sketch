package workflows

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/workflows/offerassignment"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/workflows/shipmentwf"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/workflows/taskassignment"
	"github.com/google/uuid"
	"go.temporal.io/sdk/worker"
)

var (
	starters = map[string]WfStartParameters{
		"task-assignment": {
			Wf: taskassignment.StartTaskAssignmentWorkflow,
			Parameters: []interface{}{
				&models.Story{}, &models.TaskPlanTemplate{
					TaskPlanId:       uuid.New(),
					CompletionPeriod: 24 * time.Hour,
					Stages: []*models.TaskStage{
						{
							TaskStageId:    uuid.New(),
							Automatable:    true,
							AcceptedPeriod: 6 * time.Hour,
							Roles:          []*models.Role{},
						},
					},
				},
			},
		},
		"offer-assignment": {
			Wf:         offerassignment.StartOfferAssignmentWorkflow,
			Parameters: []interface{}{},
		},
		"shipmentwf": {
			Wf:         shipmentwf.RunShpWf,
			Parameters: []interface{}{},
		},
	}
)

type WfStartParameters struct {
	Wf         interface{}
	Parameters []interface{}
}

func GetStarterWorkflow(wfname string) (WfStartParameters, bool) {
	flow, ok := starters[wfname]
	return flow, ok
}

func RegisterWorkflows(w worker.Worker) {
	taskassignment.RegisterWorkflows(w)
	offerassignment.RegisterWorkflows(w)
	shipmentwf.RegisterWorkflows(w)
}
