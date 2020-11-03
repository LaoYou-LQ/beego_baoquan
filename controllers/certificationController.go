package controllers

import "github.com/astaxie/beego"

type CertificationController struct {
	beego.Controller
}

func (c *CertificationController) Post() {
	c.TplName="certification.html"
}