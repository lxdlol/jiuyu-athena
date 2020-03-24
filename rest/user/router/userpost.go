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
	//登录
	router.POST("")
	//注册
	router.POST("")
	//忘记密码
	router.POST("")
	//重设密码
	router.POST("")
	//退出登录
	router.POST("")
}

func userRouter() {

}
