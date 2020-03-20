package models

import "time"

type EventType struct {
	/**
	1.create post
	2.comment post
	3.upload report
	4.user upgrade 用户升级
	*/

}

type Event struct {
	EventType EventType `json:"event_type" bson:"event_type"` //事件类型
	Params    []string  `json:"entity_id" json:"entity_id"`
	Title     string    `json:"title" json:"title"`
	CreatedAt string    `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Comments  []Comment `json:"comments" bson:"comments"` //评论
	Views     int64     `json:"views" bson:"views"`       //阅读数
	Flowers   int64     `json:"flowers" bson:"flowers"`   //鲜花数
	Eggs      int64     `json:"eggs" bson:"eggs"`         //鸡蛋数

}
