package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	greeter "greeter-http/proto/greeter"
	"net/http"
)

type Greeter struct{}

func (s *Greeter) Call(w http.ResponseWriter, r *http.Request) {
	de := json.NewDecoder(r.Body)
	var req greeter.Request
	de.Decode(&req)
	en := json.NewEncoder(w)
	en.Encode(&greeter.Response{Msg: "Hello " + req.Name})
}

type GreeterFastHttp struct{}

func (s *GreeterFastHttp) Call(ctx *fasthttp.RequestCtx) {
	var req greeter.Request
	json.Unmarshal(ctx.PostBody(), &req)
	bs, _ := json.Marshal(&greeter.Response{Msg: "Hello " + req.Name})
	ctx.Write(bs)
}

type GreeterGin struct{}

func (s *GreeterGin) Call(c *gin.Context) {
	var req greeter.Request
	c.ShouldBindJSON(&req)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello " + req.Name,
	})
}
