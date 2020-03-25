package router

import (
	"athena/rest/project/handler"
	"github.com/gin-gonic/gin"
)

// SetObjectRouter 设置项目路由
func SetObjectRouter(r *gin.Engine, auth gin.HandlerFunc) {
	object := r.Group("object", auth)
	{
		object.GET("/index", handler.IndexGet)
	}

}
