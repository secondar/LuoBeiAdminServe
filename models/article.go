package models

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"LuoBeiAdminServeForGolang/extend/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type Article struct {
	Id        int      `orm:"pk;auto;size(11)" json:"id" form:"id"`
	Sort      int      `orm:"size(11)" json:"sort" form:"sort"`
	Title     string   `orm:"size(255)" json:"title" form:"title"`
	Thumbnail *string  `orm:"size(255)" json:"thumbnail" form:"thumbnail"`
	Content   string   `orm:"type(text)" json:"content" form:"content"`
	Describe  string   `orm:"size(255)" json:"describe" form:"describe"`
	Hot       int      `orm:"size(11)" json:"hot" form:"sort"`
	Addtime   lib.Time `orm:"auto_now_add" json:"addtime"`
}

func (s Article) ValidAdd() error {
	valid := validation.Validation{}
	valid.Required(utils.TrimSpace(s.Title), "标题").Message("必须填写")
	valid.Required(utils.TrimSpace(s.Content), "内容").Message("必须填写")
	valid.Required(s.Sort, "分类").Message("必须填写")
	valid.Min(s.Sort, 1, "分类").Message("必须大于等于1")
	return s.validation(&valid)
}
func (s Article) ValidEdit() error {
	valid := validation.Validation{}
	valid.Required(utils.TrimSpace(s.Title), "标题").Message("必须填写")
	valid.Required(utils.TrimSpace(s.Content), "内容").Message("必须填写")
	valid.Required(s.Sort, "分类").Message("必须填写")
	valid.Min(s.Id, 1, "文章ID").Message("必须大于等于1")
	return s.validation(&valid)
}

// 验证
func (s Article) validation(valid *validation.Validation) error {
	if valid.HasErrors() {
		errs := ""
		for _, err := range valid.Errors {
			if errs != "" {
				errs = fmt.Sprintf("%s，", errs)
			}
			errs = fmt.Sprintf("%s%s:%s", errs, err.Key, err.Message)
		}
		return errors.New(errs)
	}
	return nil
}

type ArticleModel struct {
	Orm orm.Ormer
	Qs  orm.QuerySeter
}

// 初始化
func init() {
	tablePrefix, err := beego.AppConfig.String("mysql::tableprefix")
	if err != nil {
		logs.Error(err)
	}
	orm.RegisterModelWithPrefix(tablePrefix, new(Article))
}

// 获取ORM
func (_this *ArticleModel) NewArticleOrm() {
	if _this.Orm == nil {
		_this.Orm = orm.NewOrm()
	}
}

// 获取QS
func (_this *ArticleModel) NewArticleQs() {
	if _this.Orm == nil {
		_this.NewArticleOrm()
	}
	if _this.Qs == nil {
		_this.Qs = _this.Orm.QueryTable(new(Article))
	}
}