package handler

import (
	"athena/common"
	"github.com/gin-gonic/gin"
	"strings"
)

func Login(c *gin.Context) {
	var p UserParam
	if e := c.ShouldBind(&p); e != nil {
		common.NewErrorResponse(c, 400, "数据格式错误", nil)
		return
	}
	pwd := strings.Replace(p.PassWord, " ", "", -1)
	//判断pwd
	if len(pwd) < 6 || len(pwd) > 20 {
		common.NewErrorResponse(c, 10001, "密码长度错误", nil)
		return
	}
	username := strings.Replace(p.UserName, " ", "", -1)

}

func Register(c *gin.Context) {
	var p UserParam
	if e := c.ShouldBind(&p); e != nil {
		common.NewErrorResponse(c, 400, "数据格式错误", nil)
		return
	}
	pwd := strings.Replace(p.PassWord, " ", "", -1)
	//判断pwd
	if len(pwd) < 6 || len(pwd) > 20 {
		common.NewErrorResponse(c, 10001, "密码长度错误", nil)
		return
	}
	username := strings.Replace(p.UserName, " ", "", -1)
	if username == "" {
		common.NewErrorResponse(c, 400, "数据格式错误", nil)
		return
	}

}
