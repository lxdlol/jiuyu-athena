package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Trade 行业
type Trade struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`    // 行业id
	Name string        `json:"name" bson:"name"` // 行业
}

// Project 项目
type Project struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`                // 项目id
	User      User          `json:"user" bson:"user"`             // 发布人
	Trade     Trade         `json:"trade" bson:"trade"`           // 行业
	Tags      []Tag         `json:"tags" bson:"tags"`             // 标签
	Comments  []Comment     `json:"comments" bson:"comments"`     // 备注
	Follows   []User        `json:"follows" bson:"follows"`       // 关注这个项目的人
	Reports   []Report      `json:"reports" bson:"reports"`       // 评测报告
	Suggest   int8          `json:"suggest" bson:"suggest"`       // 推荐
	CreatedAt time.Time     `json:"created_at" bson:"created_at"` // 创建时间
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"` // 更新时间
}

// Report 报告
type Report struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`                // 报告id
	User      User          `json:"user" bson:"user"`             // 上传人
	Title     string        `json:"title" bson:"title"`           // 标题
	FilePath  string        `json:"file_path" bson:"file_path"`   // 文件路径
	CreatedAt time.Time     `json:"created_at" bson:"created_at"` // 创建时间
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"` // 更新时间
}
