/*


 */
package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Topic struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`       //required
	Comment string        `json:"comment" bson:"comment"` //描述
	Image   string        `json:"image" bson:"image"`     //图标

	CreatedAt time.Time `json:"created_at" bson:"created_at"` //required
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"` //required
}

type Post struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Title   string        `json:"title" bson:"title"`
	Content string        `json:"content" bson:"content"` //required
	Topic   Topic         `json:"topic" bson:"topic"`     //required
	Tags    []string      `json:"tags" bson:"tags"`
	User    User          `json:"user" bson:"user"` //required

	Views   int64 `json:"views" bson:"views"`     //阅读数
	Flowers int64 `json:"flowers" bson:"flowers"` //鲜花数
	Eggs    int64 `json:"eggs" bson:"eggs"`       //鸡蛋数

	CreatedAt time.Time `json:"created_at" bson:"created_at"` //required
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"` //required

}

/*
	用户是否有查看权限
*/

func (c *Post) HasViewPermission(user User) (error, bool) {
	return nil, true
}

type Comment struct {
	Content string `json:"content" bson:"content"`
	User    User   `json:"user" bson:"user"`

	Flowers int64 `json:"flowers" bson:"flowers"` //鲜花数
	Eggs    int64 `json:"eggs" bson:"eggs"`       //鸡蛋数

	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type Tag struct {
	Name   string `json:"name" bson:"Name"`
	Number int64  `json:"number" bson:"number"` //内容数量
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
