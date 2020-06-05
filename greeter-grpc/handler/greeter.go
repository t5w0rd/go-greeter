package handler

import (
	"context"
	"log"

	greeter "greeter-grpc/proto/greeter"
)

type Greeter struct{}

// Call is a single request handler called via client.Call or the generated client code
func (s *Greeter) Call(ctx context.Context, req *greeter.Request) (*greeter.Response, error) {
	return &greeter.Response{Msg: "Hello " + req.GetName()}, nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (s *Greeter) Stream(req *greeter.StreamingRequest, stream greeter.Greeter_StreamServer) error {
	log.Printf("Received Greeter.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Printf("Responding: %d", i)
		if err := stream.Send(&greeter.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (s *Greeter) PingPong(stream greeter.Greeter_PingPongServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("Got ping %v", req.Stroke)
		if err := stream.Send(&greeter.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
