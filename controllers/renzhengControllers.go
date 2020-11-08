package controllers

import (
	"DataCertProject/blockchain"
	"DataCertProject/models"
	"DataCertProject/util"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type RenZhengControllers struct {
	beego.Controller
}

func (z *RenZhengControllers) Get() {
	phone := z.GetString("phone") //key和html页面中的一样（不是模版语法中的）
	z.Data["Phone"] = phone
	z.TplName = "home.html"
}
/*
 *1、获取客户端上传的文件以及其他form表单的信息
 *关闭文件
 *2、将文件保存在本地的一个目录中
 *3、将上传的记录保存到数据库中
 *4、从数据库中读取phone用户对应的所有认证数据记录
 */



func (z *RenZhengControllers) Post() {

	//获取上传的文件
	//标题21

	fileTitle := z.Ctx.Request.PostFormValue("upload_title")
	phone := z.Ctx.Request.PostFormValue("phone")
	//文件
	f, h, err := z.GetFile("renzhen_file")
	if err != nil {
		z.Ctx.WriteString("用户数据解析失败")
		return
	}
	defer f.Close()
	//路径
	//2、将文件保存在本地的一个目录中
	//文件全路径： 路径 + 文件名 + "." + 扩展名
	//要的文件的路径
	uploadDir := "./static/img/" + h.Filename
	saveFile, err := os.OpenFile(uploadDir, os.O_RDWR|os.O_CREATE, 777)
	defer saveFile.Close()
	//创建一个writer: 用于向硬盘上写一个文件
	writer := bufio.NewWriter(saveFile)
	_, err = io.Copy(writer, f)
	if err != nil {
		//如果返回错误页面 r.tplName="xxx.html"
		z.Ctx.WriteString("保存数据出错")
		return
	}
	//打开文件，获取文件内容
	hashfile, err := os.Open(uploadDir)
	defer hashfile.Close()
	//将获取的内容加密
	hash, err := util.Md5HashReader(hashfile) //保全号加密：10.16上午
	t :=time.Now().Unix()
	//将上传的记录保存到数据库中
	record := models.UploadRecord{}
	record.FileName = h.Filename
	record.FileSize = h.Size
	record.FileTitle = fileTitle
	record.CertTime = t
	record.FileCert = hash
	record.Phone = phone
	//fmt.Println(record.FileCert)
	//fmt.Println("hash",hash)
	_, err = record.SeveRecord()
	if err != nil {
		fmt.Println(err)
		z.Ctx.WriteString("数据认证错误")
		return
	}
	//1.准备用户相关信息
	us,err:=models.QueryUserByPhone(phone)
	if err!=nil {
		z.Ctx.WriteString("数据认证失败")
		return
	}
	//2准备认证文件的SHA256哈希
	certhash,_:=util.SHA256HashReader(hashfile)
	//3实例化一个认证数据结构体实例
	certRecord:=models.CertRecord{
		CertId:[]byte(hash),
		CertHash:[]byte(certhash),
		CertAuthor:us.Name,
		AuthorCard:us.Card,
		Phone:us.Phone,
		FileName:h.Filename,
		FileSize:h.Size,
		CertTime:t,
	}
	//4序列化
	certBytes,err:=certRecord.SerializeRecord()
	_, err = blockchain.CHAIN.SaveData(certBytes)
	if err != nil {
		z.Ctx.WriteString("认证数据上链失败")
		return
	}
	records, err := models.QueryRecordbyPhone(phone)
	if err != nil {
		z.Ctx.WriteString("获取认证数据失败")
		return
	}
	//fmt.Println(records)
	z.Data["Records"] = records
	z.Data["Phone"] = phone
	z.TplName = "list_record.html"
}
