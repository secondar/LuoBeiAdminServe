package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models"
	"LuoBeiAdminServeForGolang/service"
	beego "github.com/beego/beego/v2/server/web"
)

type SystemController struct {
	beego.Controller
}

// 获取
func (_this *SystemController) GetSystem() {
	ResultJson := utils.ResultJson{}
	SystemService := service.SystemService{}
	SystemService.Init()
	SystemInfo, err := SystemService.GetSystem()
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	} else {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = SystemInfo
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 保存
func (_this *SystemController) SaveSystem() {
	ResultJson := utils.ResultJson{}
	SystemInfo := models.System{}
	err := _this.ParseForm(&SystemInfo)
	if err != nil {
		ResultJson.Code = 500
		ResultJson.Msg = "请求数据接收失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	SystemService := service.SystemService{}
	SystemService.Init()
	_, err = SystemService.SaveSystem(SystemInfo)
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	} else {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = SystemInfo
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
