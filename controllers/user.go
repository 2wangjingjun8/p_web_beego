package controllers

import (
	"p_web/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)
// UserController 用户结构体
type UserController struct {
	beego.Controller
}

// ShowLogin 登陆显示
func (c *UserController) ShowLogin() {
	username := c.Ctx.GetCookie("username")
	if username != "" {
		c.Data["username"] = username
		c.Data["check"] = "checked"
	}
	
	c.TplName = "login.html"
}

// HandleLogin 登陆处理
func (c *UserController) HandleLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	remember := c.GetString("remember")
	beego.Info(remember)
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
	} else {
		if user.Password == password {
			// c.Ctx.WriteString("登录成功")
			if remember == "on" {
				c.Ctx.SetCookie("username",username,time.Second*3600)
			}else{
				c.Ctx.SetCookie("username",username,-1)
			}
			c.SetSession("username",username)
			
			c.Redirect("/index",302)
		} else {
			c.Ctx.WriteString("密码错误")
			c.Redirect("/login", 302)
		}
	}

}

// HandleLogout 退出登录
func (c *UserController) HandleLogout()  {
	c.DelSession("username")
	c.Redirect("/",302)
}

// ShowRegister 注册显示
func (c *UserController) ShowRegister() {
	c.TplName = "register.html"
}

// HandleRegister 注册处理
func (c *UserController) HandleRegister() {
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