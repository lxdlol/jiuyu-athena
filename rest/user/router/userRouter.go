package router

import (
	"athena/rest/user/handler"
	"github.com/gin-gonic/gin"
)

// rObject
type rObject struct{}

// NewRouter 初始化项目路由
func NewRouter() *rObject {
	return new(rObject)
}

// SetObjectRouter 设置项目路由
func (o *rObject) Set(r *gin.Engine, auth gin.HandlerFunc) {
	o.authRouter(r)
	o.userRouter(r)
}

func (o *rObject) authRouter(r *gin.Engine) {
	router := r.Group("/auth")
	{
		//登录
		router.POST("/login", handler.LoginHandler)
		//注册
		router.POST("/:namespace/register", handler.RegisterHandler)
		//忘记密码
		router.POST("/forget-password")
		//重设密码
		router.POST("/reset-password")
		//退出登录
		router.POST("/exit")
	}

}

func (o *rObject) userRouter(r *gin.Engine) {
	router := r.Group("/user")
	{
		//vip等级增删改查
		router.POST("/vip-Grade")
		router.PUT("/vip-Grade")
		router.DELETE("/vip-Grade")
		router.GET("/vip-Grade")
	}

}
