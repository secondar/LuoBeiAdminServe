package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/service"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type RouterController struct {
	beego.Controller
}

// 权限权限
func (_this *RouterController) Setting() {
	ResultJson := utils.ResultJson{}
	role, err := _this.GetInt("role")
	if err != nil || role <= 0 {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	router := utils.TrimSpace(_this.GetString("router"))
	if router == "" {
		ResultJson.Code = 402
		ResultJson.Msg = "请至少选择一个权限"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRouterService := service.AdminRouterService{}
	AdminRouterService.Init()
	_, err = AdminRouterService.SaveRouter(role, router)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 获取角色权限路由
func (_this *RouterController) GetRoleRouter() {
	ResultJson := utils.ResultJson{}
	role, err := _this.GetInt("role")
	logs.Error(role)
	if err != nil || role <= 0 {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRouterService := service.AdminRouterService{}
	AdminRouterService.Init()
	AdminRouterList, err := AdminRouterService.GetRoleRouter(role)
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	} else {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = AdminRouterList
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
