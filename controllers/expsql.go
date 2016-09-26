package controllers

import (
	"github.com/astaxie/beego"
)

type ExpSQLController struct {
	beego.Controller
}

func (this *ExpSQLController) Get() {
	this.Data["title"] = "expSQL - X-SYSTEM"
	this.TplNames = "expsql.html"

}
