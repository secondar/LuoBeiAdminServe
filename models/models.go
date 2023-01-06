package models

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// 初始化数据库连接
func InitDb() {
	user, err := beego.AppConfig.String("mysql::user")
	if err != nil {
		logs.Error("读取数据库用户名失败，请检查是否配置完成")
		os.Exit(1)
	}
	pass, err := beego.AppConfig.String("mysql::pass")
	if err != nil {
		logs.Error("读取数据库密码失败，请检查是否配置完成")
		os.Exit(1)
	}
	host, err := beego.AppConfig.String("mysql::host")
	if err != nil {
		logs.Error("读取数据库主机失败，请检查是否配置完成")
		os.Exit(1)
	}
	port, err := beego.AppConfig.String("mysql::port")
	if err != nil {
		logs.Error("读取数据库端口失败，请检查是否配置完成")
		os.Exit(1)
	}
	dbname, err := beego.AppConfig.String("mysql::dbname")
	if err != nil {
		logs.Error("读取数据库表名失败，请检查是否配置完成")
		os.Exit(1)
	}
	charset, err := beego.AppConfig.String("mysql::charset")

	// 设置db连接
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Error("数据库连接失败，请检查是否配置正确")
		os.Exit(1)
	}
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", user, pass, host, port, dbname, charset)
	err = orm.RegisterDataBase("default", "mysql", conn)
	if err != nil {
		logs.Error("数据库连接失败，请检查是否配置正确")
		os.Exit(1)
	}
	runmode, _ := beego.AppConfig.String("runmode")
	// 开发模式开启sql调试
	if runmode == "dev" {
		orm.Debug = true
	}
}
func CreateWhere(where map[string]string) []string {
	wh := make([]string, len(where))
	index := 0
	for _, item := range where {
		wh[index] = item
		index++
	}
	return wh
}
