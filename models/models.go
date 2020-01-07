package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User 用户表
type User struct {
	ID       int
	UserName string
	Password string
}

// Article 文章表
type Article struct {
	ID       int    `orm:"pk;auto"`
	Artiname string `orm:"size(20)"`
	Atime    time.Time
	Acount   int `orm:"default(20);null"`
	Acontent string
	Aimg     string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test_beego?charset=utf8")
	orm.RegisterModel(new(User), new(Article))
	orm.RunSyncdb("default", false, true)
}
