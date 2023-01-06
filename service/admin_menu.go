package service

import (
	"LuoBeiAdminServeForGolang/models"
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type AdminMenuService struct {
	Models models.AdminMenuModel
}

func (s *AdminMenuService) Init() {
	s.Models.NewAdminMenuQs()
}
func (s *AdminMenuService) ValidData(AdminMenuData models.AdminMenu) error {
	if AdminMenuData.Type == 2 && AdminMenuData.Link < 0 || AdminMenuData.Type == 2 && AdminMenuData.Link > 1 {
		return errors.New("请选择是否为外链")
	}
	if AdminMenuData.Type == 3 || AdminMenuData.Sort != nil && *AdminMenuData.Sort < 0 && AdminMenuData.Type == 3 {
		return errors.New("请填写正确的排序")
	}
	if AdminMenuData.Type != 3 && AdminMenuData.Router != nil && *AdminMenuData.Router == "" {
		return errors.New("请填写路由地址")
	}
	if AdminMenuData.Type == 2 && AdminMenuData.Link == 0 {
		if AdminMenuData.Component != nil && *AdminMenuData.Component == "" {
			return errors.New("请填写组件名称")
		}
		if AdminMenuData.Path != nil && *AdminMenuData.Path == "" {
			return errors.New("请填写组件路径")
		}
	}
	return nil
}

// 添加菜单
func (s *AdminMenuService) Add(AdminMenuData models.AdminMenu) (int64, error) {
	err := AdminMenuData.ValidAdd()
	if err != nil {
		return 0, err
	}
	err = s.ValidData(AdminMenuData)
	if err != nil {
		return 0, err
	}
	if s.Models.Qs == nil {
		s.Models.NewAdminMenuQs()
	}
	AdminMenuDataCheck := models.AdminMenu{}
	if AdminMenuData.Pid != 0 {
		AdminMenuDataCheck.Id = AdminMenuData.Pid
		err := s.Models.Orm.Read(&AdminMenuDataCheck)
		if err != nil || AdminMenuDataCheck.Title == "" {
			return 0, errors.New("上级菜单不存在")
		}
	}
	i, err := s.Models.Orm.Insert(&AdminMenuData)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("添加失败，如果您是系统管理员可以查看错误日志")
	}
	return i, nil
}

// 编辑菜单
func (s *AdminMenuService) Edit(AdminMenuData models.AdminMenu) (int64, error) {
	err := AdminMenuData.ValidEdit()
	if err != nil {
		return 0, err
	}
	err = s.ValidData(AdminMenuData)
	if err != nil {
		return 0, err
	}
	if s.Models.Qs == nil {
		s.Models.NewAdminMenuQs()
	}
	AdminMenuDataCheck := models.AdminMenu{}
	if AdminMenuData.Pid != 0 {
		AdminMenuDataCheck.Id = AdminMenuData.Pid
		err := s.Models.Orm.Read(&AdminMenuDataCheck)
		if err != nil || AdminMenuDataCheck.Title == "" {
			return 0, errors.New("上级菜单不存在")
		}
	}
	tmp := models.AdminMenu{}
	tmp.Id = AdminMenuData.Id
	err = s.Models.Orm.Read(&tmp)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("找不到需要编辑的菜单ID")
	}
	AdminMenuData.Id = tmp.Id
	i, err := s.Models.Orm.Update(&AdminMenuData)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("修改失败，如果您是系统管理员可以查看错误日志")
	}
	return i, nil
}

// 删除
func (s *AdminMenuService) Delete(Id int) (int64, error) {
	if s.Models.Qs == nil {
		s.Models.NewAdminMenuQs()
	}
	cou, err := s.Models.Qs.Filter("pid", Id).Count()
	if err == nil && cou > 0 {
		return 0, errors.New("删除失败，该菜单下还有子菜单，如果需要删除，请先删除该菜单下的所有子菜单")
	}
	row, err := s.Models.Qs.Filter("id", Id).Delete()
	if err != nil {
		return row, errors.New("删除失败，如果您的系统管理员，您可以查看错误日志")
	} else {
		return row, err
	}
}

// 获取列表
func (s *AdminMenuService) GetList() []models.AdminMenu {
	AdminMenu := []models.AdminMenu{}
	s.Models.NewAdminMenuQs()
	s.Models.Qs.OrderBy("sort").All(&AdminMenu)
	return s.ToTree(AdminMenu, 0)
}

// 获取用户权限路由
func (s *AdminMenuService) GetAdminMenuRouter(role int) ([]models.AdminMenu, error) {
	if s.Models.Orm == nil {
		s.Models.Orm = orm.NewOrm()
	}
	table_prefix, err := beego.AppConfig.String("mysql::tableprefix")
	AdminMenu := []models.AdminMenu{}
	if err != nil {
		logs.Error(err)
	}
	Tsql := fmt.Sprintf("SELECT admin_menu.* FROM %sadmin_menu admin_menu RIGHT JOIN %sadmin_router admin_router ON admin_router.menu=admin_menu.id WHERE admin_router.role = ?", table_prefix, table_prefix)
	_, err = s.Models.Orm.Raw(Tsql, role).QueryRows(&AdminMenu)
	if err != nil {
		return AdminMenu, errors.New("获取用户权限路由时出现错误，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return s.ToTree(AdminMenu, 0), nil
}

// 转树形菜单
func (s *AdminMenuService) ToTree(list []models.AdminMenu, ParentId int) []models.AdminMenu {
	tree := []models.AdminMenu{}
	for _, item := range list {
		if item.Pid == ParentId {
			item.Children = s.ToTree(list, item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}
