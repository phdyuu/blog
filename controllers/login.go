package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

var globalSessions *session.Manager

type LoginController struct {
	beego.Controller
}

func (this *LoginController) ShowAdminIndex() {
	this.TplName = "admin/login.html"
}


func (this *LoginController) HandleLogin() {
	User := this.GetString("user")
	Passwd := this.GetString("passwd")
	if User != "phdyu" || Passwd != "gotoburnout" {
		this.Redirect("/AdminSystem", 302)
	} else {
		this.SetSession("Admin", 1)
		this.Redirect("/ShowAdminSystem", 302)
	}
}


func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:"adminPhdYu",
		EnableSetCookie: true,
		Gclifetime:60,
		Maxlifetime: 60,
		Secure: false,
		CookieLifeTime: 30,
		ProviderConfig: "./tmp",
	}
	globalSessions, _ = session.NewManager("memory",sessionConfig)
	go globalSessions.GC()
}