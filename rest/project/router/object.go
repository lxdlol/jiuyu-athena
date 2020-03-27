package router

import (
	"athena/rest/project/handler"
	"github.com/gin-gonic/gin"
)

// rObject
type rObject struct{}

// NewRouter 初始化项目路由
func NewRouter() *rObject {
	return new(rObject)
}

// Set 设置项目路由
func (o *rObject) Set(r *gin.Engine, auth gin.HandlerFunc) {
	router := r.Group("object", auth)
	{
		router.GET("/index", handler.IndexGet)
	}

}
