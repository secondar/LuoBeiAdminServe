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

type AdminRole struct {
	Id      int      `orm:"pk;auto;size(11)" json:"id" form:"id"`
	Title   string   `orm:"size(32)" json:"title" form:"title"`
	Remarks *string  `orm:"size(255)" json:"remarks" form:"remarks"`
	State   int8     `orm:"size(1);default(1)" json:"state" form:"state"`
	Addtime lib.Time `orm:"auto_now_add" json:"addtime"`
}

func (s AdminRole) ValidAdd() error {
	valid := validation.Validation{}
	valid.Required(utils.TrimSpace(s.Title), "名称").Message("必须填写")
	valid.Range(s.State, 0, 2, "状态").Message("必须在0~1之间")
	return s.validation(&valid)
}
func (s AdminRole) ValidEdit() error {
	valid := validation.Validation{}
	valid.Min(s.Id, 1, "角色ID").Message("必须大于等于1")
	valid.Required(utils.TrimSpace(s.Title), "名称").Message("必须填写")
	valid.Range(s.State, 0, 2, "状态").Message("必须在0~1之间")
	return s.validation(&valid)
}

// 验证
func (s AdminRole) validation(valid *validation.Validation) error {
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

type AdminRoleModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(AdminRole))
}

// 获取ORM
func (_this *AdminRoleModel) NewAdminRoleOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminRoleModel) NewAdminRoleQs() {
	if _this.Orm == nil {
		_this.NewAdminRoleOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(AdminRole))
	}
}
