package handler

import (
	"net/http"

	"gin-gorose/internals/db"
	"gin-gorose/internals/ret"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// 增加
func UserAdd(c *gin.Context) {
	username := c.Query("username")
	age := c.DefaultQuery("age", "0")
	var data = make(map[string]interface{})
	data["username"] = username
	data["age"] = age
	affected_rows, err := db.DB().Table("users").Data(data).Insert() // 执行入库
	if err != nil {
		c.JSON(http.StatusOK, ret.FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, ret.SuccessReturn(affected_rows)) // api接口返回
}

// 删除
func UserDelete(c *gin.Context) {
	uid := c.Query("uid") // 主键
	affected_rows, err := db.DB().Table("users").Where("uid", uid).Delete()
	if err != nil {
		c.JSON(http.StatusOK, ret.FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, ret.SuccessReturn(affected_rows))
}

// 修改
func UserEdit(c *gin.Context) {
	uid := c.Query("uid")
	username := c.DefaultQuery("username", "")
	age := c.DefaultQuery("age", "0")
	var data = make(map[string]interface{})
	data["username"] = username
	data["age"] = age
	affected_rows, err := db.DB().Table("users").Where("uid", uid).Data(data).Update() // 执行入库
	if err != nil {
		c.JSON(http.StatusOK, ret.FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, ret.SuccessReturn(affected_rows))
}

// 查询
func UserList(c *gin.Context) {
	UserList, err := db.DB().Table("users").OrderBy("uid desc").Limit(50).Get()
	if err != nil {
		c.JSON(http.StatusOK, ret.FailReturn(err.Error()))
		return
	}
	c.JSON(http.StatusOK, ret.SuccessReturn(UserList))
}
