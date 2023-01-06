package service

import (
	"LuoBeiAdminServeForGolang/models"
	"github.com/beego/beego/v2/client/orm"
)

type AdminOnLineService struct {
	Models models.AdminOnLineModel
}

func (s *AdminOnLineService) Init() {
	s.Models.NewAdminOnLineQs()
}

//添加
func (s *AdminOnLineService) Add(AdminOnLineData models.AdminOnLine) (int64, error) {
	if s.Models.Orm == nil {
		s.Models.Orm = orm.NewOrm()
	}
	return s.Models.Orm.Insert(&AdminOnLineData)
}
