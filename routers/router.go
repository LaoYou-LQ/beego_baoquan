package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//注册页面
    beego.Router("/", &controllers.MainController{})
    //用户注册接口请求（从注册页面注册后转登入页面）
    beego.Router("/user_register",&controllers.RegisterControllers{})
    //直接登入的接口请求（从注册页面跳转至登入页面）
    beego.Router("/login",&controllers.RegisterControllers{})
	//用户登入请求接口，（登入页面转主页面）
    beego.Router("/home",&controllers.RequControllers{})
    //文件上传接口
    beego.Router("/renzheng",&controllers.RenZhengControllers{})
	//在认证数据列表页面，点击新增认证按钮，跳转"新增页面"
    beego.Router("/upload_file.html",&controllers.RenZhengControllers{})
	//查看认证数据的证书
	beego.Router("/cert_detail.html", &controllers.CertDetailController{})
	beego.Router("user_kyc.html",&controllers.CertificationController{})
    //用户实名认证接口
    beego.Router("/user_kyc",&controllers.CertificationController{})

}
