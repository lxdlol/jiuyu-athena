package handler

import (
	"athena/common"
	"github.com/gin-gonic/gin"
	"strings"
)

func Login(c *gin.Context) {
	var p UserParam
	if e := c.ShouldBind(&p); e != nil {
		common.GResp.Failure(c, common.CodeIllegalParam)
		return
	}
	pwd := strings.Replace(p.PassWord, " ", "", -1)
	//判断pwd
	if len(pwd) < 6 || len(pwd) > 20 {
		common.GResp.Failure(c, common.CodeIllegalPwdLength)
		return
	}
	//username := strings.Replace(p.UserName, " ", "", -1)

}

func Register(c *gin.Context) {
	var p UserParam
	if e := c.ShouldBind(&p); e != nil {
		common.GResp.Failure(c, common.CodeIllegalParam)
		return
	}
	pwd := strings.Replace(p.PassWord, " ", "", -1)
	//判断pwd
	if len(pwd) < 6 || len(pwd) > 20 {
		common.GResp.Failure(c, common.CodeIllegalPwdLength)
		return
	}
	username := strings.Replace(p.UserName, " ", "", -1)
	if username == "" {
		common.GResp.Failure(c, common.CodeIllegalParam)
		return
	}

}
