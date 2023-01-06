package service

import (
	"LuoBeiAdminServeForGolang/models"
	"errors"
	"github.com/astaxie/beego/logs"
)

type ArticleSortService struct {
	Models models.ArticleSortModel
}

func (s *ArticleSortService) Init() {
	s.Models.NewArticleSortQs()
}

// 添加
func (s *ArticleSortService) Add(ArticleSortInfo models.ArticleSort) (int64, error) {
	err := ArticleSortInfo.ValidAdd()
	if err != nil {
		return 0, err
	}
	if s.Models.Qs == nil {
		s.Models.NewArticleSortOrm()
	}
	if ArticleSortInfo.Pid != 0 {
		CheckPid, err := s.Models.Qs.Filter("id", ArticleSortInfo.Pid).Count()
		if err == nil && CheckPid <= 0 {
			return 0, errors.New("父级分类不存在")
		}
	}
	row, err := s.Models.Orm.Insert(&ArticleSortInfo)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("添加文章分类在写入数据库时出现错误，如果您是系统管理员，您可以通过错误日志查看详细错误信息")
	}
	return row, nil
}

// 编辑
func (s *ArticleSortService) Edit(ArticleSortInfo models.ArticleSort) (int64, error) {
	err := ArticleSortInfo.ValidEdit()
	if err != nil {
		return 0, err
	}
	if s.Models.Qs == nil {
		s.Models.NewArticleSortOrm()
	}
	if ArticleSortInfo.Pid != 0 {
		CheckPid, err := s.Models.Qs.Filter("id", ArticleSortInfo.Pid).Count()
		if err == nil && CheckPid <= 0 {
			return 0, errors.New("父级分类不存在")
		}
	}
	row, err := s.Models.Orm.Update(&ArticleSortInfo)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("编辑文章分类在写入数据库时出现错误，如果您是系统管理员，您可以通过错误日志查看详细错误信息")
	}
	return row, nil
}

// 删除
func (s *ArticleSortService) Delete(Id int) (int64, error) {
	if s.Models.Qs == nil {
		s.Models.NewArticleSortOrm()
	}
	cou, err := s.Models.Qs.Filter("pid", Id).Count()
	if err == nil && cou > 0 {
		return 0, errors.New("删除失败，该分类下还有子分类，如果需要删除，请先删除该分类下的所有子分类")
	}
	row, err := s.Models.Qs.Filter("id", Id).Delete()
	if err != nil {
		return row, errors.New("删除失败，如果您的系统管理员，您可以查看错误日志")
	} else {
		return row, err
	}
}

// 获取列表
func (s *ArticleSortService) GetList() []models.ArticleSort {
	if s.Models.Qs == nil {
		s.Models.NewArticleSortQs()
	}
	ArticleSort := []models.ArticleSort{}
	s.Models.Qs.OrderBy("sort").All(&ArticleSort)
	return s.ToTree(ArticleSort, 0)
}

// 转树形菜单
func (s *ArticleSortService) ToTree(list []models.ArticleSort, ParentId int) []models.ArticleSort {
	tree := []models.ArticleSort{}
	for _, item := range list {
		if item.Pid == ParentId {
			item.Children = s.ToTree(list, item.Id)
			tree = append(tree, item)
		}
	}
	return tree
}
