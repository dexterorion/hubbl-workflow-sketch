package taskassignment

import "go.temporal.io/sdk/worker"

func RegisterWorkflows(w worker.Worker) {
	w.RegisterWorkflow(StartTaskAssignmentWorkflow)
	w.RegisterWorkflow(ConstructTaskPlanWorkflow)
	w.RegisterWorkflow(CleanupAndFlush)
	w.RegisterWorkflow(DispatchAutomatableTask)
	w.RegisterWorkflow(ExecuteTaskPlan)
	w.RegisterWorkflow(NotifyAndWaitAcceptance)
	w.RegisterWorkflow(WaitOrDeadline)
}