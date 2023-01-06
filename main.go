package main

import (
	"LuoBeiAdminServeForGolang/controllers/public"
	"LuoBeiAdminServeForGolang/extend/customemplatefunction"
	"LuoBeiAdminServeForGolang/models"
	_ "LuoBeiAdminServeForGolang/routers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.InitDb()
	// 跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//其中Options跨域复杂请求预检
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "X-Token", "X-Requested-With", "withCredentials", "Accept", "Accept-Encoding", "Accept-Language", "Connection", "Host", "Referer", "Sec-Fetch-Dest", "Sec-Fetch-Mode", "Sec-Fetch-Site", "User-Agent"},
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	// 自定义模板函数
	customemplatefunction.Init()
	// session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
}
func main() {
	beego.ErrorController(&public.ErrorController{})
	beego.Run()
}
