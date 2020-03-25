package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// NewResponse 成功返回
func NewResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: 200,
		Msg:  "",
		Data: data,
	})
}

// NewErrorResponse 错误返回
func NewErrorResponse(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
