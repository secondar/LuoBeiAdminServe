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

// 注意，这个如果变动的话记得在/extend/jwt/jwt_admin/jwt.go 第五十四行一并更改
type AdminMenu struct {
	Id             int         `orm:"pk;auto;size(11)" json:"id" form:"id"`
	Pid            int         `orm:"size(11)" json:"pid" form:"pid"`
	Title          string      `orm:"size(255)" json:"title" form:"title"`
	Type           int8        `orm:"size(1)" json:"type" form:"type"`
	Icon           *string     `orm:"size(255)" json:"icon" form:"icon"`
	Show           int8        `orm:"size(1)" json:"show" form:"show"`
	Link           int8        `orm:"size(1)" json:"link" form:"link"`
	ApiPath        *string     `orm:"size(255)" json:"api_path" form:"api_path"`
	Characteristic *string     `orm:"size(255)" json:"characteristic" form:"characteristic"`
	Router         *string     `orm:"size(255)" json:"router" form:"router"`
	Sort           *int        `orm:"size(11)" json:"sort" form:"sort"`
	Component      *string     `orm:"size(255)" json:"component" form:"component"`
	Path           *string     `orm:"size(255)" json:"path" form:"path"`
	Addtime        lib.Time    `orm:"auto_now_add" json:"addtime"`
	Children       []AdminMenu `orm:"-" json:"children"`
}

func (s AdminMenu) ValidAdd() error {
	valid := validation.Validation{}
	valid.Min(s.Pid, 0, "父级菜单").Message("必须大于等于0")
	valid.Required(utils.TrimSpace(s.Title), "标题").Message("必须填写")
	valid.Range(s.Type, 1, 3, "类型").Message("必须在1~3之间")
	valid.Range(s.Show, 0, 1, "是否显示").Message("必须在0~1之间")
	valid.Range(s.Link, 0, 1, "是否外链").Message("必须在0~1之间")
	return s.validation(&valid)
}
func (s AdminMenu) ValidEdit() error {
	valid := validation.Validation{}
	valid.Min(s.Id, 1, "菜单ID").Message("必须大于等于1")
	valid.Min(s.Pid, 0, "父级菜单").Message("必须大于等于0")
	valid.Required(utils.TrimSpace(s.Title), "标题").Message("必须填写")
	valid.Range(s.Type, 1, 3, "类型").Message("必须在1~3之间")
	valid.Range(s.Show, 0, 1, "是否显示").Message("必须在0~1之间")
	valid.Range(s.Link, 0, 1, "是否外链").Message("必须在0~1之间")
	return s.validation(&valid)
}

// 验证
func (s AdminMenu) validation(valid *validation.Validation) error {
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

type AdminMenuModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(AdminMenu))
}

// 获取ORM
func (_this *AdminMenuModel) NewAdminMenuOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminMenuModel) NewAdminMenuQs() {
	if _this.Orm == nil {
		_this.NewAdminMenuOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(AdminMenu))
	}
}
