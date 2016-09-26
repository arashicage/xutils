package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

type RedisLoadController struct {
	beego.Controller
}

func (this *RedisLoadController) Get() {
	var fpxx struct {
		A1 string `redis:"1"`
		A2 string `redis:"2"`
		A3 string `redis:"3"`
		A4 string `redis:"4"`
		A5 string `redis:"5"`
		A6 string `redis:"6"`
		A7 string `redis:"7"`
	}

	rs, err := redis.Dial("tcp", "172.30.11.231:22121")
	defer rs.Close()
	if err != nil {
		fmt.Println("ERROR create connection to redis:", err)
	}

	val, err := redis.Values(rs.Do("HGETALL", "01:110012414002835213"))
	if err != nil {
		fmt.Println("ERROR create connection to redis:", err)
	}

	if err := redis.ScanStruct(val, &fpxx); err != nil {
		fmt.Println("scan err", err)
	}

	fmt.Printf("%+v\n", fpxx)
	this.Data["data"] = fpxx

	this.Data["title"] = "redisload - X-SYSTEM"
	this.TplNames = "redis.html"
}
