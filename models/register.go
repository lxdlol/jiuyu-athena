package models

type Auth struct {
	Uid        string `json:"uid" bson:"uid"`                 //用户全局ID
	UserName   string `json:"user_name" bson:"user_name"`     //昵称
	Password   string `json:"password" bson:"password"`       //密码
	InviteCode string `json:"invite_code" bson:"invite_code"` //邀请码
}
