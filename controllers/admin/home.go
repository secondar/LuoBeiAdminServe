package admin

import beego "github.com/beego/beego/v2/server/web"

type AdminHomeController struct {
	beego.Controller
}

func (_this *AdminHomeController) AdminHome() {
	_this.TplName = "admin/index.tpl"
}
