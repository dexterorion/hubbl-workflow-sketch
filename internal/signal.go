package internal

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func SignalWorkflow(c client.Client, workflowID, runID, signalName, signalValue string) {
	err := c.SignalWorkflow(context.Background(), workflowID, runID, signalName, signalValue)

	if err != nil {
		log.Fatalln("Error signaling workflow")
	}

	log.Printf("Workflow %s signaled", workflowID)
}
