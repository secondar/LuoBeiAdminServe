package models

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminRouter struct {
	Id      int      `orm:"pk;auto;size(11)" json:"id"`
	Role    int      `orm:"size(11)" json:"role"`
	Menu    int      `orm:"size(11)" json:"menu"`
	IsPid   int8     `orm:"size(1)" json:"is_pid"`
	Addtime lib.Time `orm:"auto_now_add" json:"addtime"`
}
type AdminRouterModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(AdminRouter))
}

// 获取ORM
func (_this *AdminRouterModel) NewAdminRouterOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminRouterModel) NewAdminRouterQs() {
	if _this.Orm == nil {
		_this.NewAdminRouterOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(AdminRouter))
	}
}