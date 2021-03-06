package post

import (
	"athena/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIError struct {
	ErrorCode    int
	ErrorMessage string
}

// @Summary 添加文章
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param some_id path int true "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /post [post]
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

// @Summary 文章分页列表
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /post [get]
func ListPost(c *gin.Context) {
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

// @Summary 获取文章详情
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /post/{id} [get]
func GetPost(c *gin.Context) {
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

// @Summary 增加评论
// @Description 给文章添加一条评论
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /post/{postId}/comment [get]
func AddComment(c *gin.Context) {
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

// @Summary 评论列表
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /post/{postId}/comment [post]
func ListComments(c *gin.Context) {
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

// @Summary 删除评论
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /post/{postId}/comment/{commentId} [delete]
func DeleteComment(c *gin.Context) {
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
