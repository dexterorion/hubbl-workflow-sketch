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
)

func init() {
	flag.StringVar(&action, "action", "worker", "Action type: worker | trigger")
	flag.StringVar(&wfname, "wfname", "task-assignment", "If -action=trigger should have valid wfname")

	flag.Parse()
}

func main() {
	if action != "worker" && action != "trigger" {
		log.Fatal("action must be `worker` or `trigger`")
	}

	if action == "trigger" && wfname == "" {
		log.Fatal("when -action=trigger some `wfname` should be sent")
	}

	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	if action == "worker" {
		internal.StartWorker(c)
	} else {
		internal.TriggerWorkflow(c, wfname)
	}

	defer c.Close()

}
