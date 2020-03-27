package rest

import (
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
		router.POST("/login")
		//注册
		router.POST("/register")
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
		router.POST("/vip")
		router.PUT("/vip")
		router.DELETE("/vip")
		router.GET("/vip")
	}

}
