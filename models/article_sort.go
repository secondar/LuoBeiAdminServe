package models

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"LuoBeiAdminServeForGolang/extend/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type ArticleSort struct {
	Id       int           `orm:"pk;auto;size(11)" json:"id" form:"id"`
	Pid      int           `orm:"size(11)" json:"pid" form:"pid"`
	Title    string        `orm:"size(32)" json:"title" form:"title"`
	State    int8          `orm:"size(1)" json:"state" form:"state"`
	Sort     int           `orm:"size(11)" json:"sort" form:"sort"`
	Addtime  lib.Time      `orm:"auto_now_add" json:"addtime"`
	Children []ArticleSort `orm:"-" json:"children"`
}

func (s ArticleSort) ValidAdd() error {
	valid := validation.Validation{}
	valid.Min(s.Pid, 0, "父级分类").Message("必须大于等于1")
	valid.Required(utils.TrimSpace(s.Title), "标题").Message("必须填写")
	valid.Required(s.State, "状态").Message("必须填写")
	valid.Required(s.Sort, "分类").Message("必须填写")
	valid.Min(s.Sort, 1, "分类").Message("必须大于等于1")
	return s.validation(&valid)
}
func (s ArticleSort) ValidEdit() error {
	valid := validation.Validation{}
	valid.Required(s.Id, "分类ID").Message("必须填写")
	valid.Min(s.Id, 1, "分类ID").Message("必须大于等于1")
	valid.Min(s.Pid, 0, "父级分类").Message("必须大于等于0")
	valid.Required(utils.TrimSpace(s.Title), "标题").Message("必须填写")
	valid.Range(s.State, 0, 2, "状态").Message("必须在0~1之间")
	valid.Required(s.Sort, "分类").Message("必须填写")
	valid.Min(s.Sort, 1, "分类").Message("必须大于等于1")
	return s.validation(&valid)
}

// 验证
func (s ArticleSort) validation(valid *validation.Validation) error {
	if valid.HasErrors() {
		errs := ""
		for _, err := range valid.Errors {
			if errs != "" {
				errs = fmt.Sprintf("%s，", errs)
			}
			errs = fmt.Sprintf("%s%s:%s", errs, err.Key, err.Message)
		}
		return errors.New(errs)
	}
	return nil
}

type ArticleSortModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(ArticleSort))
}

// 获取ORM
func (_this *ArticleSortModel) NewArticleSortOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *ArticleSortModel) NewArticleSortQs() {
	if _this.Orm == nil {
		_this.NewArticleSortOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(ArticleSort))
	}
}
