package models

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// 我对这个写法也不是很满意，只不过现在技能树不够，先这么写着吧
type AdminOnLine struct {
	Id             int      `orm:"pk;auto;size(11)" json:"id"`
	Aid            int      `orm:"size(11)" json:"aid"`
	Account        string   `orm:"type(char);size(32)" json:"account"`
	Token          string   `orm:"size(500)" json:"token"`
	ExpirationTime lib.Time `json:"expiration_time"`
}

type AdminOnLineModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(table_prefix, new(AdminOnLine))
}

// 获取ORM
func (_this *AdminOnLineModel) NewAdminOnLineOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *AdminOnLineModel) NewAdminOnLineQs() {
	if _this.Orm == nil {
		_this.NewAdminOnLineOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(AdminOnLine))
	}
}
