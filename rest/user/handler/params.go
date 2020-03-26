package handler

type UserParam struct {
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"pass_word" form:"pass_word"`
	RCode    string `json:"r_code" form:"r_code"`
}
