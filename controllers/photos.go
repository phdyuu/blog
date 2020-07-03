package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"phdyu/models"
)

type PhotoController struct {
	beego.Controller
}

func (this *PhotoController) ShowPersonal() {
	var moments[] models.Moment
	o := orm.NewOrm()
	qs := o.QueryTable("moment")
	qs.Filter("class", "personal").OrderBy("-id").All(&moments)
	this.Data["moments"] = moments
	this.TplName = "personal.html"
}

func (this *PhotoController) ShowFriends() {
	var moments[] models.Moment
	o := orm.NewOrm()
	qs := o.QueryTable("moment")
	qs.Filter("class", "friends").OrderBy("-id").All(&moments)
	this.Data["moments"] = moments
	this.TplName = "friends.html"
}

func (this *PhotoController) ShowCats() {
	var moments[] models.Moment
	o := orm.NewOrm()
	qs := o.QueryTable("moment")
	qs.Filter("class", "cats").OrderBy("-id").All(&moments)
	this.Data["moments"] = moments
	this.TplName = "cats.html"
}

func (this *PhotoController) ShowScenery() {
	var moments[] models.Moment
	o := orm.NewOrm()
	qs := o.QueryTable("moment")
	qs.Filter("class", "scenery").OrderBy("-id").All(&moments)
	this.Data["moments"] = moments
	this.TplName = "scenery.html"
}

func (this *PhotoController) ShowPhotoClass() {
	var friends models.Moment
	var cats models.Moment
	var personal models.Moment
	var scenery models.Moment
	o := orm.NewOrm()
	qs := o.QueryTable("moment")
	qs.Filter("class", "friends").OrderBy("-id").Limit(1).All(&friends)
	qs.Filter("class", "cats").OrderBy("-id").Limit(1).All(&cats)
	qs.Filter("class", "personal").OrderBy("-id").Limit(1).All(&personal)
	qs.Filter("class", "scenery").OrderBy("-id").Limit(1).All(&scenery)
	this.Data["Friends"] = friends
	this.Data["Cats"] = cats
	this.Data["Personal"] = personal
	this.Data["Scenery"] = scenery
	this.TplName = "portfolio.html"

}