package controllers

import (
	"github.com/astaxie/beego"
)

type QRcodeController struct {
	beego.Controller
}

func (this *QRcodeController) Get() {
	this.Data["title"] = "QRcode - X-SYSTEM"
	this.TplNames = "qrcode.html"

}
