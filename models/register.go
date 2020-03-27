package models

type Register struct {
	Uid        string `json:"uid" bson:"uid"`                 //用户全局ID
	NickName   string `json:"nick_name" bson:"nick_name"`     //昵称
	PassWord   string `json:"pass_word" bson:"pass_word"`     //密码
	InviteCode string `json:"invite_code" bson:"invite_code"` //邀请码
}

type Team struct {
	Uid      string `json:"uid" bson:"uid"`             //用户id
	FatherId string `json:"father_id" bson:"father_id"` //推荐人id
}
