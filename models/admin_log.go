package models

import (
	"LuoBeiAdminServeForGolang/extend/lib"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminLog struct {
	Id      int      `orm:"pk;auto;size(11)" json:"id"`
	Aid     int      `orm:"size(11)" json:"aid"`
	Type    int8     `orm:"size(1)" json:"type"`
	Content string   `orm:"type(text);size(32)" json:"content"`
	Addtime lib.Time `orm:"auto_now_add" json:"addtime"`
}

type AdminLogModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(AdminLog))
}

// 获取ORM
func (_this *AdminLogModel) NewAdminLogOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminLogModel) NewAdminLogQs() {
	if _this.Orm == nil {
		_this.NewAdminLogOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(AdminLog))
	}
}
