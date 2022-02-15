package shipmentwf

import "go.temporal.io/sdk/worker"

func RegisterWorkflows(w worker.Worker) {
	w.RegisterWorkflow(RunShpWf)
	w.RegisterWorkflow(InitShpWf)
	w.RegisterWorkflow(InitTeamWf)
	w.RegisterWorkflow(ActivateStoryWf)
	w.RegisterWorkflow(SspApprovingWf)
	w.RegisterWorkflow(ValidateRequestWf)
	w.RegisterWorkflow(FinishShpWf)
	w.RegisterWorkflow(DeactivateTeamWf)
	w.RegisterWorkflow(InactivateHistoryWf)
}
