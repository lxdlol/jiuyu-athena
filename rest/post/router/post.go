package router

import (
	post2 "athena/rest/post"
	"github.com/gin-gonic/gin"
)

// rObject
type rObject struct{}

// NewObjectRouter 初始化项目路由
func NewPostRouter() *rObject {
	return new(rObject)
}

// SetObjectRouter 设置项目路由
func (o *rObject) Set(r *gin.Engine, auth gin.HandlerFunc) {
	post := r.Group("/post")
	post.POST("/", post2.AddPost)
	post.DELETE("/:id", func(context *gin.Context) {

	})
	post.PUT("/:id", func(context *gin.Context) {

	})
	post.GET("/:id", func(context *gin.Context) {

	})

	post.POST("/:id/comment", func(context *gin.Context) {

	})

	post.DELETE("/:id/comment/:commentId", func(context *gin.Context) {

	})

	post.POST("/:id/flower", func(context *gin.Context) {

	})
	post.POST("/:id/egg", func(context *gin.Context) {

	})

	//收藏，关注
	post.POST("/:id/favorite", func(context *gin.Context) {

	})
	post.DELETE("/:id/favorite", func(context *gin.Context) {

	})

}
