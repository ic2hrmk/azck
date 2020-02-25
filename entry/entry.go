package main

import (
	"log"

	"github.com/ic2hrmk/azck/app/services/consumer"
	"github.com/ic2hrmk/azck/app/services/producer"
	"github.com/ic2hrmk/azck/registry"
	"github.com/ic2hrmk/azck/shared/conf"
	"github.com/ic2hrmk/azck/shared/runner"
)

//go:generate go run entry.go --kind=producer
//go:generate go run entry.go --kind=consumer

func main() {
	//
	// Load startup flags
	//
	flags := conf.LoadFlags()

	//
	// Select service
	//
	reg := registry.NewRegistryContainer()

	reg.Add(consumer.ServiceName, consumer.Factory)
	reg.Add(producer.ServiceName, producer.Factory)

	serviceFactory, err := reg.Get(flags.Kind)
	if err != nil {
		log.Fatalf("failed to select service mode to run, %s", err)
	}

	//
	// Create service
	//
	service, err := serviceFactory()
	if err != nil {
		log.Fatalf("failed to create service, %s", err)
	}

	//
	// Run till the death comes
	//
	log.Printf("service [%s] started!", flags.Kind)

	go func() { // Detach service main routine
		if err := service.Run(); err != nil {
			log.Fatalf("service stopped with critical error, %s", err)
		}
	}()

	runner.AwaitShutdown()

	if err = service.Stop(); err != nil {
		log.Printf("error ocurred durring service stop procedure, %s", err)
	}

	log.Println("service stopped!")
}
