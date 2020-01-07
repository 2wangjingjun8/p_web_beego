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
	beego.Router("/article_list", &controllers.MainController{}, "get:ShowArticle")
	beego.Router("/article_add", &controllers.MainController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/article_edit", &controllers.MainController{}, "get:ShowEditArticle;post:HandleEditArticle")
}
