package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// User 用户表
type User struct {
	ID       int
	UserName string
	Password string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test_beego?charset=utf8")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
