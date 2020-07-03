package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"path"
	"phdyu/models"
	"time"
)


type AdminController struct {
	beego.Controller
}




func (this *AdminController) ShowAdminAlbum() {
	if !isLogin(this) {
		this.TplName = "notfound.html"
		return
	}

	var moments[] models.Moment
	var cont = 0
	o := orm.NewOrm()
	qs := o.QueryTable("moment")
	qs.OrderBy("-id").All(&moments)
	this.Data["moments"] = moments
	this.Data["cont"] = cont
	this.TplName = "admin/album.html"
}


func (this *AdminController) ShowAddMoment() {
	if !isLogin(this) {
		this.TplName = "notfound.html"
		return
	}
	this.TplName = "admin/addMoment.html"
}


func (this *AdminController) HandleAddMoment() {
	Name := this.GetString("articleName")
	Content := this.GetString("content")
	Class := this.GetString("pclass")
	f, h, _ := this.GetFile("uploadname")
	defer f.Close()
	ext := path.Ext(h.Filename)
	fileName := time.Now().Format("2006-01-02-15-04-05")
	_ = this.SaveToFile("uploadname", "./static/moments/"+fileName+ext)

	o := orm.NewOrm()
	moment := models.Moment{}
	moment.Title = Name
	moment.Content = Content
	moment.Class = Class
	moment.Img = "./static/moments/"+fileName+ext
	_, err := o.Insert(&moment)
	if err != nil {
		beego.Info(err)
	}
	this.TplName = "admin/addMoment.html"
}

func (this *AdminController) ShowUpdateMoment() {
	if !isLogin(this) {
		this.TplName = "notfound.html"
		return
	}
	uid := this.GetString("id")
	title := this.GetString("title")
	content := this.GetString("content")
	this.Data["uid"] = uid
	this.Data["title"] = title
	this.Data["content"] = content
	this.TplName = "admin/UpdateMoment.html"
}

func (this *AdminController) HandleUpdateMoment() {
	uid, _ := this.GetInt("id")
	title := this.GetString("articleName")
	o := orm.NewOrm()
	moment := models.Moment{Id:uid}
	err :=o.Read(&moment)
	if err != nil {
		this.Ctx.WriteString("Err")
		return
	}

	content := this.GetString("content")
	class := this.GetString("pclass")
	f, h, _ := this.GetFile("uploadname")
	ext := path.Ext(h.Filename)
	fileName := time.Now().Format("2006-01-02-15-04-05")
	_ = this.SaveToFile("uploadname", "./static/moments/"+fileName+ext)
	defer f.Close()
	moment.Img = "./static/moments/"+fileName+ext
	moment.Content = content
	moment.Title = title
	moment.Class = class
	_, _ = o.Update(&moment)
	this.Redirect("/ShowAdminAlbum", 302)
}

func (this *AdminController) HandleDeleteMoment() {
	uid, _ := this.GetInt("id")
	o := orm.NewOrm()
	moment := models.Moment{Id:uid}
	o.Delete(&moment)
	this.Redirect("/ShowAdminAlbum", 302)

}



func (this *AdminController) ShowAdminSystem() {
	if !isLogin(this) {
		this.TplName = "notfound.html"
		return
	}
	this.TplName = "admin/index.html"
}

func isLogin(this *AdminController) bool {
	uid := this.GetSession("Admin")
	if uid != 1 {
		return false
	}
	return true
}





