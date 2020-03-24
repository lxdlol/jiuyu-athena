package rest

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() *gin.Engine {
	if Router == nil {
		Router = gin.Default()
	}
	return Router
}
