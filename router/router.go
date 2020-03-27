package router

import (
	"athena/config"
	_ "athena/docs"
	post "athena/rest/post/router"
	object "athena/rest/project/router"
	"athena/rest/user/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gRouter 获取路由
type gRouter struct {
	*gin.Engine
	rSlice []interface{}
}

var (
	// r
	r *gRouter
)

func init() {
	r = &gRouter{
		Engine: gin.Default(),
		rSlice: []interface{}{
			object.NewRouter(),
			post.NewRouter(),
			router.NewRouter(),
		},
	}
}

func InitRouter() *gin.Engine {
	// 依赖注入 类继承检查
	if len(r.rSlice) > 0 {
		for _, v := range r.rSlice {
			if node, ok := v.(routerObj); ok {
				node.Set(r.Engine, r.basicAuth())
			}
		}
	}

	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))
	r.Use(setCROSOptions)
	return r.Engine
}

// setCROSOptions 设置跨域
func setCROSOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
	c.Header("Content-Type", "application/json")
}

// basicAuth 权限路由
func (r *gRouter) basicAuth() gin.HandlerFunc {
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
