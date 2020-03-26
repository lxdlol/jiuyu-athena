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
	if !utils.Rightful(p.PassWord, 6, 20) {
		common.NewResponse(c, common.ErrCode(10001))
		return
	}
	if !utils.Rightful(p.UserName, 1, 20) {
		common.NewResponse(c, common.ErrCode(10002))
		return
	}

}

//注册
func Register(c *gin.Context) {
	var p UserParam
	if e := c.ShouldBind(&p); e != nil {
		common.GResp.Failure(c, common.CodeIllegalParam)
		return
	}
	//检验数据正确
	if !utils.Rightful(p.PassWord, 6, 20) {
		common.NewResponse(c, common.ErrCode(10001))
		return
	}
	if !utils.Rightful(p.UserName, 1, 20) {
		common.NewResponse(c, common.ErrCode(10002))
		return
	}
	//检验推荐码是否正确
	if p.RCode != "" {
		if _, e := models.FindByFilter("r_code", p.RCode); e != nil {
			common.NewResponse(c, common.ErrCode(10003))
			return
		}
	}

}
