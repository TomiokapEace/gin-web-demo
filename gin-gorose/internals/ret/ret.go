package main

import (
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

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
