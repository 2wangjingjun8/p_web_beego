package routers

import (
	"p_web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{}, "get:ShowIndex;post:HandleIndex")
}
