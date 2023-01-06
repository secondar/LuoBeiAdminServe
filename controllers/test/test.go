package test

import (
	"LuoBeiAdminServeForGolang/extend/utils"


	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type TestController struct {
	beego.Controller
}

func (_this *TestController) Test() {
	ArticleSortModel := models.ArticleSortModel{}
	a := ArticleSortModel.GetList()
	logs.Error(a)
	_this.Ctx.WriteString(utils.GetRandomKey(3))
}

func (_this *TestController) UserInfo() {
	type R struct {
		Roles        []string `json:"roles"`
		Introduction string   `json:"introduction"`
		Avatar       string   `json:"avatar"`
		Name         string   `json:"name"`
	}
	ResultJson := utils.ResultJson{}
	ResultJson.Code = 200
	ResultJson.Msg = "111"
	ResultJson.Data = R{
		Roles:        []string{"admin"},
		Introduction: "I am a super administrator",
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:         "Super Admin",
	}
	_this.Data["json"] = ResultJson
	_this.ServeJSON()
}
