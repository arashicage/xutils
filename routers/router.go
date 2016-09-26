package routers

import (
	"github.com/astaxie/beego"
	"xutils/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
