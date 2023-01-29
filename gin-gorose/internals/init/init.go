package init

import (
	"gopro/handler"
	"net/http"
	"sync"

	"gin-gorose/internals/cors"
	"gin-gorose/internals/db"

	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose/v2"
	_ "github.com/mattn/go-sqlite3"
)

var once sync.Once
var engin *gorose.Engin

func InitGorose() {
	var err error
	once.Do(func() {
		engin, err = gorose.Open(&gorose.Config{
			Driver: "sqlite3",
			Dsn:    "db.sqlite",
		})
		if err != nil {
			panic(err.Error())
		}
	})
}

func InitUser() {
	dbSql := `CREATE TABLE IF NOT EXISTS "users" ( "uid" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, "username" TEXT NOT NULL default "", "age" integer NOT NULL default 0)`
	affected_rows, err := db.DB().Execute(dbSql)
	if err != nil {
		panic(err.Error())
	}
	if affected_rows == 0 {
		return
	}
}

func InitGin() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK,
			`<meta http-equiv="refresh" content="3;URL=/html"><br><br>
				<h1>简单实现增删改查</h1>
				</center>`)
	})
	router.Use(cors.Cors())
	router.GET("/UserAdd", handler.UserAdd)
	router.GET("/UserList", handler.UserList)
	router.GET("/UserEdit", handler.UserEdit)
	router.GET("/UserDelete", handler.UserDelete)
	router.Static("/html", "./") // 静态文件服务
	router.Run()
}
