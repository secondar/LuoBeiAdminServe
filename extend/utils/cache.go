package utils

import (
	"github.com/astaxie/beego/cache"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type Cache struct {
	cache.Cache
}

// 初始化缓存
func (s *Cache) InitCache() error {
	adapterName, err := beego.AppConfig.String("cache::adaptername")
	if err != nil {
		logs.Error(err)
		return err
	}
	config, err := beego.AppConfig.String("cache::config")
	if err != nil {
		logs.Error(err)
		return err
	}
	Cache, err := cache.NewCache(adapterName, config)
	if err != nil {
		logs.Error(err)
	}
	s.Cache = Cache
	return err
}
