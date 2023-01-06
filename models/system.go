package models

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type System struct {
	Id       int    `orm:"pk;auto;size(11)" json:"id" form:"id"`
	Title    string `orm:"size(255)" json:"title"  form:"title"`
	Tail     string `orm:"size(255)" json:"tail"  form:"tail"`
	Keyword  string `orm:"size(255)" json:"keyword"  form:"keyword"`
	Describe string `orm:"size(255)" json:"describe"  form:"describe"`
}

func (s System) ValidAdd() error {
	valid := validation.Validation{}
	valid.Required(utils.TrimSpace(s.Title), "名称").Message("必须填写")
	valid.Required(utils.TrimSpace(s.Tail), "小尾巴").Message("必须填写")
	valid.Required(utils.TrimSpace(s.Keyword), "关键词").Message("必须填写")
	valid.Required(utils.TrimSpace(s.Describe), "描述").Message("必须填写")

	return s.validation(&valid)
}
func (s System) ValidEdit() error {
	valid := validation.Validation{}
	valid.Min(s.Id, 1, "角色ID").Message("必须大于等于1")
	valid.Required(utils.TrimSpace(s.Title), "名称").Message("必须填写")
	valid.Required(utils.TrimSpace(s.Tail), "小尾巴").Message("必须填写")
	valid.Required(utils.TrimSpace(s.Keyword), "关键词").Message("必须填写")
	valid.Required(utils.TrimSpace(s.Describe), "描述").Message("必须填写")
	return s.validation(&valid)
}

// 验证
func (s System) validation(valid *validation.Validation) error {
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

type SystemModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	tablePrefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(tablePrefix, new(System))
}

// 获取ORM
func (_this *SystemModel) NewSystemOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *SystemModel) NewSystemQs() {
	if _this.Orm == nil {
		_this.NewSystemOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(System))
	}
}
