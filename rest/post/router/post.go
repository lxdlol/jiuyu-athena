package router

import (
	post2 "athena/rest/post"
	"github.com/gin-gonic/gin"
)

// rObject
type rObject struct{}

// NewRouter 初始化项目路由
func NewRouter() *rObject {
	return new(rObject)
}

// Set 设置项目路由
func (o *rObject) Set(r *gin.Engine, auth gin.HandlerFunc) {
	router := r.Group("/post")
	{
		router.POST("/", post2.AddPost)
		router.DELETE("/:id", func(context *gin.Context) {})
		router.PUT("/:id", func(context *gin.Context) {})         //修改一篇文章
		router.GET("/:id", func(context *gin.Context) {})         //查看一篇文章详情
		router.POST("/:id/flower", func(context *gin.Context) {}) //点赞一篇文章
		router.POST("/:id/egg", func(context *gin.Context) {})    //丢鸡蛋一篇文章
		//收藏，关注
		router.POST("/:id/favorite", func(context *gin.Context) {})
		router.DELETE("/:id/favorite", func(context *gin.Context) {})

		//评论
		router.POST("/:id/comment", func(context *gin.Context) {})
		router.GET("/:id/comment", func(context *gin.Context) {})
		router.DELETE("/:id/comment/:commentId", func(context *gin.Context) {})
		router.POST("/:id/comment/:commentId/flower", func(context *gin.Context) {}) //点赞
		router.POST("/:id/comment:commentId/egg", func(context *gin.Context) {})     //丢鸡蛋
	}

}
