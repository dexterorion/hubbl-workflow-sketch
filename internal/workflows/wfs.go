package workflows

import (
	"time"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"github.com/google/uuid"
	"go.temporal.io/sdk/worker"
)

var (
	starters = map[string]WfStartParameters{
		"task-assignment": {
			Wf: StartTaskAssignmentWorkflow,
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
	w.RegisterWorkflow(StartTaskAssignmentWorkflow)
	w.RegisterWorkflow(ConstructTaskPlanWorkflow)
	w.RegisterWorkflow(CleanupAndFlush)
	w.RegisterWorkflow(DispatchAutomatableTask)
	w.RegisterWorkflow(ExecuteTaskPlan)
	w.RegisterWorkflow(NotifyAndWaitAcceptance)
	w.RegisterWorkflow(WaitOrDeadline)
}
