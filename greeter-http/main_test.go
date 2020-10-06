package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	greeter "greeter-http/proto/greeter"
	"log"
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
	b.ResetTimer()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		en := json.NewEncoder(buf)
		en.Encode(&greeter.Request{Name: name})

		b.StartTimer()
		r, err := c.Post(url, contentType, buf)
		b.StopTimer()

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

func BenchmarkTimeNow(b *testing.B) {
	var buf bytes.Buffer
	m := map[string]string {
		"a": "aaa",
		"b": "bbb",
	}
	buf.Grow(1000)
	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(m)
		buf.Write(data)
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

func TestHttpKeepAlive(t *testing.T) {
	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%#v", request.Header)
		writer.Write([]byte("OK\n"))
	})
	http.ListenAndServe(":28080", nil)
}
