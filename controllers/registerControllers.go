package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterControllers struct {
	beego.Controller
}
//直接登入接口
func (r *RegisterControllers) Get()  {
		r.TplName = "login.html"
	}


func (r *RegisterControllers) Post() {
	//fmt.Println("也执行了")
	//1、解析请求数据
	var user models.User
	err :=r.ParseForm(&user)
	if err!=nil {
		//返回错误信息给浏览器,提示用户
		r.Ctx.WriteString("对不起，数据解析错误")
		return
	}
	//2、保存用户信息到数据库
	_,err1 :=user.SeveUser()
	//3、返回前端结果(成功跳登录页面，失败弹出错误信息）
	if err1!=nil {
		fmt.Println(err1)
		r.Ctx.WriteString("对不起，用户注册失败")
		return
	}
	//用户注册成功
	r.TplName="login.html"
}
