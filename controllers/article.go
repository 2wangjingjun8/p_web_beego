package controllers

import (
	"bytes"
	"encoding/gob"
	"math"
	"p_web/models"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
)

// ArticleController is a struct
type ArticleController struct {
	beego.Controller
}

// SelectArticle 选择文章类型
func (c *ArticleController) SelectArticle() {
	TypeID := c.GetString("TypeID")
	o := orm.NewOrm()
	var articles []models.Article
	o.QueryTable("Article").RelatedSel("ArticleType").Filter("ArticleType__ID", TypeID).All(&articles)
	beego.Info(articles)
}

// ShowArticle 文章列表
func (c *ArticleController) ShowArticle() {
	c.Layout = "base/layout.html"
	c.TplName = "article_list.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Sidebar"] = "base/nav.html"
	c.Data["page_head"] = "文章列表"
	c.Data["addurl"] = "/article_add"

	TypeID := c.GetString("TypeID")
	o := orm.NewOrm()
	var articles []models.Article
	qs := o.QueryTable("Article").RelatedSel("ArticleType")

	var count int64
	if TypeID == "" {
		count, _ = qs.Count()
	} else {
		count, _ = qs.Filter("ArticleType__ID", TypeID).Count()
	}

	pagesize := 2
	// 总页码
	pageNum := math.Ceil(float64(count) / float64(pagesize))
	pageIndex, _ := c.GetInt("page")
	beego.Info(pageIndex)
	if pageIndex == 0 {
		pageIndex = 1
	}
	start := pagesize * (pageIndex - 1)

	if TypeID == "" {
		qs.Limit(pagesize, start).All(&articles)
	} else {
		qs.Limit(pagesize, start).Filter("ArticleType__ID", TypeID).All(&articles)
	}

	var aType []models.ArticleType
	// redis 在就获取，不再就存储
	// conn, err := redis.Dial("tcp", ":6379")
	// if err != nil {
	// 	beego.Info("redis连接失败")
	// }
	// defer conn.Close()
	// reply, err := conn.Do("set", "atype", aType)
	// if err != nil {
	// 	beego.Info("redis存储失败")
	// }
	// beego.Info(reply)

	// 序列化与发序列化
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		beego.Info("redis连接失败")
	}
	defer conn.Close()
	ok, _ := redis.Bool(conn.Do("EXISTS", "atype"))
	beego.Info(ok)
	if ok == true {
		ReadBuffer, _ := redis.Bytes(conn.Do("get", "atype"))
		beego.Info(ReadBuffer)
		dec := gob.NewDecoder(bytes.NewReader(ReadBuffer))
		err = dec.Decode(&aType)
		if err != nil {
			beego.Info("获取不到解码后的数据:", err)
		}
		beego.Info(aType)
	} else {
		// 查询
		o.QueryTable("ArticleType").All(&aType)

		// redis存储序列化的数据
		var buffer bytes.Buffer
		enc := gob.NewEncoder(&buffer)
		err = enc.Encode(aType)
		reply, err := conn.Do("set", "atype", buffer.Bytes())
		if err != nil {
			beego.Info("redis存储失败")
		}
		beego.Info(reply)
	}

	c.Data["articles"] = articles
	c.Data["count"] = count
	c.Data["pageNum"] = pageNum
	c.Data["pageIndex"] = pageIndex
	c.Data["aType"] = aType
	c.Data["TypeID"] = TypeID

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
	o := orm.NewOrm()
	var aType []models.ArticleType
	o.QueryTable("ArticleType").All(&aType)
	c.Data["aType"] = aType
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
			var ArticleType models.ArticleType
			TypeID := c.GetString("TypeID")
			ArticleType.ID, _ = strconv.Atoi(TypeID)
			o.Read(&ArticleType)

			Arti.ArticleType = &ArticleType
			_, err = o.Insert(&Arti)
			if err != nil {
				beego.Info("插入数据失败", err)
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
	err = o.QueryTable("Article").RelatedSel("ArticleType").One(&article)
	if err != nil {
		beego.Info("获取数据错误")
	}
	beego.Info("article", article)
	// 点击数更新
	article.Acount++
	_, err = o.Update(&article)
	if err != nil {
		beego.Info("更新点击数错误")
	}
	// 多对多插入，user
	// 1 获取多对多操对象
	m2m := o.QueryM2M(&article, "Users")
	// 2 获取插入对象
	username := c.GetSession("username").(string)
	beego.Info(username)
	var user = models.User{UserName: username}
	o.Read(&user, "UserName")
	// 3 多对多插入
	_, err = m2m.Add(&user)
	if err != nil {
		beego.Info("插入多对多失败")
	}
	// 多对多查询
	// o.QueryTable("Article").Filter("Users__User__UserName", username).Distinct().Filter("ID", id).One(&article)
	// err = o.Read(&article)

	beego.Info(username)
	o.LoadRelated(&article, "Users")

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
	} else {
		Arti.Aimg = c.GetString("img_old")
	}
	_, err = o.Update(&Arti, "Artiname", "Acontent", "Aimg")
	if err != nil {
		beego.Info("更新数据失败", err)
		return
	}
	c.Redirect("/article_list", 302)

}

// HandleArticleDel 处理文章删除
func (c *ArticleController) HandleArticleDel() {
	id, err := c.GetInt("id")
	if err != nil {
		beego.Info("获取参数错误", err)
	}
	o := orm.NewOrm()
	arti := models.Article{ID: id}
	err = o.Read(&arti)
	if err != nil {
		beego.Info("获取数据错误")
	}
	_, err = o.Delete(&arti)
	if err != nil {
		beego.Info("删除数据失败", err)
		return
	}
	c.Redirect("/article_list", 302)

}
