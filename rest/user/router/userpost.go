package router

import (
	"github.com/gin-gonic/gin"
)

// rObject
type rObject struct{}

// NewUserRouter 初始化项目路由
func NewUserRouter() *rObject {
	return new(rObject)
}

// SetObjectRouter 设置项目路由
func (o *rObject) Set(r *gin.Engine, auth gin.HandlerFunc) {
	o.authRouter(r)
	o.userRouter(r)
}

func (o *rObject) authRouter(r *gin.Engine) {
	auth := r.Group("/auth")
	//登录
	auth.POST("/login")
	//注册
	auth.POST("/register")
	//忘记密码
	auth.POST("/forget-password")
	//重设密码
	auth.POST("/reset-password")
	//退出登录
	auth.POST("/exit")
}

func (o *rObject) userRouter(r *gin.Engine) {
	user := r.Group("/user")
	//vip等级增删改查
	user.POST("/vip")
	user.PUT("/vip")
	user.DELETE("/vip")
	user.GET("/vip")
}
