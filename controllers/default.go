package controllers

import (
	"p_web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// 插入
	// o := orm.NewOrm()
	// user := models.User{}
	// user.UserName = "maomao"
	// user.Password = "123456"
	// _, err := o.Insert(&user)
	// if err != nil {
	// 	beego.Info("插入失败", err)
	// 	return
	// } else {
	// 	beego.Info("插入成功")
	// }

	// 查询1
	// o := orm.NewOrm()
	// user := models.User{}
	// user.ID = 1
	// err := o.Read(&user)
	// if err != nil {
	// 	beego.Info("查询失败")
	// } else {
	// 	beego.Info("查询成功", user)
	// }
	// 查询2
	// o := orm.NewOrm()
	// user := models.User{}
	// user.UserName = "maomao"
	// err := o.Read(&user, "UserName")
	// if err != nil {
	// 	beego.Info("查询失败")
	// } else {
	// 	beego.Info("查询成功", user)
	// }

	// 更新
	// o := orm.NewOrm()
	// user := models.User{}
	// user.ID = 1
	// err := o.Read(&user)
	// if err == nil {
	// 	user.Password = "w123456"
	// 	_, err = o.Update(&user)
	// 	if err == nil {
	// 		beego.Info("更新成功")
	// 	} else {
	// 		beego.Info("更新失败", err)
	// 	}
	// }

	// 删除
	// o := orm.NewOrm()
	// user := models.User{}
	// user.ID = 1
	// _, err := o.Delete(&user)
	// if err != nil {
	// 	beego.Info("删除失败", err)
	// } else {
	// 	beego.Info("删除" + strconv.Itoa(user.ID) + "成功")
	// }

	// c.Data["username"] = "xiaoxiao"
	// c.Data["user"] = user
	c.TplName = "register.html"
}

func (c *MainController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		beego.Info("数据不能为空！")
		c.Redirect("/register", 302)
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.UserName = username
	user.Password = password
	_, err := o.Insert(&user)
	if err != nil {
		beego.Info("数据插入失败！", err)
		c.Redirect("/register", 302)
		return
	}
	c.Ctx.WriteString("注册成功！")
	// c.Data["username"] = "maomao"
	// c.TplName = "register.html"
}

// ShowIndex 首页
func (c *MainController) ShowIndex() {
	// username := c.GetSession("username")
	// if username == nil {
	// 	c.Redirect("/login",302)
	// }
	c.TplName = "index.html"
	c.Layout = "base/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "首页"
}

// HandleIndex 处理首页
func (c *MainController) HandleIndex() {

}
