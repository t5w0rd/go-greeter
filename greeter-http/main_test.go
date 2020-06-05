package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	greeter "greeter-http/proto/greeter"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func BenchmarkServer(b *testing.B) {
	c := &http.Client{}
	const url = "http://127.0.0.1:50000/call"
	const contentType = "application/json"
	const name = "t5w0rd"
	const assertRsp = "Hello t5w0rd"

	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.Reset()
		en := json.NewEncoder(buf)
		en.Encode(&greeter.Request{Name: name})
		r, err := c.Post(url, contentType, buf)
		if err != nil {
			b.Fatalf("could not greet: %v", err)
		}
		de := json.NewDecoder(r.Body)
		var rsp greeter.Response
		de.Decode(&rsp)
		if rsp.Msg != assertRsp {
			b.Error(rsp.Msg)
		}
	}
}

func BenchmarkRedisSet(b *testing.B) {
	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	buf := make([]byte, 10)
	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		rand.Read(buf)
		res := cli.Set(ctx, string(buf), buf, -1)
		cancel()
		if res.Err() != nil {
			b.Error(res.Err())
		}
	}
}
