package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"phdyu/models"
)

type MainController struct {
	beego.Controller
}


func (this *MainController) ShowIndex() {
	var moments[] models.Moment
	o := orm.NewOrm()
	qs := o.QueryTable("moment")
	qs.OrderBy("-id").Limit(4).All(&moments)
	this.Data["moments"] = moments
	this.TplName = "index.html"
}



