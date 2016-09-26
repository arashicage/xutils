package controllers

import (
	// "beeblog/models"
	"github.com/astaxie/beego"
)

type UtilsController struct {
	beego.Controller
}

func (this *UtilsController) Get() {
	this.Data["title"] = "工具 - X-SYSTEM"
	this.TplNames = "dump.html"

}
