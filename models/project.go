package models

import "time"

type Trade struct {
	Name string `json:"name" bson:"name"`
}

type Project struct {
	User      User      `json:"user" bson:user`
	Trade     Trade     `json:"trade" bson:"trade"`
	Tags      []Tag     `json:"tags" bson:"tags"`
	Comments  []Comment `json:"comments" bson:"comments"`
	Follows   []User    `json:"follows" bson:"follows"`
	Reports   []Report  `json:"reports" bson:"reports"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Report struct {
	User      User
	Title     string
	FilePath  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
