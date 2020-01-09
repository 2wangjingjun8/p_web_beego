package controllers

import (
	"p_web/models"
	"path"
	"time"
	"math"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// ArticleController is a struct
type ArticleController struct {
	beego.Controller
}

// ShowArticle 文章列表
func (c *ArticleController) ShowArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "文章列表"
	c.Data["addurl"] = "/article_add"

	o := orm.NewOrm()
	var articles []models.Article
	qs := o.QueryTable("Article")
	count,_ := qs.Count()
	pagesize := 2
	// 总页码
	pageNum := math.Ceil(float64(count)/float64(pagesize))
	pageIndex,err := c.GetInt("page")
	beego.Info(pageIndex)
	beego.Info(err)
	start := pagesize * (pageIndex - 1)
	_, err = qs.Limit(pagesize, start).All(&articles)
	// _, err := qs.All(&articles)
	if err != nil {
		beego.Info("查询数据失败", err)
		return
	}
	beego.Info(articles)
	c.Data["articles"] = articles
	c.Data["count"] = count
	c.Data["pageNum"] = pageNum
	c.Data["pageIndex"] = pageIndex

	preIndex := pageIndex - 1
	if preIndex <= 0 {
		preIndex = 1
	}
	nextIndex := pageIndex + 1
	if nextIndex > int(pageNum) {
		nextIndex = int(pageNum)
	}
	c.Data["preIndex"] = preIndex
	c.Data["nextIndex"] = nextIndex
}

// ShowAddArticle 文章添加
func (c *ArticleController) ShowAddArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_add.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "新增文章"
}

// HandleAddArticle 处理文章添加
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
				beego.Info("插入数据失败",err)
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

// ShowArticleDetail 文章详情
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

// ShowEditArticle 文章编辑
func (c *ArticleController) ShowEditArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_edit.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "更新文章"

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
	c.Data["article"] = article

}

// HandleEditArticle 处理文章编辑
func (c *ArticleController) HandleEditArticle() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取参数错误", err)
	}
	beego.Info("id", id)
	Artiname := c.GetString("Artiname")
	// AType := c.GetString("AType")
	Acontent := c.GetString("Acontent")
	if Artiname == "" || Acontent == "" {
		beego.Info("数据不能为空")
		return
	}
	f, h, err := c.GetFile("img")
	beego.Info(f, h, err)
	var filename string
	if err != nil {
		beego.Info("没有上传图片")
	} else {
		defer f.Close()
		fileext := path.Ext(h.Filename)
		beego.Info(fileext)
		if fileext == ".jpg" || fileext == ".png" || fileext == ".jpeg" {
			filename = time.Now().Format("2006-01-02_150405") + fileext
			beego.Info(filename)
			c.SaveToFile("img", "./static/img/"+filename)


		} else {
			beego.Info("图片格式不正确")
			return
		}
		if h.Size > 500000000 {
			beego.Info("图片大小不符合")
			return
		}
	}


	o := orm.NewOrm()
	Arti := models.Article{ID: id}
	err = o.Read(&Arti)
	if err != nil {
		beego.Info("获取数据错误")
	}
	
	Arti.Artiname = Artiname
	// Arti.AType = AType
	Arti.Acontent = Acontent
	if filename != "" {
		Arti.Aimg = "/static/img/" + filename
	}else{
		Arti.Aimg = c.GetString("img_old")
	}
	_, err = o.Update(&Arti,"Artiname","Acontent","Aimg")
	if err != nil {
		beego.Info("更新数据失败",err)
		return
	}
	c.Redirect("/article_list", 302)

}

// HandleArticleDel 处理文章删除
func (c *ArticleController) HandleArticleDel()  {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取参数错误", err)
	}
	o:= orm.NewOrm()
	arti := models.Article{ID:id}
	err = o.Read(&arti)
	if err != nil{
		beego.Info("获取数据错误")
	}
	_,err = o.Delete(&arti)
	if err != nil {
		beego.Info("删除数据失败",err)
		return
	}
	c.Redirect("/article_list", 302)
	
}

// ShowArticleType 文章分类显示
func (c *ArticleController) ShowArticleType()  {
	c.Layout = "base/layout.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "文章分类列表"
	c.TplName = "article_add_type.html"
	// 查询分类信息
	o :=orm.NewOrm()
	var ArticleTypes []*models.ArticleType
	qs := o.QueryTable("ArticleType")
	_, err := qs.All(&ArticleTypes)
	if err != nil{
		beego.Info("没有分类数据")
	}
	c.Data["ArticleTypes"] = ArticleTypes

}

// HandleArticleType 文章分类添加
func (c *ArticleController) HandleArticleType()  {
	Atype := c.GetString("Atype")
	beego.Info(Atype)
	if Atype == "" {
		beego.Info("数据不能为空")
	}else{
		o :=orm.NewOrm()
		ArticleType := models.ArticleType{}
		ArticleType.TypeName = Atype
		_,err := o.Insert(&ArticleType)
		if err != nil{
			beego.Info("插入数据失败")
		}
	}
	c.Redirect("/atype_add",302)

}