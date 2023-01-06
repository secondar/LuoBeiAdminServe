package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models"
	"LuoBeiAdminServeForGolang/service"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"math"
)

type ArticleController struct {
	beego.Controller
}

// 添加
func (_this *ArticleController) Add() {
	ResultJson := utils.ResultJson{}
	ArticleInfo := models.Article{}
	err := _this.ParseForm(&ArticleInfo)
	if err != nil {
		ResultJson.Code = 500
		ResultJson.Msg = "请求数据接收失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	ArticleService := service.ArticleService{}
	ArticleService.Init()
	_, err = ArticleService.Add(ArticleInfo)
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

// 编辑
func (_this *ArticleController) Edit() {
	ResultJson := utils.ResultJson{}
	ArticleInfo := models.Article{}
	err := _this.ParseForm(&ArticleInfo)
	if err != nil {
		ResultJson.Code = 500
		ResultJson.Msg = "请求数据接收失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	ArticleService := service.ArticleService{}
	ArticleService.Init()
	_, err = ArticleService.Edit(ArticleInfo)
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

// 删除
func (_this *ArticleController) Delete() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求，没有传递ID"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	ArticleService := service.ArticleService{}
	ArticleService.Init()
	_, err = ArticleService.Delete(id)
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

// 列表
func (_this *ArticleController) GetList() {
	ResultJson := utils.ResultJson{}
	key := utils.TrimSpace(_this.GetString("key"))
	page, err := _this.GetInt("page")
	if err != nil {
		page = 0
	} else {
		page = page - 1
	}
	limit, err := _this.GetInt("limit")
	if err != nil {
		limit = 30
	}
	sort, _ := _this.GetInt("sort")
	ArticleModel := models.ArticleModel{}
	ArticleModel.NewArticleQs()
	var ArticleList []models.Article
	qs := ArticleModel.Qs
	if key != "" {
		qs = qs.Filter("title__icontains", key)
	}
	if sort > 0 {
		qs = qs.Filter("sort", sort)
	}
	row, err := qs.OrderBy("id").Limit(limit, page).All(&ArticleList)
	if err != nil {
		logs.Error(err)
		ResultJson.Code = 503
		ResultJson.Msg = "获取失败，如果您是系统管理员，您可以通过错误日志查看详细错误信息"
	} else {
		type ResData struct {
			Total   int64            `json:"total"`
			Last    int64            `json:"last"`
			Current int              `json:"current"`
			List    []models.Article `json:"list"`
		}
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = ResData{Total: row, Last: int64(math.Ceil(float64(row) / float64(limit))), Current: page + 1, List: ArticleList}
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}

// 详情
func (_this *ArticleController) Details() {
	ResultJson := utils.ResultJson{}
	id, err := _this.GetInt("id")
	if err != nil {
		ResultJson.Code = 401
		ResultJson.Msg = "非法请求，没有传递ID"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	ArticleService := service.ArticleService{}
	ArticleService.Init()
	ArticleInfo, err := ArticleService.Details(id)
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = err.Error()
	} else {
		ResultJson.Code = 200
		ResultJson.Msg = "成功"
		ResultJson.Data = ArticleInfo
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
