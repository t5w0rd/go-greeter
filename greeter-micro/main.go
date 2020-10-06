package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"greeter-micro/handler"
	greeter "greeter-micro/proto/greeter"
	"greeter-micro/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.tutils.service.greeter"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	greeter.RegisterGreeterHandler(service.Server(), new(handler.Greeter))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.tutils.service.greeter", service.Server(), new(subscriber.Greeter))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
