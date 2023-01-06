package service

import (
	"LuoBeiAdminServeForGolang/models"
	"errors"
	"github.com/astaxie/beego/logs"
)

type SystemService struct {
	Models models.SystemModel
}

func (s *SystemService) Init() {
	s.Models.NewSystemQs()
}

// 获取
func (s *SystemService) GetSystem() (models.System, error) {
	if s.Models.Orm == nil {
		s.Models.NewSystemOrm()
	}
	SystemInfo := models.System{Id: 1}
	err := s.Models.Orm.Read(&SystemInfo)
	if err != nil {
		logs.Error(err)
		return SystemInfo, errors.New("读取系统配置失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	} else {
		return SystemInfo, err
	}
}

// 修改
func (s *SystemService) SaveSystem(SystemInfo models.System) (int64, error) {
	if s.Models.Orm == nil {
		s.Models.NewSystemOrm()
	}
	SystemInfoUpdate := models.System{Id: 1}
	err := s.Models.Orm.Read(&SystemInfoUpdate)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("读取系统配置失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	SystemInfoUpdate.Title = SystemInfo.Title
	SystemInfoUpdate.Tail = SystemInfo.Tail
	SystemInfoUpdate.Keyword = SystemInfo.Keyword
	SystemInfoUpdate.Describe = SystemInfo.Describe
	row, err := s.Models.Orm.Update(&SystemInfoUpdate)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("保存网站配置时失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return row, err
}
