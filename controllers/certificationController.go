package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CertificationController struct {
	beego.Controller
}

func (c CertificationController) Get() {
	c.TplName="user_kyc.html"
}

func (c *CertificationController) Post() {
	var user models.User
	err:=c.ParseForm(&user)
	if err!=nil {
		fmt.Println(err)
		c.Ctx.WriteString("认证失败")
	}
	_,err=user.Update()
	if err!=nil {
		fmt.Println(err)
		c.Ctx.WriteString("用户认证失败，请重试")
		return
	}
	records,err:=models.QueryRecordbyPhone(user.Phone)
	if err!=nil {
		c.Ctx.WriteString("抱歉，获取数据失败")
		return
	}
	c.Data["Records"] =records
	c.Data["Phone"]=user.Phone
	c.TplName ="list_record.html"
}