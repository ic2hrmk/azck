package main

import (
	"log"

	"github.com/ic2hrmk/azck/app/services/consumer"
	"github.com/ic2hrmk/azck/app/services/producer"
	"github.com/ic2hrmk/azck/registry"
	"github.com/ic2hrmk/azck/shared/conf"
)

//go:generate go run entry.go --kind=producer --env=.env
//go:generate go run entry.go --kind=consumer --env=.env

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
		log.Fatal(err)
	}

	//
	// Create service
	//
	service, err := serviceFactory()
	if err != nil {
		log.Fatal(err)
	}

	//
	// Run till the death comes
	//
	log.Printf("service [%s] started!", flags.Kind)
	log.Fatal(service.Run())
}
