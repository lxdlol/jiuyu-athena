package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Trade struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

/*
项目
*/
type Project struct {
	User      User      `json:"user" bson:user`           //发布人
	Trade     Trade     `json:"trade" bson:"trade"`       //行业
	Tags      []Tag     `json:"tags" bson:"tags"`         //标签
	Comments  []Comment `json:"comments" bson:"comments"` //备注
	Follows   []User    `json:"follows" bson:"follows"`   //关注这个项目的人
	Reports   []Report  `json:"reports" bson:"reports"`   //评测报告
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
type Report struct {
	User      User   `json:"user" bson:"user"`  //上传人
	Title     string `json:"title" bson:"tile"` //标题
	FilePath  string
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
