package customemplatefunction

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models"
	"LuoBeiAdminServeForGolang/service"
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	"reflect"
	"time"
)

var Cache utils.Cache

func Init() {
	_ = beego.AddFuncMap("GetSystemConf", GetSystemConf)
	err := Cache.InitCache()
	if err != nil {
		logs.Error("初始化缓存失败,错误信息：%s", err.Error())
	}
}
func GetSystemConf() (SystemConfInfo models.System) {
	SystemConf := Cache.Get("SystemConf")
	if SystemConf == nil || reflect.TypeOf(SystemConf).String() != "models.System" {
		SystemModel := service.SystemService{}
		SystemData, err := SystemModel.GetSystem()
		if err != nil {
			logs.Error("从数据可获取网站配置失败,错误信息：%s", err.Error())
			return models.System{}
		}
		SystemConf = SystemData
		err = Cache.Put("SystemConf", SystemData, 12000*60*time.Second)
		if err != nil {
			logs.Error("写网站配置缓存失败,错误信息：%s", err.Error())
			return models.System{}
		}
	}
	return SystemConf.(models.System)
}
