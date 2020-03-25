package handler

import "github.com/gin-gonic/gin"

// @Summary 获取首页
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Failure 400 {object} APIError "We need ID!!"
// @Failure 404 {object} APIError "Can not find ID"
// @Router /index [get]
func IndexGet(c *gin.Context) {

}
