package controllers

import (
	"DataCertProject/blockchain"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController) Get() {
	certId:=c.GetString("cert_id")
	fmt.Println(certId)
	block,err:=blockchain.CHAIN.QueryBlockByCertId([]byte(certId))
	if err!=nil {
		fmt.Println("123234344")
		c.Ctx.WriteString("链上数据查询错误")
		return
	}
	if block==nil {
		c.Ctx.WriteString("未查询到链上数据")
		return
	}
	c.Data["CertId"]=strings.ToUpper(string(block.Data))
	c.TplName ="cert_detail.html"
}