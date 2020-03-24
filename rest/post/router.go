package post

import (
	"athena/rest"
	"github.com/gin-gonic/gin"
)

func PostRouter() {
	r := rest.InitRouter()
	post := r.Group("/post", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	post.POST("/", func(context *gin.Context) {

	})

	post.PUT("/:id", func(context *gin.Context) {

	})

	post.DELETE("/:id", func(context *gin.Context) {

	})
	post.PUT("/:id/comment", func(context *gin.Context) {

	})

	post.DELETE("/:id/comment/:commentId", func(context *gin.Context) {

	})

}
