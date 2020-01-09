package controllers

import (
	"p_web/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// ArticleTypeController is a struct
type ArticleTypeController struct {
	beego.Controller
}

// ShowArticleType 文章分类显示
func (c *ArticleTypeController) ShowArticleType() {
	c.Layout = "base/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "文章分类列表"
	c.TplName = "article_add_type.html"
	// 查询分类信息
	o := orm.NewOrm()
	var ArticleTypes []*models.ArticleType
	qs := o.QueryTable("ArticleType")
	_, err := qs.All(&ArticleTypes)
	if err != nil {
		beego.Info("没有分类数据")
	}
	c.Data["ArticleTypes"] = ArticleTypes

}

// HandleArticleType 文章分类添加
func (c *ArticleTypeController) HandleArticleType() {
	Atype := c.GetString("Atype")
	beego.Info(Atype)
	if Atype == "" {
		beego.Info("数据不能为空")
	} else {
		o := orm.NewOrm()
		ArticleType := models.ArticleType{}
		ArticleType.TypeName = Atype
		_, err := o.Insert(&ArticleType)
		if err != nil {
			beego.Info("插入数据失败")
		}
	}
	c.Redirect("/atype_add", 302)

}
