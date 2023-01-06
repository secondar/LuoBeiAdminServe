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

type Admin struct {
	Id        int      `orm:"pk;auto;size(11)" json:"id" form:"id"`
	Account   string   `orm:"type(char);size(32)" json:"account"  form:"account"`
	Password  string   `orm:"type(char);size(32)" json:"-"  form:"password"`
	Interfere string   `orm:"type(char);size(32)" json:"-"  form:"interfere"`
	Role      int      `orm:"size(11)" json:"role"  form:"role"`
	State     int8     `orm:"default(1);size(1)" json:"state"  form:"state"`
	Addtime   lib.Time `orm:"auto_now_add" json:"addtime"`
	Token     string   `orm:"-" json:"token"`
}

func (s Admin) ValidAdd() error {
	valid := validation.Validation{}
	valid.Required(utils.TrimSpace(s.Account), "账户").Message("必须填写")
	valid.Required(utils.TrimSpace(s.Password), "密码").Message("必须填写")
	valid.Required(s.Role, "权限组").Message("必须填写")
	valid.Min(s.Role, 1, "权限组").Message("必须大于等于1")
	valid.Range(s.State, 0, 2, "状态").Message("必须在0~1之间")
	return s.validation(&valid)
}
func (s Admin) ValidEdit() error {
	valid := validation.Validation{}
	valid.Min(s.Id, 1, "账户ID").Message("必须大于等于1")
	valid.Required(utils.TrimSpace(s.Account), "账户").Message("必须填写")
	valid.Required(s.Role, "权限组").Message("必须填写")
	valid.Min(s.Role, 1, "权限组").Message("必须大于等于1")
	valid.Range(s.State, 0, 2, "状态").Message("必须在0~1之间")
	return s.validation(&valid)
}

// 验证
func (s Admin) validation(valid *validation.Validation) error {
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

type AdminModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(Admin))
}

// 获取ORM
func (_this *AdminModel) NewAdminOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminModel) NewAdminQs() {
	if _this.Orm == nil {
		_this.NewAdminOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(Admin))
	}
}
