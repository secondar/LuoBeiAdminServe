package service

import (
	"LuoBeiAdminServeForGolang/models"
	"errors"
	"github.com/beego/beego/v2/core/logs"
)

type AdminRoleService struct {
	Models models.AdminRoleModel
}

func (s *AdminRoleService) Init() {
	s.Models.NewAdminRoleQs()
}

// 添加角色
func (s *AdminRoleService) Add(AdminRoleData models.AdminRole) (int64, error) {
	err := AdminRoleData.ValidAdd()
	if err != nil {
		return 0, err
	}
	if s.Models.Qs == nil {
		s.Models.NewAdminRoleQs()
	}
	if AdminRoleData.Title == "" {
		return 0, errors.New("请填写角色名称")
	}
	AdminRoleDataCheck := models.AdminRole{}
	err = s.Models.Qs.Filter("title", AdminRoleData.Title).One(&AdminRoleDataCheck)
	if err == nil {
		return 0, errors.New("名称已存在")
	}
	row, err := s.Models.Orm.Insert(&AdminRoleData)
	if err != nil {
		logs.Error(err)
		return row, errors.New("添加失败，如果您是系统管理员，您可以通过错误日志查看错误信息")
	}
	return row, err
}

// 编辑角色
func (s *AdminRoleService) Edit(AdminRoleData models.AdminRole) (int64, error) {
	err := AdminRoleData.ValidEdit()
	if err != nil {
		return 0, err
	}
	if s.Models.Qs == nil {
		s.Models.NewAdminRoleQs()
	}
	if AdminRoleData.Title == "" {
		return 0, errors.New("请填写角色名称")
	}
	AdminRoleDataCheck := models.AdminRole{}
	AdminRoleDataCheck.Id = AdminRoleData.Id
	err = s.Models.Orm.Read(&AdminRoleDataCheck)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("找不到需要编辑的角色ID")
	}
	AdminRoleData.Id = AdminRoleDataCheck.Id
	row, err := s.Models.Orm.Update(&AdminRoleData)
	if err != nil {
		logs.Error(err)
		return row, errors.New("修改失败，如果您是系统管理员，您可以通过错误日志查看错误信息")
	}
	return row, err
}
func (s *AdminRoleService) Details(Id int) (models.AdminRole, error) {
	if s.Models.Orm == nil {
		s.Models.NewAdminRoleOrm()
	}
	AdminRole := models.AdminRole{Id: Id}
	err := s.Models.Orm.Read(&AdminRole)
	return AdminRole, err
}

// 获取列表
func (s *AdminRoleService) GetList() ([]models.AdminRole, error) {
	if s.Models.Qs == nil {
		s.Models.NewAdminRoleQs()
	}
	var RoleList []models.AdminRole
	_, err := s.Models.Qs.OrderBy("id").All(&RoleList)
	if err != nil {
		logs.Error(err)
		return RoleList, errors.New("获取失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return RoleList, err
}

// 删除
func (s *AdminRoleService) Delete(Id int) (int64, error) {
	if s.Models.Qs == nil {
		s.Models.NewAdminRoleQs()
	}
	row, err := s.Models.Qs.Filter("id", Id).Delete()
	if err != nil {
		logs.Error(err)
		return row, errors.New("删除失败，如果您的系统管理员，您可以查看错误日志")
	} else {
		return row, err
	}
}
