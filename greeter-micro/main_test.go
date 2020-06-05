package main

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	greeter "greeter-micro/proto/greeter"
	"testing"
	"time"
)

func BenchmarkServer(b *testing.B) {
	c := greeter.NewGreeterService("com.tutils.service.greeter", client.DefaultClient)
	const name = "t5w0rd"
	const assertRsp = "Hello t5w0rd"

	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Call(ctx, &greeter.Request{Name: name})
		if err != nil {
			b.Fatalf("could not greet: %v", err)
		}
		if r.GetMsg() != assertRsp {
			b.Error(r.GetMsg())
		}
	}
}
