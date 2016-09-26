package controllers

import (
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/juju/errgo/errors"
	"github.com/tgulacsi/goracle/examples/connect"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Expdata2Controller struct {
	beego.Controller
}

func (this *Expdata2Controller) Get() {
	this.Data["title"] = "expdata - X-SYSTEM"
	this.TplNames = "expsql.html"

}

func expsql(w io.Writer, dsn, tab, commit string) error {
	commitN, _ := strconv.Atoi(commit)
	qry := "select * from " + tab
	cols := "insert into " + tab + " ("

	cx, err := connect.GetRawConnection(dsn)
	defer cx.Close()
	if err != nil {
		return errors.Newf("error connecting to database: %s", err)
	}

	cu := cx.NewCursor()
	defer cu.Close()
	err = cu.Execute(qry, nil, nil)
	if err != nil {
		return errors.Newf("error executing %q: %s", qry, err)
	}

	columns, err := GetColumns(cu)
	if err != nil {
		return errors.Newf("error getting column converters: %s", err)
	}
	log.Printf("columns: %#v", columns)

	bw := bufio.NewWriterSize(w, 65536)
	defer bw.Flush()

	for i, col := range columns {
		if i > 0 {
			cols = cols + ","
		}
		cols = cols + col.Name
	}
	cols = cols + ") values ("

	bw.WriteString("\nprompt 加载表 " + tab + "\n")

	n := 0
	rows, err := cu.FetchMany(1000)
	if len(rows) == 0 {
		bw.WriteString("prompt 加载 0 行")
	}

	for err == nil && len(rows) > 0 {
		for j, row := range rows {
			r := ""
			for i, data := range row {
				if i > 0 {
					r = r + `,`
				}
				if data == nil {
					r = r + `NULL`
				} else {
					r = r + `'` + strings.TrimSuffix(strings.TrimPrefix(columns[i].String(data), `"`), `"`) + `'`
				}

			}
			r = r + `);`
			if j == 1 {
				fmt.Println("\n*************\n", cols+r)
			}

			bw.WriteString(cols + r)
			bw.Write([]byte{'\n'})
			n++
			b := n % commitN //取余数
			if b == 0 {
				bw.WriteString("prompt commit at " + strconv.Itoa(n) + " 行")
				bw.WriteString("\ncommit;\n")
			}
		}
		bw.WriteString("prompt 加载 " + strconv.Itoa(n) + " 行\n")
		rows, err = cu.FetchMany(1000)
	}

	bw.WriteString("prompt commit at " + strconv.Itoa(n) + " 行")
	bw.WriteString("\ncommit;\n")

	log.Printf("written %d rows.", n)
	if err != nil && err != io.EOF {
		return errors.Newf("error fetching rows from %s: %s", cu, err)
	}
	return nil
}

func (this *Expdata2Controller) Post() {
	this.Data["title"] = "导出 - X-Utils"

	nls := this.Input().Get("nls")
	dsn := this.Input().Get("dsn")
	tablist := this.Input().Get("tablist")
	dir := this.Input().Get("dir")
	commit := this.Input().Get("commit")
	mode := this.Input().Get("mode")

	last := strings.LastIndex(dir, `\`)
	pdir := SubStr(dir, 0, last) //获取父目录

	this.Ctx.WriteString(`<table border="1">
							<tr>
							<td> nls </td>
							<td>` + nls + `</td>
							</tr>
							<tr>
							<td> commit </td>
							<td>` + commit + `</td>
							</tr>
							<tr>
							<td> mode </td>
							<td>` + mode + `</td>
							</tr>
							<tr>
							<td> dsn </td>
							<td>` + dsn + `</td>
							</tr>
							<tr>
							<td> tablist </td>
							<td>` + tablist + `</td>
							</tr>
							<tr>
							<td> dir </td>
							<td>` + dir + `</td>
							</tr>
							<tr>
							<td> pdir </td>
							<td>` + pdir + `</td>
							</tr>
							<tr>
							<td> pos </td>
							<td>` + strconv.Itoa(last) + `</td>
							</tr>
							<tr>
							<td> pdir-replaced </td>
							<td>` + strings.Replace(pdir, `\`, `\\`, -1) + `</td>
							</tr>							
						  </table>`)
	this.Ctx.WriteString(`<br/>`)
	this.Ctx.WriteString(`<div> nls:` + nls + `</div>`)
	this.Ctx.WriteString(`<div> commit:` + commit + `</div>`)
	this.Ctx.WriteString(`<div> mode:` + mode + `</div>`)
	this.Ctx.WriteString(`<div> dsn:` + dsn + `</div>`)
	this.Ctx.WriteString(`<div> tablist:` + tablist + `</div>`)
	this.Ctx.WriteString(`<div> dir:` + dir + `</div>`)
	this.Ctx.WriteString(`<div> pdir:` + pdir + `</div>`)
	this.Ctx.WriteString(`<div> idx:` + strconv.Itoa(last) + `</div>`)

	pdir = strings.Replace(pdir, `\`, `\\`, -1)
	this.Ctx.WriteString(`<div> pdir-replaced:` + pdir + `</div>`)

	//递归创建父目录
	err := os.MkdirAll(pdir, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println("Create Directory OK!")
	}

	//文件不存在创建文件，文件存在覆写该文件
	file, err := os.OpenFile(dir, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("writer", err)
	}
	defer file.Close()

	file.WriteString("prompt 2014年12月17日")
	tabs := strings.Split(tablist, `,`)
	for _, tab := range tabs {
		//fmt.Println("--------------", tab)
		//fmt.Println(">>>>>>>>>>>>>>", strings.TrimSpace(tab))
		if err := expsql(file, dsn, strings.TrimSpace(tab), commit); err != nil {
			log.Printf("error dumping: %s", err)
			this.Ctx.WriteString(err.Error())
		}
	}

	this.Ctx.WriteString(`<br/><a href="http://localhost:8080/expdata">back</a>`)

}
