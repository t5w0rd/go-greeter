package main

import (
	"context"
	"google.golang.org/grpc"
	greeter "greeter-grpc/proto/greeter"
	"testing"
	"time"
)

func BenchmarkServer(b *testing.B) {
	conn, err := grpc.Dial("127.0.0.1:50000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		b.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := greeter.NewGreeterClient(conn)
	const name = "t5w0rd"
	const assertRsp = "Hello t5w0rd"

	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Call(ctx, &greeter.Request{Name: name})
		if err != nil {
			b.Fatalf("could not greet: %v", err)
		}
		if r.Msg != assertRsp {
			b.Error(r.Msg)
		}
	}
}
