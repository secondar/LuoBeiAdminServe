package service

import (
	"LuoBeiAdminServeForGolang/models"
	"github.com/beego/beego/v2/client/orm"
)

type AdminLogService struct {
	Models models.AdminLogModel
}

func (s *AdminLogService) Init() {
	s.Models.NewAdminLogQs()
}

//添加日志
func (s *AdminLogService) Add(LogData models.AdminLog) (int64, error) {
	if s.Models.Orm == nil {
		s.Models.Orm = orm.NewOrm()
	}
	return s.Models.Orm.Insert(&LogData)
}
