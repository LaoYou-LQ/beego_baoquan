package controllers

import (
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
	phone :=z.GetString("phone")//key和html页面中的一样（不是模版语法中的）
	z.Data["Phone"]=phone
	z.TplName="home.html"
}

func (z *RenZhengControllers) Post() {

	//获取上传的文件
	//标题21

	fileTitle := z.Ctx.Request.PostFormValue("upload_title")
	phone :=z.Ctx.Request.PostFormValue("Phone")
	//文件
	f, h, err := z.GetFile("renzhen_file")
	if err !=nil {
		z.Ctx.WriteString("用户数据解析失败")
		return
	}
	defer f.Close()
	//路径
	uploadDir:="./static/img/"+h.Filename
	saveFile , err :=os.OpenFile(uploadDir,os.O_RDWR|os.O_CREATE,777)
	defer saveFile.Close()
	writer:=bufio.NewWriter(saveFile)
	file_size ,err := io.Copy(writer,f)
	if err!=nil {
		//如果返回错误页面 r.tplName="xxx.html"
		z.Ctx.WriteString("保存数据出错")
		return
	}
	fmt.Println("拷贝的文件大小是",file_size)
	fmt.Println("文件标题是",fileTitle)
	hashfile,err:=os.Open(uploadDir)
	defer hashfile.Close()
	hash, err := util.Md5HashReader(hashfile)
	//将上传的记录保存到数据库中
	record := models.UploadRecord{}
	record.FileName = h.Filename
	record.FileSize = h.Size
	record.FileTitle = fileTitle
	record.CertTime = time.Now().Unix()
	record.FileCert = hash
	record.Phone =phone
	fmt.Println(record.FileCert)
	fmt.Println("hash",hash)
	_, err =record.SeveRecord()
	if err !=nil {
		fmt.Println(err)
		z.Ctx.WriteString("数据认证错误")
		return
	}
	records,err :=models.QueryRecordbyPhone(phone)
	if err!=nil {
		z.Ctx.WriteString("获取认证数据失败")
		return
	}
	fmt.Println(records)
	z.Data["Records"]=records
	z.TplName = "list_record.html"
}
