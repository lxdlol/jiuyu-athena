package models

import (
	"fmt"
	"testing"
)

func TestInsertPost(t *testing.T) {
	//id := bson.NewObjectId()

	post := Post{Content: "测试新增post"}
	err := InsertPost(post)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("TestInsertPost")
}

func TestPost_DoEgg(t *testing.T) {
	fmt.Println("TestPost_DoEgg")
}
