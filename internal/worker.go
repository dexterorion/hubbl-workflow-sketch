package internal

import (
	"log"
	"os"
	"os/signal"

	"github.com/dexterorion/hubbl-workflow-sketch/internal/activities"
	"github.com/dexterorion/hubbl-workflow-sketch/internal/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func StartWorker(c client.Client) {
	w := worker.New(c, "task-assignment", worker.Options{
		EnableSessionWorker: true,
	})

	workflows.RegisterWorkflows(w)
	activities.RegisterActivities(w)

	log.Println("Starting worker...")
	err := w.Start()
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

	log.Println("Worker started...")

	startBlocking()
}

func startBlocking() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	<-signals
	log.Printf("received signal.. finishing process..")
}
