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

    beego.Router("/upload_file.html",&controllers.RenZhengControllers{})

   // beego.Router("/cert_detail.html",&controllers.)
}
