package home

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/models"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type HomeController struct {
	beego.Controller
}

// 首页
func (_this *HomeController) Home() {
	_this.Data["IsHome"] = "true"
	_this.TplName = "home/index.tpl"
}

// 文章
func (_this *HomeController) Article() {
	Title := "全部动态"
	tmp := _this.Ctx.Input.Param(":sort")
	Sort, err := strconv.Atoi(tmp)
	if err != nil {
		Sort = 0
	}
	tmp = _this.Ctx.Input.Param(":page")
	Page, err := strconv.Atoi(tmp)
	if err != nil {
		Page = 1
	}

	var ArticleList []models.Article
	var ArticleSortList []models.ArticleSort
	ArticleModel := models.ArticleModel{}
	ArticleSortModel := models.ArticleSortModel{}
	ArticleModel.NewArticleQs()
	ArticleSortModel.NewArticleSortQs()
	// 先获得分类
	_, err = ArticleSortModel.Qs.Filter("state", 1).All(&ArticleSortList)
	if err != nil {
		_this.Ctx.WriteString("获取文章分类失败")
		return
	}
	// 获得文章列表
	ArticleModelQS := ArticleModel.Qs
	if Sort > 0 {
		ArticleModelQS.Filter("sort", Sort)
		for _, item := range ArticleSortList {
			if item.Id == Sort {
				Title = item.Title
			}
		}
	}
	row, err := ArticleModelQS.OrderBy("-id").Count()
	if err != nil {
		_this.Ctx.WriteString("获取文章总数失败")
		return
	}
	_, err = ArticleModelQS.OrderBy("-id").Limit(30, (Page-1)*1).All(&ArticleList)
	if err != nil {
		_this.Ctx.WriteString("获取文章列表失败")
		return
	}
	_this.Data["ArticleList"] = ArticleList
	_this.Data["SortList"] = ArticleSortList
	_this.Data["Sort"] = Sort
	_this.Data["NextPage"] = Page + 1
	_this.Data["Title"] = Title
	_this.Data["Page"] = utils.ResPageInfo{
		Count: row,
		Limit: 30,
		Curr:  Page,
	}
	_this.TplName = "home/article.tpl"
}

// 详情
func (_this *HomeController) ArticleDetails() {
	tmp := _this.Ctx.Input.Param(":id")
	Id, err := strconv.Atoi(tmp)
	if err != nil {
		_this.Redirect("/404.html", 302)
		return
	}
	var Article models.Article
	ArticleModel := models.ArticleModel{}
	ArticleModel.NewArticleOrm()
	Article.Id = Id
	err = ArticleModel.Orm.Read(&Article)
	if err != nil {
		_this.Ctx.WriteString("文章获取失败")
		return
	}
	_this.Data["Info"] = Article
	_this.Data["Describe"] = Article.Describe
	_this.Data["Title"] = Article.Title
	_this.Data["InfoAddtime"] = Article.Addtime.GetStr()
	_this.TplName = "home/article_details.tpl"
}
