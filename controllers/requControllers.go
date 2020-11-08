package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type RequControllers struct {
	beego.Controller
}

func (m *RequControllers) Post() {
	var user models.User
	err := m.ParseForm(&user)
	if err != nil {
		m.Ctx.WriteString("登入失败")
		return
	}
	u, err := user.Querys()
	if err != nil {
		fmt.Println("23",err)
		m.Ctx.WriteString("账号或密码错误")
		return
	}
	name := strings.TrimSpace(u.Name)
	card := strings.TrimSpace(u.Card)
	sex := strings.TrimSpace(u.Sex)
	if name == "" || card == "" || sex == "" {
		//直接跳转到
		m.Data["Phone"]=u.Phone
		m.TplName="user_kyc.html"
		return
	}
	m.Data["Phone"] = u.Phone //html中的模板语法添加的就是Data中的"Phone"
	m.TplName = "home.html"
}
