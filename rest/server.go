package rest

import (
	"athena/config"
	_ "athena/docs"
	"athena/router"
	"github.com/siddontang/go-log/log"
)

// HttpServer 服务开启
type HttpServer struct {
	addr string
}

// NewHttpServer 初始化请求服务
func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

// Start 服务启动
func (server *HttpServer) Start() {
	r := router.InitRouter()
	err := r.Run(server.addr)
	if err != nil {
		panic(err)
	}
}

// StartServer 启动服务
func StartServer() {
	athenaConfig := config.GetConfig()

	httpServer := NewHttpServer(athenaConfig.RestServer.Addr)
	go httpServer.Start()

	log.Info("rest server ok")
}
