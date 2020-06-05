package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	greeter "greeter-micro/proto/greeter"
)

type Greeter struct{}

func (e *Greeter) Handle(ctx context.Context, msg *greeter.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *greeter.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
