package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["title"] = "首页 - X-SYSTEM"
	this.TplNames = "index.html"

}

func (m *MainController) X() {
	m.Data["title"] = "首页 - X-SYSTEM"
	m.TplNames = "index.html"
}
