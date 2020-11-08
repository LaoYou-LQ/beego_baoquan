package controllers

import (
	"DataCertProject/blockchain"
	"DataCertProject/models"
	"DataCertProject/util"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController) Get() {
	certId := c.GetString("cert_id")
	fmt.Println(certId)
	block, err := blockchain.CHAIN.QueryBlockByCertId([]byte(certId))
	if err != nil {
		fmt.Println("123234344")
		c.Ctx.WriteString("链上数据查询错误")
		return
	}
	if block == nil {
		c.Ctx.WriteString("未查询到链上数据")
		return
	}
	certRecord, err := models.DeSerializeRecord(block.Data)
	certRecord.CertHashStr = string(certRecord.CertHash)
	certRecord.CertIdStr = strings.ToUpper(string(certRecord.CertId))
	certRecord.CertTimeFormat =util.TimeFormat(certRecord.CertTime,0,util.TIME_FORMAT_ONE)
	c.Data["CertRecord"] = certRecord
	c.TplName = "cert_detail.html"
}
