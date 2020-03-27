package models

import (
	"athena/db"
	"gopkg.in/mgo.v2/bson"
)

func InsertPost(post Post) error {
	_, c := db.Connect("post")
	err := c.Insert(post)
	return err
}
func ListPostByUserId(userId string, pageSize int64, page int64) error {
	_, c := db.Connect("post")
	_ = c.Find(bson.M{"user._id": userId})
	return nil
}
func ListPost(topic Topic, pageSize, page int64) (error, []Post) {
	var posts []Post
	_, c := db.Connect("post")
	_ = c.Find(bson.M{"$sort": "-updated_at"})
	return nil, posts
}

func DeletePost(id string) error {
	_, c := db.Connect("post")
	return c.RemoveId(id)
}

func (c *Post) Update(update interface{}) error {
	_, collect := db.Connect("post")
	err := collect.UpdateId(c.Id, c)
	return err
}
func (c *Post) InsertComment(comment Comment) error {
	return c.Update(bson.M{"$push": bson.M{"comments": comment}})
}

func (c *Post) DeleteComment(comment Comment) error {
	return c.Update(bson.M{"$pop": bson.M{"comments": comment}})
}

/*
丢鸡蛋
*/
func (c *Post) DoEgg() error {
	return c.Update(bson.M{"$inc": "eggs"})
}

/*
点赞
*/
func (c *Post) DoFlower() error {
	return c.Update(bson.M{"$inc": "flowers"})
}

//阅读数
func (c *Post) DoView() error {
	c.Views = c.Views + 1
	return c.Update(bson.M{"$inc": "views"})
}

/*
	用户是否有查看权限
*/

func (c *Post) HasViewPermission(user User) (error, bool) {
	return nil, true
}
