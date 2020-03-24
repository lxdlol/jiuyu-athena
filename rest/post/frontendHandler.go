package post

import (
	"athena/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 发布
func AddPost(c *gin.Context) {
	var post models.Post
	c.Bind(post)
	err := models.InsertPost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 201,
		"msg":    "success",
		"data":   nil,
	})
}
