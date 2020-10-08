/*
 * @Descripttion: NewGo
 * @version: go1.14.1 darwin/amd64
 * @Author: Newball
 * @Date: 2020-08-12 15:43:22
 * @LastEditors: Newball
 * @LastEditTime: 2020-08-19 21:52:31
 */
package server

import (
	"goFrame/api"
	"goFrame/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.New()

	// 中间件, 顺序不能改
	r.Use(gin.Recovery())
	r.Use(middleware.Logger()) //日志
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors()) //跨域
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
