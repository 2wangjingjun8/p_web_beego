package controllers

import (
	"p_web/models"
	"path"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// ArticleController is a struct
type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) ShowArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "文章列表"

	o := orm.NewOrm()
	var articles []models.Article
	_, err := o.QueryTable("Article").All(&articles)
	if err != nil {
		beego.Info("查询数据失败", err)
		return
	}
	beego.Info(articles)
	c.Data["articles"] = articles
}

func (c *ArticleController) ShowAddArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_add.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "新增文章"
}

func (c *ArticleController) HandleAddArticle() {
	Artiname := c.GetString("Artiname")
	// AType := c.GetString("AType")
	Acontent := c.GetString("Acontent")
	if Artiname == "" || Acontent == "" {
		beego.Info("数据不能为空")
		return
	}
	// var filename string
	f, h, err := c.GetFile("img")
	defer f.Close()
	if err != nil {
		beego.Info("上传失败")
		return
	} else {
		fileext := path.Ext(h.Filename)
		beego.Info(fileext)
		if fileext == ".jpg" || fileext == ".png" || fileext == ".jpeg" {
			filename := time.Now().Format("2006-01-02_150405") + fileext
			beego.Info(filename)
			c.SaveToFile("img", "./static/img/"+filename)

			o := orm.NewOrm()
			Arti := models.Article{}
			Arti.Artiname = Artiname
			// Arti.AType = AType
			Arti.Acontent = Acontent
			Arti.Aimg = "/static/img/" + filename
			_, err = o.Insert(&Arti)
			if err != nil {
				beego.Info("插入数据失败")
				return
			}
			c.Redirect("/article_list", 302)

		} else {
			beego.Info("图片格式不正确")
			return
		}
		if h.Size > 500000000 {
			beego.Info("图片大小不符合")
			return
		}
	}
}

func (c *ArticleController) ShowArticleDetail() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取参数错误", err)
	}
	beego.Info("id", id)
	o := orm.NewOrm()
	article := models.Article{ID: id}
	err = o.Read(&article)
	if err != nil {
		beego.Info("获取数据错误")
	}
	beego.Info("article", article)

	c.Layout = "base/layout.html"
	c.TplName = "article_detail.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "文章详情"
	c.Data["article"] = article
}

func (c *ArticleController) ShowEditArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_edit.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "更新文章"
}
func (c *ArticleController) HandleEditArticle() {
	Artiname := c.GetString("Artiname")
	beego.Info(Artiname)
}
