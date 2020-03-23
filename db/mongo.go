package db

import (
	"athena/config"
	"fmt"
	"gopkg.in/mgo.v2"
	"strconv"
	"time"
)

var mgosess *mgo.Session

//初始化链接
func init() {
	athenaConfig := config.GetConfig()
	dialinfo := mgo.DialInfo{
		Addrs:     []string{athenaConfig.DataSource.Host + ":" + strconv.Itoa(int(athenaConfig.DataSource.Port))},
		Timeout:   500 * time.Millisecond,
		Username:  athenaConfig.DataSource.User,
		Source:    "admin",
		Password:  athenaConfig.DataSource.Password,
		PoolLimit: 100,
	}
	session, e := mgo.DialWithInfo(&dialinfo)
	//path, _ := cnf.Cnf.GetValue("cnf", "DbPath")
	//session, e := mgo.Dial(cnf.CnfStr.Dbpath)
	if e != nil {
		fmt.Println("connect mongo error:" + e.Error())
	}
	mgosess = session
}

//copy链接选择数据库和表
func Connect(collection string) (*mgo.Session, *mgo.Collection) {
	athenaConfig := config.GetConfig()
	ms := mgosess.Copy()
	c := ms.DB(athenaConfig.DataSource.Database).C(collection)
	ms.SetMode(mgo.Monotonic, true)
	return ms, c
}
