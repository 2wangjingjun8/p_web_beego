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

func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}

func (c *MainController) HandleLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" || password == "" {
		beego.Info("数据不能为空！")
		c.Redirect("/login", 302)
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.UserName = username
	err := o.Read(&user, "UserName")
	beego.Info(user)
	if err != nil {
		beego.Info("账号不存在", err)
		c.Redirect("/login", 302)
		return
	} else {
		if user.Password == password {
			c.Ctx.WriteString("登录成功")
		} else {
			c.Ctx.WriteString("密码错误")
			c.Redirect("/login", 302)
			return
		}
	}

}

func (c *MainController) ShowIndex() {
	c.TplName = "index.html"
}

func (c *MainController) HandleIndex() {

}

func (c *MainController) ShowArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "文章列表"
}

func (c *MainController) ShowAddArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_add.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "新增文章"
}

func (c *MainController) HandleAddArticle() {
	Artiname := c.GetString("Artiname")
	AType := c.GetString("AType")
	Acontent := c.GetString("Acontent")
	beego.Info(Artiname, AType, Acontent)
}

func (c *MainController) ShowEditArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_edit.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "更新文章"
}
func (c *MainController) HandleEditArticle() {
	Artiname := c.GetString("Artiname")
	beego.Info(Artiname)
}
