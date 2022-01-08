package workflows

import (
	"github.com/dexterorion/hubbl-workflow-sketch/internal/models"
	"go.temporal.io/sdk/worker"
)

var (
	starters = map[string]WfStartParameters{
		"task-assignment": {
			Wf: StartTaskAssignment,
			Parameters: []interface{}{
				models.Story{}, models.TaskPlanTemplate{}, []models.StoryAssignment{},
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
	w.RegisterWorkflow(StartTaskAssignment)
}
