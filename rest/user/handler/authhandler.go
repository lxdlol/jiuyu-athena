package handler

import (
	"athena/common"
	"athena/log"
	"athena/models"
	"athena/rest/middleware"
	"athena/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func LoginHandler(c *gin.Context) {
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
	var auth models.Auth
	auth.Username = p.UserName
	filter, e := auth.FindByFilter("nick_name", auth.Username)
	if e != nil {
		common.GResp.Failure(c, common.CodeUserError)
		return
	}
	Salt := utils.NewPwdSalt(auth.Uid, 1)
	if filter.Password != utils.Hash256(p.PassWord, Salt) {
		common.GResp.Failure(c, common.CodePasswordError)
		return
	}
	ok, token := middleware.GenerateToken(auth)
	if !ok {
		common.GResp.Failure(c, common.CodeInternalServerError)
		return
	}
	common.GResp.Success(c, gin.H{"token": token})
}

//注册
func RegisterHandler(c *gin.Context) {
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
	var auth models.Auth
	Salt := utils.NewPwdSalt(user.Uid, 1)
	pwd := utils.Hash256(p.PassWord, Salt)
	auth = models.Auth{
		Uid:      user.Uid,
		Username: user.NickName,
		Password: pwd,
	}
	if err := auth.Insert(); err != nil {
		log.Log.Error(err)
		common.GResp.Failure(c, common.CodeDbInsertError)
		return
	}
	i, _ := strconv.Atoi(user.Uid)
	user.InviteCode = utils.InviteCode.IdToCode(uint64(i))
	if e := user.Insert(); e != nil {
		log.Log.Error(e)
		common.GResp.Failure(c, common.CodeIsExist)
		return
	}
	//异步计算团队
	models.TeamChan <- user
	common.GResp.Success(c, nil)
}
