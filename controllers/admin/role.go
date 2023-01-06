package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models"
	"LuoBeiAdminServeForGolang/service"
	beego "github.com/beego/beego/v2/server/web"
)

type RoleController struct {
	beego.Controller
}

// 添加菜单
func (_this *RoleController) Add() {
	ResultJson := utils.ResultJson{}
	AdminRoleData := models.AdminRole{}
	err := _this.ParseForm(&AdminRoleData)
	if err != nil {
		ResultJson.Code = 500
		ResultJson.Msg = "请求数据接收失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRoleService := service.AdminRoleService{}
	AdminRoleService.Init()
	_, err = AdminRoleService.Add(AdminRoleData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 添加菜单
func (_this *RoleController) Edit() {
	ResultJson := utils.ResultJson{}
	AdminRoleData := models.AdminRole{}
	err := _this.ParseForm(&AdminRoleData)
	if err != nil {
		ResultJson.Code = 500
		ResultJson.Msg = "请求数据接收失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRoleService := service.AdminRoleService{}
	AdminRoleService.Init()
	_, err = AdminRoleService.Edit(AdminRoleData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 获取列表
func (_this *RoleController) List() {
	ResultJson := utils.ResultJson{}
	AdminRoleService := service.AdminRoleService{}
	AdminRoleList, err := AdminRoleService.GetList()
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = "失败"
	} else {
		ResultJson.Data = AdminRoleList
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 删除
func (_this *RoleController) Delete() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["josn"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminRoleService := service.AdminRoleService{}
	_, err = AdminRoleService.Delete(id)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 权限权限
func (_this *RoleController) SettingRouter() {
	ResultJson := utils.ResultJson{}
	role, err := _this.GetInt("role")
	if err != nil {
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
