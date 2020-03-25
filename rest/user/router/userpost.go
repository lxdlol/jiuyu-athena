package router

import (
	"athena/rest"
)

func authRouter() {
	router := rest.InitRouter()
	auth := router.Group("/auth")
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

func userRouter() {
	router := rest.InitRouter()
	user := router.Group("/user")
	//vip等级增删改查
	user.POST("/vip")
	user.PUT("/vip")
	user.DELETE("/vip")
	user.GET("/vip")
}
