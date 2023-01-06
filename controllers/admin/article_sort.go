package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models"
	"LuoBeiAdminServeForGolang/service"
	beego "github.com/beego/beego/v2/server/web"
)

type ArticleSortController struct {
	beego.Controller
}

// 获取列表
func (_this *ArticleSortController) GetList() {
	ResultJson := utils.ResultJson{}
	ArticleSortService := service.ArticleSortService{}
	ArticleSortService.Init()
	ArticleSortList := ArticleSortService.GetList()
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	ResultJson.Data = ArticleSortList
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 添加
func (_this *ArticleSortController) Add() {
	ResultJson := utils.ResultJson{}
	ArticleSortService := service.ArticleSortService{}
	ArticleSortService.Init()
	ArticleSort := models.ArticleSort{}
	err := _this.ParseForm(&ArticleSort)
	if err != nil {
		ResultJson.Code = 500
		ResultJson.Msg = "请求数据接收失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	_, err = ArticleSortService.Add(ArticleSort)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 修改
func (_this *ArticleSortController) Edit() {
	ResultJson := utils.ResultJson{}
	ArticleSortService := service.ArticleSortService{}
	ArticleSortService.Init()
	ArticleSort := models.ArticleSort{}
	err := _this.ParseForm(&ArticleSort)
	if err != nil {
		ResultJson.Code = 500
		ResultJson.Msg = "请求数据接收失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	_, err = ArticleSortService.Edit(ArticleSort)
	ResultJson.Code = 200
	ResultJson.Msg = "成功"
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

//删除
func (_this *ArticleSortController) Delete() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "获取不到分类ID"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	ArticleSortService := service.ArticleSortService{}
	ArticleSortService.Init()
	_, err = ArticleSortService.Delete(id)
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	} else {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
