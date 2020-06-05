package main

import (
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"greeter-http/handler"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func nethttpServer() {
	h := &handler.Greeter{}
	mux := http.NewServeMux()
	mux.HandleFunc("/call", h.Call)
	s := &http.Server{Addr: ":50000", Handler: mux}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func fasthttpServer() {
	h := &handler.GreeterFastHttp{}
	m := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/call":
			h.Call(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}
	fasthttp.ListenAndServe(":50000", m)
}

func ginServer() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()
	h := &handler.GreeterGin{}
	r.POST("/call", h.Call)
	r.Run(":50000")
}

func main() {
	if len(os.Args) <= 1 {
		nethttpServer()
		return
	}

	switch os.Args[1] {
	case "nethttp":
		nethttpServer()
	case "fasthttp":
		fasthttpServer()
	case "gin":
		ginServer()
	}
}
