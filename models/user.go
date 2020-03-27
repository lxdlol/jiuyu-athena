package models

import (
	"athena/db"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var TeamChan = make(chan User, 1000)

type Level struct {
	Name  string `json:"name" bson:"name"`   //等级名称
	Image string `json:"image" bson:"image"` //等级图标
	Value int8   `json:"value" bson:"value"` //等级
}

type User struct {
	Uid           string `json:"uid" bson:"uid"` //用户全局ID
	NickName      string `json:"nick_name" bson:"nick_name"`
	Level         Level  `json:"level" bson:"level"` //用户等级
	Avatar        string `json:"avatar" bson:"avatar"`
	City          string `json:"city" bson:"city"`
	IsGuru        bool   `json:"is_guru" bson:"is_guru"`               //是否是达人
	GuruRanking   int64  `json:"guru_ranking" json:"guru_ranking"`     //达人指数
	IsMaster      bool   `json:"is_master" bson:"is_master"`           //是否是大咖
	MasterRanking int64  `json:"master_ranking" bson:"master_ranking"` //大咖指数
	Follows       int64  `json:"follows" bson:"follows"`               //关注数
	InviteCode    string `json:"invite_code" bson:"invite_code"`       //邀请码,
	ParentId      string `json:"parent_id"`
	IsStaff       bool   `json:"is_staff" bson:"is_staff"` //是否后台账户
}

type Team struct {
	Owner   User
	Members []Member
}

type Member struct {
	Uid   string
	Level int64
	User  User
}

//条件查询
func (u User) FindByFilter(filter string, value interface{}) (User, error) {
	s, c := db.Connect("user")
	defer s.Close()
	var user User
	err := c.Find(bson.M{filter: value}).One(&user)
	return user, err
}

func (u User) Insert() error {
	s, c := db.Connect("user")
	defer s.Close()
	insert := c.Insert(&u)
	return insert
}

type Follow struct {
	/*
		关注列表
	*/
	User      User      `json:"user" bson:"user"`
	Follow    User      `json:"follow" bson:"follow"` //被关注的用户
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
