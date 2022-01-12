package main

import (
	"flag"
	"log"

	"github.com/dexterorion/hubbl-workflow-sketch/internal"
	"go.temporal.io/sdk/client"
)

var (
	action string
	wfname string

	signalName  string
	signalValue string
	workflowID  string
	runID       string
)

func init() {
	flag.StringVar(&action, "action", "worker", "Action type: worker | trigger | signal")
	flag.StringVar(&wfname, "wfname", "task-assignment", "If -action=trigger should have valid wfname")

	// for signal purposes
	flag.StringVar(&signalName, "signalName", "DispatchAutomatableTaskCompletion", "Signal name to be sent")
	flag.StringVar(&signalValue, "signalValue", "Some value", "Signal value to be sent")
	flag.StringVar(&workflowID, "workflowID", "", "Workflow to signal")
	flag.StringVar(&runID, "runID", "", "Workflow run to signal")

	flag.Parse()
}

func main() {
	if action != "worker" && action != "trigger" && action != "signal" {
		log.Fatal("action must be `worker` or `trigger`")
	}

	if action == "trigger" && wfname == "" {
		log.Fatal("when -action=trigger some `wfname` should be sent")
	}

	if action == "signal" &&
		(signalName == "" || signalValue == "" || workflowID == "" || runID == "") {
		log.Fatal("wrong parameters for signaling")
	}

	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	switch action {
	case "worker":
		internal.StartWorker(c)
	case "trigger":
		internal.TriggerWorkflow(c, wfname)
	case "signal":
		internal.SignalWorkflow(c, workflowID, runID, signalName, signalValue)
	}

	defer c.Close()

}
