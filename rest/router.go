package rest

import (
	"athena/config"
	"athena/rest/project/router"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() *gin.Engine {
	if Router == nil {
		Router = gin.Default()
	}

	// TODO: 依赖注入
	SetPostRouter(Router)
	router.SetObjectRouter(Router, authRouter())

	return Router
}

// authRouter 权限路由
func authRouter() gin.HandlerFunc {
	mode := config.GetConfig().RunMode
	if mode == "dev" {
		return gin.BasicAuth(gin.Accounts{
			"foo":    "bar",
			"austin": "1234",
			"lena":   "hello2",
			"manu":   "4321",
		})
	}

	return func(c *gin.Context) {
		c.Next()
	}
}
