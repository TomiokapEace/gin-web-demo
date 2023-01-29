package main

import (
	"gin-gorose/internals/init"

	_ "github.com/mattn/go-sqlite3"
)

// 初始化

func main() {
	init.InitGorose()
	init.InitUser()
	init.InitGin()
}
