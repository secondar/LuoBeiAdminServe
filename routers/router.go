package routers

import (
	"LuoBeiAdminServeForGolang/controllers/admin"
	"LuoBeiAdminServeForGolang/controllers/home"
	"LuoBeiAdminServeForGolang/middleware/jwt_admin"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	InitHomeRouter()
	InitAdminRouter()
}
func InitAdminRouter() {
	beego.Router("/lb_secondar_admin", &admin.AdminHomeController{}, "GET:AdminHome")

	ns := beego.NewNamespace("/admin/api",
		beego.NSNamespace("/auth",
			beego.NSRouter("/login", &admin.AuthController{}, "POST:Login"),
			beego.NSRouter("/captcha", &admin.AuthController{}, "GET:GetCaptcha"),
		),
		beego.NSNamespace("/v1",
			beego.NSNamespace("/menu",
				beego.NSRouter("/add", &admin.MenuController{}, "POST:Add"),
				beego.NSRouter("/edit", &admin.MenuController{}, "POST:Edit"),
				beego.NSRouter("/delete", &admin.MenuController{}, "POST:Delete"),
				beego.NSRouter("/list", &admin.MenuController{}, "GET:List"),
				beego.NSRouter("/getadminmenurouter", &admin.MenuController{}, "GET:GetAdminMenuRouter"),
				//	/admin/api/v1/menu/getadminmenurouter 这个是基础，也就是获取菜单列表这个接口，每个账户必须有这个接口的访问权限
				//  如果业务需求需要更改，请将 /extend/jwt/jwt_admin/jwt.go 第四十二行一并更改
			),
			beego.NSNamespace("/role",
				beego.NSRouter("/add", &admin.RoleController{}, "POST:Add"),
				beego.NSRouter("/edit", &admin.RoleController{}, "POST:Edit"),
				beego.NSRouter("/delete", &admin.RoleController{}, "POST:Delete"),
				beego.NSRouter("/list", &admin.RoleController{}, "GET:List"),
				beego.NSRouter("/settingrouter", &admin.RoleController{}, "POST:SettingRouter"),
			),
			beego.NSNamespace("/router",
				beego.NSRouter("/setting", &admin.RouterController{}, "POST:Setting"),
				beego.NSRouter("/getrouter", &admin.RouterController{}, "POST:GetRoleRouter"),
			),
			beego.NSNamespace("/admin",
				beego.NSRouter("/getlist", &admin.AdminController{}, "POST:GetList"),
				beego.NSRouter("/add", &admin.AdminController{}, "POST:Add"),
				beego.NSRouter("/edit", &admin.AdminController{}, "POST:Edit"),
				beego.NSRouter("/delete", &admin.AdminController{}, "POST:Delete"),
			),
			beego.NSNamespace("/system",
				beego.NSRouter("/get", &admin.SystemController{}, "GET:GetSystem"),
				beego.NSRouter("/save", &admin.SystemController{}, "POST:SaveSystem"),
			),
			beego.NSNamespace("/article",
				beego.NSRouter("/getlist", &admin.ArticleController{}, "POST:GetList"),
				beego.NSRouter("/details", &admin.ArticleController{}, "GET:Details"),
				beego.NSRouter("/add", &admin.ArticleController{}, "POST:Add"),
				beego.NSRouter("/edit", &admin.ArticleController{}, "POST:Edit"),
				beego.NSRouter("/delete", &admin.ArticleController{}, "POST:Delete"),
				beego.NSNamespace("/sort",
					beego.NSRouter("/getlist", &admin.ArticleSortController{}, "POST:GetList"),
					beego.NSRouter("/add", &admin.ArticleSortController{}, "POST:Add"),
					beego.NSRouter("/edit", &admin.ArticleSortController{}, "POST:Edit"),
					beego.NSRouter("/delete", &admin.ArticleSortController{}, "POST:Delete"),
				),
			),
			beego.NSNamespace("/upload",
				beego.NSRouter("/img", &admin.UpFileController{}, "POST:UploadImg"),
			),
		),
	)
	beego.InsertFilterChain("/admin/api/v1/*", jwt_admin.JwtAuth) //这里设置鉴权
	beego.AddNamespace(ns)
}
func InitHomeRouter() {
	beego.Router("/", &home.HomeController{}, "GET:Home")
	beego.Router("/article/:sort-:page.html", &home.HomeController{}, "GET:Article")
	beego.Router("/article/details/:id.html", &home.HomeController{}, "GET:ArticleDetails")
}
