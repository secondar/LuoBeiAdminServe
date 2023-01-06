package service

import (
	"LuoBeiAdminServeForGolang/models"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/client/orm"
)

type ArticleService struct {
	Models models.ArticleModel
}

func (s *ArticleService) Init() {
	s.Models.NewArticleQs()
}

// 添加
func (s *ArticleService) Add(ArticleInfo models.Article) (int64, error) {
	err := ArticleInfo.ValidAdd()
	if err != nil {
		return 0, err
	}
	if s.Models.Orm == nil {
		s.Models.NewArticleOrm()
	}
	if ArticleInfo.Title == "" {
		return 0, errors.New("请填写文章标题")
	}
	if ArticleInfo.Sort <= 0 {
		return 0, errors.New("请选择文章分类")
	}
	if ArticleInfo.Content == "" {
		return 0, errors.New("请填写文章内容")
	}
	row, err := s.Models.Orm.Insert(&ArticleInfo)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("添加文章时出现错误，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return row, err
}

// 修改
func (s *ArticleService) Edit(ArticleInfo models.Article) (int64, error) {
	err := ArticleInfo.ValidEdit()
	if err != nil {
		return 0, err
	}
	if s.Models.Orm == nil {
		s.Models.NewArticleOrm()
	}
	ArticleInfoUpdate := models.Article{}
	ArticleInfoUpdate.Id = ArticleInfo.Id
	err = s.Models.Orm.Read(&ArticleInfoUpdate)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("查询不到需要修改的文章")
	}
	ArticleInfoUpdate.Sort = ArticleInfo.Sort
	ArticleInfoUpdate.Title = ArticleInfo.Title
	ArticleInfoUpdate.Content = ArticleInfo.Content
	ArticleInfoUpdate.Describe = ArticleInfo.Describe
	row, err := s.Models.Orm.Update(&ArticleInfoUpdate)
	if err != nil {
		logs.Error(err)
		return 0, errors.New("修改文章时失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return row, err
}

// 删除
func (s *ArticleService) Delete(id int) (int64, error) {
	if s.Models.Qs == nil {
		s.Models.NewArticleQs()
	}
	row, err := s.Models.Qs.Filter("id", id).Delete()
	if err != nil {
		logs.Error(err)
		return 0, errors.New("删除文章失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return row, err
}

// 获取详情
func (s *ArticleService) Details(id int) (models.Article, error) {
	if s.Models.Orm == nil {
		s.Models.NewArticleOrm()
	}
	ArticleInfo := models.Article{}
	ArticleInfo.Id = id
	err := s.Models.Orm.Read(&ArticleInfo)
	if err != nil {
		if err == orm.ErrNoRows {
			return ArticleInfo, errors.New("找不到文章信息")
		} else {
			logs.Error(err)
			return ArticleInfo, errors.New("获取文章信息失败，如果您是系统管理员，您可以通过错误日志查看错误信息")
		}
	}
	return ArticleInfo, err
}
