package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type RequControllers struct {
	beego.Controller
}

func (m *RequControllers) Post() {
	var user models.User
	err :=m.ParseForm(&user)
	if err!=nil {
		m.Ctx.WriteString("登入失败")
		return
	}
	u,err:=user.Querys()
	if err!=nil {
		m.Ctx.WriteString("账号或密码错误")
		return
	}
	m.Data["Phone"]=u.Phone//html中的模板语法添加的就是Data中的"Phone"
	m.TplName="home.html"
}