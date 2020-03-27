package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type PostType int32
type PostEventType int32 //kafka消息类型

const (
	Message PostType = 0
	Event   PostType = 1
)

const (
	UploadReport      PostEventType = 0 //上传报告
	AddReportComment  PostEventType = 1 //报告新评论
	AddProject        PostEventType = 2 //系统发布(推送)的消息
	AddProjectComment PostEventType = 3 //项目评论
	NewMaster         PostEventType = 4 //晋升大咖
	NewGuru           PostEventType = 5 //晋升达人

)

type Topic struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`       //required
	Comment string        `json:"comment" bson:"comment"` //描述
	Image   string        `json:"image" bson:"image"`     //图标

	CreatedAt time.Time `json:"created_at" bson:"created_at"` //required
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"` //required
}

/*
文章,博客,动态
*/
type Post struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Type     PostType      `json:"type" bson:"type"` //1.message, 2.event
	Images   []string      `json:"images" bson:"images"`
	Content  string        `json:"content" bson:"content"`       //required
	Comments []Comment     `json:"comments" bson:"comments"`     //评论
	Topic    Topic         `json:"topic" bson:"topic,omitempty"` //required
	Tags     []string      `json:"tags" bson:"tags"`
	User     User          `json:"user" bson:"user,omitempty"` //required,发布人

	Views   int64 `json:"views" bson:"views"`     //阅读数
	Flowers int64 `json:"flowers" bson:"flowers"` //鲜花数
	Eggs    int64 `json:"eggs" bson:"eggs"`       //鸡蛋数

	CreatedAt time.Time `json:"created_at" bson:"created_at"` //required
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"` //required

}

// Comment  评论
type Comment struct {
	Id        bson.ObjectId `json:"id" bson:"id"`
	Content   string        `json:"content" bson:"content"`
	User      User          `json:"user" bson:"user,omitempty"`
	Flowers   int64         `json:"flowers" bson:"flowers"` //鲜花数
	Eggs      int64         `json:"eggs" bson:"eggs"`       //鸡蛋数
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
}

type Tag struct {
	Name   string `json:"name" bson:"Name"`     //unique
	Number int64  `json:"number" bson:"number"` //内容数量
}

type GalleyGroup struct {
	Name   string       `json:"name" bson:"name"`
	Parent *GalleyGroup `json:"parent" bson:"parent"`
}

type Gallery struct {
	Group    GalleyGroup `json:"group" bson:"group"`
	Name     string      `json:"name" bson:"name"`
	FilePath string      `json:"file_path" bson:"file_path"`
}

/*
以下设计还不稳定
*/
type PostConstraint struct {
	Topic Topic  `json:"topic" json:"topic"`
	Rules []Rule `json:"rules" bson:"rules"` //满足一条rule即可
}

type Rule struct {
	/*
		user_level_gte：大于或等于指定等级用户
		user_level_only:只允许指定等级用户
		user_level_pay_and:达到用户等级并购买
		pay_only:支付
	*/
	Value interface{} `json:"value" bson:"value"`
}
