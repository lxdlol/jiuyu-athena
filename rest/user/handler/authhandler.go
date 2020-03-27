package handler

import (
	"athena/common"
	"athena/models"
	"athena/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Login(c *gin.Context) {
	var p UserParam
	if e := c.ShouldBind(&p); e != nil {
		common.GResp.Failure(c, common.CodeIllegalParam)
		return
	}
	if !utils.Rightful(p.PassWord, 6, 20) {
		common.GResp.Failure(c, common.CodeIllegalParam)
		return
	}
	if !utils.Rightful(p.UserName, 1, 20) {
		common.GResp.Failure(c, common.CodeUserError)
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
		common.GResp.Failure(c, common.CodeIllegalPwdLength)
		return
	}
	if !utils.Rightful(p.UserName, 1, 20) {
		common.GResp.Failure(c, common.CodeUserError)
		return
	}
	//检验推荐码是否正确
	var user = models.User{InviteCode: p.RCode}
	if p.RCode != "" {
		if _, e := user.FindByFilter("r_code", p.RCode); e != nil {
			common.GResp.Failure(c, common.CodeRCodeError)
			return
		}
	}
	//注册
	user = models.User{
		Uid:           utils.Rid(),
		NickName:      p.UserName,
		Level:         models.Level{},
		Avatar:        "",
		City:          "",
		IsGuru:        false,
		GuruRanking:   0,
		IsMaster:      false,
		MasterRanking: 0,
		Follows:       0,
		InviteCode:    "",
	}
	i, _ := strconv.Atoi(user.Uid)
	user.InviteCode = utils.InviteCode.IdToCode(uint64(i))
	if e := user.Insert(); e != nil {
		common.GResp.Failure(c, common.CodeIsExist)
		return
	}
	common.GResp.Success(c, nil)
}
