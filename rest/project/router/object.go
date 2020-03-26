package router

import (
	"athena/rest/project/handler"
	"github.com/gin-gonic/gin"
)

// rObject
type rObject struct{}

// NewObjectRouter 初始化项目路由
func NewObjectRouter() *rObject {
	return new(rObject)
}

// SetObjectRouter 设置项目路由
func (o *rObject) Set(r *gin.Engine, auth gin.HandlerFunc) {
	object := r.Group("object", auth)
	{
		object.GET("/index", handler.IndexGet)
	}

}
