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
	Articles []*Article `orm:"rel(m2m)"`
}

// Article 文章表
type Article struct {
	ID       int       `orm:"pk;auto"`
	Artiname string    `orm:"size(60)"`
	Atime    time.Time `orm:"auto_now"`
	Acount   int       `orm:"default(20);null"`
	Acontent string
	Aimg     string
	ArticleType *ArticleType `orm:"rel(fk)"`
	Users []*User `orm:"reverse(many)"`
}

// ArticleType 文章类型表
type ArticleType struct{
	ID int
	TypeName string `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/test_beego?charset=utf8")
	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	orm.RunSyncdb("default", false, true)
}
