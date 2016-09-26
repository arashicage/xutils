package controllers

import (
	// "beeblog/models"
	//	"github.com/Unknwon/goconfig"
	"github.com/astaxie/beego"
	//	"log"
	"fmt"
)

type ConfigController struct {
	beego.Controller
}

func (this *ConfigController) Post() {
	nls := this.Input().Get("nls")
	tns := this.Input().Get("tns")
	dir := this.Input().Get("expdir")
	fmt.Println(">>>tns_amdin will be reconfig to ", tns)
	fmt.Println(">>>nls_lang  will be reconfig to  nls", nls)
	fmt.Println(">>>exp_dir   will be reconfig to  dir", dir)

	//	修改控制文件

	err := beego.AppConfig.Set("nls_lang", nls)
	if err != nil {
		fmt.Println(err)
	}
	err = beego.AppConfig.Set("tns_admin", tns)
	if err != nil {
		fmt.Println(err)
	}
	err = beego.AppConfig.Set("exp_dir", dir)
	if err != nil {
		fmt.Println(err)
	}

	beego.AppConfig.SaveConfigFile(beego.AppConfigPath)

	beego.ParseConfig()

	nls_lang := beego.AppConfig.String("nls_lang")
	tns_admin := beego.AppConfig.String("tns_admin")
	exp_dir := beego.AppConfig.String("exp_dir")
	var nls_zh bool

	if nls_lang == "SIMPLIFIED CHINESE_CHINA.ZHS16GBK" {
		nls_zh = true
	} else {
		nls_zh = false
	}

	this.Data["nls_lang"] = nls_lang
	this.Data["tns_admin"] = tns_admin
	this.Data["exp_dir"] = exp_dir
	this.Data["nls_checked"] = nls_zh

	this.Data["title"] = "系统配置 - X-SYSTEM"
	this.TplNames = "setting.html"
}

func (this *ConfigController) Get() {
	beego.ParseConfig()

	nls_lang := beego.AppConfig.String("nls_lang")
	tns_admin := beego.AppConfig.String("tns_admin")
	exp_dir := beego.AppConfig.String("exp_dir")
	var nls_zh bool

	if nls_lang == "SIMPLIFIED CHINESE_CHINA.ZHS16GBK" {
		nls_zh = true
	} else {
		nls_zh = false
	}

	this.Data["nls_lang"] = nls_lang
	this.Data["tns_admin"] = tns_admin
	this.Data["exp_dir"] = exp_dir
	this.Data["nls_checked"] = nls_zh

	this.Data["title"] = "系统配置 - X-SYSTEM"
	this.TplNames = "setting.html"

}
