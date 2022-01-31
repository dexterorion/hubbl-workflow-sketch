package activities

import "go.temporal.io/sdk/worker"

func RegisterActivities(w worker.Worker) {
	// taskassignment
	w.RegisterActivity(FindUsersMatchingRoles)
	w.RegisterActivity(IsTaskAutomatable)
	w.RegisterActivity(CleanDatabase)
	w.RegisterActivity(FLushLogs)
	w.RegisterActivity(ReleaseConn)
	w.RegisterActivity(SendMessageToTopic)
	w.RegisterActivity(CreateNotice)
	w.RegisterActivity(NotifyTaskAssignment)

	// offerassignment
	w.RegisterActivity(EvalAutomation)
	w.RegisterActivity(DecidesAutomation)
	w.RegisterActivity(AutomateAutomation)
	w.RegisterActivity(EvaluateExternalSystem)
	w.RegisterActivity(DecideExternalSystem)
	w.RegisterActivity(NotifySuccess)
	w.RegisterActivity(NotifyFail)
	w.RegisterActivity(SendOfferToUsers)
	w.RegisterActivity(OwnerWillDoTheTask)
	w.RegisterActivity(OwnerAssignToSomeone)
}
