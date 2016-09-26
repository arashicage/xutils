package controllers

import (
	"strconv"
	"strings"
	//"bufio"
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/logs"
	"github.com/juju/errgo/errors"
	"github.com/tgulacsi/goracle/examples/connect"
	//"github.com/tgulacsi/goracle/oracle"
	"io"
	"log"
	//"os"
	//"path"
	//"strings"
	// "strings"
	//"strconv"
	//"time"
)

type CrudController struct {
	beego.Controller
}

type Recx struct {
	Gfsbh  string
	Gfmc   string
	Gfdzdh string
	Gfyhzh string
	Id     string
}

func getRecx(dsn, qry, del string) (r []Recx, e error) {

	r = make([]Recx, 100)

	cx, err := connect.GetRawConnection(dsn)
	if err != nil {
		return nil, errors.Newf("error connecting to database: %s", err)
	}
	defer cx.Close()

	cu := cx.NewCursor()
	err = cu.Execute(qry, nil, nil)
	if err != nil {
		return nil, errors.Newf("error executing %q: %s", qry, err)
	}
	defer cu.Close()

	columns, err := GetColumns(cu)
	if err != nil {
		return nil, errors.Newf("error getting column converters: %s", err)
	}
	log.Printf("columns: %#v", columns)

	n := 0
	rows, err := cu.FetchMany(1000)
	for err == nil && len(rows) > 0 {
		for j, row := range rows { //j 行
			for i, data := range row { //i 列
				if data == nil {
					data = ""
				}
				fmt.Println(j, i, columns[i].String(data))

				if i == 0 {
					r[j].Gfsbh = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
					//r[j].Gfsbh = columns[i].String(data)
				} else if i == 1 {
					r[j].Gfmc = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
					//r[j].Gfmc = columns[i].String(data)
				} else if i == 2 {
					r[j].Gfdzdh = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
					//r[j].Gfdzdh = columns[i].String(data)
				} else if i == 3 {
					r[j].Gfyhzh = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
					//r[j].Gfyhzh = columns[i].String(data)
				}
				r[j].Id = strconv.Itoa(j + 1)
				//r[j].Gfyhzh = columns[i].String(data)

			}
			n++
		}
		rows, err = cu.FetchMany(1000)
	}
	log.Printf("written %d rows.", n)
	if err != nil && err != io.EOF {
		return nil, errors.Newf("error fetching rows from %s: %s", cu, err)
	}
	return r, nil
}

func getRecx2(dsn, qry, del string) (rt []Recx, e error) {

	r := make([]Recx, 0)

	cx, err := connect.GetRawConnection(dsn)
	if err != nil {
		return nil, errors.Newf("error connecting to database: %s", err)
	}
	defer cx.Close()

	cu := cx.NewCursor()
	err = cu.Execute(qry, nil, nil)
	if err != nil {
		return nil, errors.Newf("error executing %q: %s", qry, err)
	}
	defer cu.Close()

	columns, err := GetColumns(cu)
	if err != nil {
		return nil, errors.Newf("error getting column converters: %s", err)
	}
	log.Printf("columns: %#v", columns)

	n := 0
	rows, err := cu.FetchMany(1000)
	for err == nil && len(rows) > 0 {
		for j, row := range rows { //j 行
			temp := Recx{}
			for i, data := range row { //i 列
				if data == nil {
					data = ""
				}
				fmt.Println(j, i, columns[i].String(data))

				if i == 0 {
					//r[j].Gfsbh = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
					temp.Gfsbh = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
				} else if i == 1 {
					//r[j].Gfmc = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
					temp.Gfmc = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
				} else if i == 2 {
					//r[j].Gfdzdh = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
					temp.Gfdzdh = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
				} else if i == 3 {
					//r[j].Gfyhzh = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
					temp.Gfyhzh = strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`)
				}
				//r[j].Id = strconv.Itoa(j + 1)
				temp.Id = strconv.Itoa(j + 1)

			}
			n++
			r = append(r, temp)
		}
		rows, err = cu.FetchMany(1000)
	}
	log.Printf("written %d rows.", n)
	if err != nil && err != io.EOF {
		return nil, errors.Newf("error fetching rows from %s: %s", cu, err)
	}
	//rt = r[0:len(r)]
	//fmt.Println(len(r))
	//fmt.Println(cap(r))
	//fmt.Println(rt)
	fmt.Println(len(r), cap(r))
	return r, nil
}

func (this *CrudController) Get() {

	dsn := "dzdz/oracle@DZDZ_dev_11200"
	qry := "SELECT gfsbh,gfmc,gfdzdh,gfyhzh FROM dzdz_fpxx_zzsfp where rownum<=100"
	del := "≡"

	x, _ := getRecx2(dsn, qry, del)

	this.Data["e"] = x

	this.Data["title"] = "CRUD - X-SYSTEM"
	this.TplNames = "crud.html"

}
