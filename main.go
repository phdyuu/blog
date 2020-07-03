package main

import (
	"phdyu/controllers"
	_ "phdyu/routers"
	"github.com/astaxie/beego"
	_"phdyu/models"
	_ "github.com/astaxie/beego/session/mysql"
)

func main() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

