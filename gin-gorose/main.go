package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gohouse/gorose/v2"
	_ "github.com/mattn/go-sqlite3"
)

// 增加
func UserAdd(c *gin.Context) {
	username := c.Query("username")
	age := c.DefaultQuery("age", "0")
	var data = make(map[string]interface{})
	data["username"] = username
	data["age"] = age
	affected_rows, err := DB().Table("users").Data(data).Insert() // 执行入库
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(affected_rows)) // api接口返回
}

// 删除
func UserDelete(c *gin.Context) {
	uid := c.Query("uid") // 主键
	affected_rows, err := DB().Table("users").Where("uid", uid).Delete()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(affected_rows))
}

// 修改
func UserEdit(c *gin.Context) {
	uid := c.Query("uid")
	username := c.DefaultQuery("username", "")
	age := c.DefaultQuery("age", "0")
	var data = make(map[string]interface{})
	data["username"] = username
	data["age"] = age
	affected_rows, err := DB().Table("users").Where("uid", uid).Data(data).Update() // 执行入库
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(affected_rows))
}

// 查询
func UserList(c *gin.Context) {
	UserList, err := DB().Table("users").OrderBy("uid desc").Limit(50).Get()
	if err != nil {
		c.JSON(http.StatusOK, FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, SuccessReturn(UserList))
}

// 初始化
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
	affected_rows, err := DB().Execute(dbSql)
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
	router.Use(Cors())
	router.GET("/UserAdd", UserAdd)
	router.GET("/UserList", UserList)
	router.GET("/UserEdit", UserEdit)
	router.GET("/UserDelete", UserDelete)
	router.Static("/html", "./") // 静态文件服务
	router.Run()
}

// /
// /
// 工具函数
// 返回正确函数
func SuccessReturn(msg interface{}) map[string]interface{} {
	var res = make(map[string]interface{})
	res["data"] = msg
	res["code"] = http.StatusOK
	res["msg"] = "success"
	return res
}

// 返回错误函数
func FailReturn(msg interface{}) map[string]interface{} {
	var res = make(map[string]interface{})
	res["data"] = ""
	res["code"] = http.StatusBadRequest
	res["msg"] = msg

	return res
}

// DB orm快捷使用函数
func DB() gorose.IOrm {
	return engin.NewOrm()
}

// 处理跨域请求, 支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" { //放行所有OPTIONS方法
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next() // 处理请求
	}
}

func main() {
	InitGorose()
	InitUser()
	InitGin()
}
