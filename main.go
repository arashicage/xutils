package main

import (
	"github.com/astaxie/beego"
	"xutils/controllers"
	_ "xutils/routers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/x", &controllers.MainController{}, "get:X")
	//beego.Router("/dump", &controllers.DumpController{})
	//beego.Router("/utils", &controllers.UtilsController{})

	beego.Router("/config", &controllers.ConfigController{})
	beego.Router("/utils/expcsv", &controllers.ExpCSVController{}) //文本文件导出
	beego.Router("/utils/expcsv/ws", &controllers.ExpCSVController{}, "get:Ws")
	beego.Router("/utils/expsql", &controllers.ExpSQLController{})
	beego.Router("/expdata2", &controllers.Expdata2Controller{})
	beego.Router("/crud", &controllers.CrudController{})
	beego.Router("/qrcode", &controllers.QRcodeController{})
	beego.Router("/routine", &controllers.RoutineController{})
	beego.Router("/redis/load", &controllers.RedisLoadController{})

	beego.Router("/about", &controllers.AboutController{})

	beego.Run()
}
