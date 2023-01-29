package main

import (
	"github.com/gohouse/gorose/v2"
	_ "github.com/mattn/go-sqlite3"
)

var engin *gorose.Engin

// DB orm快捷使用函数
func DB() gorose.IOrm {
	return engin.NewOrm()
}
