package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

DZDZ_CONF_SJCJ_FPLX

id_ number
fplx_dm varchar2(2)
fplx_desc   varchar2(200)
yxbz    char(1)


type Conf struct {
	ID             int64
	FPLX_DM        string
	FPLX_DESC      string
	YXBZ           string
}

func GetConf(id string) (*Conf, error) {
	tidnum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("Topic")
	err = qs.Filter("id", tidnum).One(topic)

	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}
