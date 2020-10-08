/*
 * @Descripttion: NewGo
 * @version: go1.14.1 darwin/amd64
 * @Author: Newball
 * @Date: 2020-08-12 15:43:22
 * @LastEditors: Newball
 * @LastEditTime: 2020-08-19 22:19:42
 */
package main

import (
	"goFrame/conf"
	"goFrame/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()
	_ = r.Run(":8080")
}
