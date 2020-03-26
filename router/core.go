package router

import "github.com/gin-gonic/gin"

// routerObj 路由对象
type routerObj interface {
	Set(r *gin.Engine, auth gin.HandlerFunc)
}
