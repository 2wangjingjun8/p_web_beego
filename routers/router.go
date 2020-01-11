package routers

import (
	"p_web/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/article*", beego.BeforeRouter, filterFunc)
	beego.Router("/", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{}, "get:ShowIndex;post:HandleIndex")

	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/logout", &controllers.UserController{}, "get:HandleLogout")
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister;post:HandleRegister")
	// 文章管理模块
	beego.Router("/article_list", &controllers.ArticleController{}, "get:ShowArticle;post:SelectArticle")
	beego.Router("/article_add", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/article_edit", &controllers.ArticleController{}, "get:ShowEditArticle;post:HandleEditArticle")
	beego.Router("/article_detail", &controllers.ArticleController{}, "get:ShowArticleDetail")
	beego.Router("/article_del", &controllers.ArticleController{}, "get:HandleArticleDel")
	// 文章类型管理模块
	beego.Router("/atype_add", &controllers.ArticleTypeController{}, "get:ShowArticleType;post:HandleArticleType")

	beego.Router("/redis", &controllers.RedisDemoController{})

}

var filterFunc = func(ctx *context.Context) {
	username := ctx.Input.Session("username")
	if username == nil {
		ctx.Redirect(302, "/login")
	}
}
