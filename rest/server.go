package rest

import (
	"athena/config"
	"github.com/gin-gonic/gin"
	"github.com/siddontang/go-log/log"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func (server *HttpServer) Start() {
	r := InitRouter()

	r.Use(setCROSOptions)
	err := r.Run(server.addr)
	if err != nil {
		panic(err)
	}
}

func setCROSOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")
}

func StartServer() {
	athenaConfig := config.GetConfig()

	httpServer := NewHttpServer(athenaConfig.RestServer.Addr)
	go httpServer.Start()

	log.Info("rest server ok")
}
