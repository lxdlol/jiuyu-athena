package models

import (
	"athena/db"
	"gopkg.in/mgo.v2/bson"
)

type Auth struct {
	Uid      string `json:"uid" bson:"uid"`           //用户全局ID
	Username string `json:"username" bson:"username"` //昵称
	Password string `json:"password" bson:"password"` //密码
}

func (a Auth) Insert() error {
	session, collection := db.Connect("auth")
	defer session.Close()
	return collection.Insert(&a)
}
func (a Auth) FindByFilter(filter string, value interface{}) (Auth, error) {
	s, c := db.Connect("auth")
	defer s.Close()
	var auth Auth
	err := c.Find(bson.M{filter: value}).One(&auth)
	return auth, err
}
