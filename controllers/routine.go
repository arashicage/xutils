package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type RoutineController struct {
	beego.Controller
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("error:", err)
	}
}

func (this *RoutineController) Get() {

	var trace []Trace

	bytes, err := ioutil.ReadFile("conf/trace.ini")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}

	err = json.Unmarshal(bytes, &trace)
	CheckError(err)

	this.Data["trace"] = trace

	for _, v := range trace {
		fmt.Println(v.Name)
		fmt.Println(v.Desc)
	}

	this.Data["title"] = "关于 - X-SYSTEM"
	this.Data["xx2"] = `[{"RRYID":"039","公共部分":"22.0440","设备管理":"0","班组管理":"0","运行管理":"-0.20","安全管理":"-0.10"},{"RRYID":"586","公共部分":"33.2670","设备管理":"0","班组管理":"0","运行管理":"-1.50","安全管理":"-0.20"},{"RRYID":"429","公共部分":"10.7290","设备管理":"0","班组管理":"0","运行管理":"-0.20","安全管理":"0"},{"RRYID":"372","公共部分":"54.4370","设备管理":"0","班组管理":"0","运行管理":"0","安全管理":"0"},{"RRYID":"061","公共部分":"29.7760","设备管理":"0","班组管理":"0","运行管理":"-0.20","安全管理":"-0.20"},{"RRYID":"008","公共部分":"19.1960","设备管理":"0","班组管理":"0","运行管理":"0","安全管理":"-0.50"},{"RRYID":"0363","公共部分":"27.6040","设备管理":"0","班组管理":"0","运行管理":"-1","安全管理":"0"},{"RRYID":"0521","公共部分":"38.0160","设备管理":"0","班组管理":"0","运行管理":"0","安全管理":"0"},{"RRYID":"0163","公共部分":"27.7390","设备管理":"0","班组管理":"0","运行管理":"0","安全管理":"-0.60"}]`
	this.TplNames = "routine.html"

}
