package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/middleware/jwt_admin"
	"LuoBeiAdminServeForGolang/models"
	"LuoBeiAdminServeForGolang/service"
	beego "github.com/beego/beego/v2/server/web"
)

type MenuController struct {
	beego.Controller
}

// 添加菜单
func (_this *MenuController) Add() {
	ResultJson := utils.ResultJson{}
	AdminMenuData := models.AdminMenu{}
	err := _this.ParseForm(&AdminMenuData)
	if err != nil {
		ResultJson.Code = 500
		ResultJson.Msg = "请求数据接收失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminMenuService := service.AdminMenuService{}
	AdminMenuService.Init()
	_, err = AdminMenuService.Add(AdminMenuData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 编辑
func (_this *MenuController) Edit() {
	ResultJson := utils.ResultJson{}
	AdminMenuData := models.AdminMenu{}
	err := _this.ParseForm(&AdminMenuData)
	if err != nil {
		ResultJson.Code = 500
		ResultJson.Msg = "请求数据接收失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminMenuService := service.AdminMenuService{}
	AdminMenuService.Init()
	_, err = AdminMenuService.Edit(AdminMenuData)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 删除
func (_this *MenuController) Delete() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil || id <= 0 {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminMenuService := service.AdminMenuService{}
	AdminMenuService.Init()
	_, err = AdminMenuService.Delete(id)
	ResultJson.Code = 200
	ResultJson.Msg = "删除成功"
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 获取列表
func (_this *MenuController) List() {
	AdminMenuService := service.AdminMenuService{}
	AdminMenuService.Init()
	AdminMenu := AdminMenuService.GetList()
	ResultJson := utils.ResultJson{}
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	ResultJson.Data = AdminMenu
	_this.Data["json"] = ResultJson
	_this.ServeJSON()
}

// 获取账户拥有的路由
func (_this *MenuController) GetAdminMenuRouter() {
	ResultJson := utils.ResultJson{}
	AdminMenuService := service.AdminMenuService{}
	AdminMenuService.Init()
	AdminService := service.AdminService{}
	AdminInfo, err := AdminService.CtxTokenGetAdminInfo(_this.Ctx.Input.GetData("admin_token_claims").(*jwt_admin.CustomClaims))
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
	}
	AdminMenu := []models.AdminMenu{}
	if AdminInfo.Role != 1 {
		AdminMenu, err = AdminMenuService.GetAdminMenuRouter(AdminInfo.Id)
	} else {
		// 如果是超级管理员组就直接放行了
		AdminMenu = AdminMenuService.GetList()
	}
	if err == nil {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = AdminMenu
	} else {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	ResultJson.Data = AdminMenu
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
