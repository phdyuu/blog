package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}
func (c *ErrorController) Error401() {
	c.TplName="notfound.html"
}
func (c *ErrorController) Error403() {
	c.TplName="notfound.html"
}
func (c *ErrorController) Error404() {
	c.TplName="notfound.html"
}

func (c *ErrorController) Error500() {
	c.TplName="notfound.html"
}
func (c *ErrorController) Error503() {
	c.TplName="notfound.html"
}
