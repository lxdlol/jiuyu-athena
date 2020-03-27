package models

type Auth struct {
	Uid      string `json:"uid" bson:"uid"`           //用户全局ID
	Username string `json:"username" bson:"username"` //昵称
	Password string `json:"password" bson:"password"` //密码
}
