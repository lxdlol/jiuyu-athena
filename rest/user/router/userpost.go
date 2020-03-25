package router

import (
	"athena/rest"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func swaggerRouter() {
	router := rest.InitRouter()
	router.GET("api/v1/auth/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func authRouter() {
	router := rest.InitRouter()
	auth := router.Group("/auth")
	//登录
	auth.POST("")
	//注册
	auth.POST("")
	//忘记密码
	auth.POST("")
	//重设密码
	auth.POST("")
	//退出登录
	auth.POST("")
}

func userRouter() {
	router := rest.InitRouter()
	user := router.Group("/user")
	user.POST("/vip")
	user.PUT("/vip")
	user.DELETE("/vip")
	user.GET("/vip")
}
