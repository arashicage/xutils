package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	//"log"
	//"strings"
)

type AboutController struct {
	beego.Controller
}

type Trace struct {
	Name, Desc, SQL string
}

func (this *AboutController) Get() {

	var trace []Trace

	bytes, err := ioutil.ReadFile("conf/trace.ini")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}

	//fmt.Println(string(bytes))

	err = json.Unmarshal(bytes, &trace)
	if err != nil {
		fmt.Println("error:", err)
	}

	this.Data["trace"] = trace

	for _, v := range trace {
		fmt.Println(v.Name)
		fmt.Println(v.Desc)
		//this.Ctx.WriteString(v.Name + v.Desc + v.SQL)
	}

	//this.Ctx.WriteString()

	this.Data["title"] = "关于 - X-SYSTEM"
	this.TplNames = "about.html"
}
