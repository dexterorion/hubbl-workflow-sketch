package offerassignment

import "go.temporal.io/sdk/worker"

func RegisterWorkflows(w worker.Worker) {
	w.RegisterWorkflow(StartOfferAssignmentWorkflow)
	w.RegisterWorkflow(AssignmentWorkflow)
	w.RegisterWorkflow(CompletionWorkflow)
	w.RegisterWorkflow(ExternalSystemWorkflow)
}
