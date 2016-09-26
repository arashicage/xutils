package controllers

import (
	"fmt"
	"github.com/juju/errgo/errors"
	//"github.com/tgulacsi/goracle/examples/connect"
	"github.com/tgulacsi/goracle/oracle"
	"log"
	"time"
)

type ColConverter func(interface{}) string

type Column struct {
	Name   string
	String ColConverter
}

func GetColumns(cu *oracle.Cursor) (cols []Column, err error) {
	desc, err := cu.GetDescription()
	if err != nil {
		return nil, errors.Newf("error getting description for %s: %s", cu, err)
	}
	//log.Printf("columns: %s", columns)
	//log.Printf("desc: %#v", desc)

	//cols := godrv.ColumnDescriber(rows).DescribeColumns()
	//log.Printf("cols: %s", cols)
	//for rows.Next() {
	var ok bool
	cols = make([]Column, len(desc))
	for i, col := range desc {
		cols[i].Name = col.Name
		if cols[i].String, ok = converters[col.Type]; !ok {
			log.Fatalf("no converter for type %d (column name: %s)", col.Type, col.Name)
		}
	}
	return cols, nil
}

var converters = map[int]ColConverter{
	1: func(data interface{}) string { //VARCHAR2
		return fmt.Sprintf("%s", data.(string))
	},
	6: func(data interface{}) string { //NUMBER
		return fmt.Sprintf("%v", data)
	},
	96: func(data interface{}) string { //CHAR
		return fmt.Sprintf("%s", data.(string))
	},
	156: func(data interface{}) string { //DATE
		// return data.(time.Time).Format(time.RFC3339)
		//return `"` + data.(time.Time).String()[:19] + `"`
		return data.(time.Time).String()[:19]
	},
}

func SubStr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0
	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length
	if start > end {
		start, end = end, start
	}
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
