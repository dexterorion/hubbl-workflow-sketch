package internal

import (
	"context"
	"log"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/workflows"
	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
)

func TriggerWorkflow(c client.Client, wfname string) {
	flow, found := workflows.GetStarterWorkflow(wfname)
	if !found {
		log.Fatalf("could not find workflow with name `%s`", wfname)
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:        wfname + uuid.New().String(),
		TaskQueue: wfname,
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, flow.Wf, flow.Parameters...)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
