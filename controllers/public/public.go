package public

import beego "github.com/beego/beego/v2/server/web"

/**
  该控制器处理页面错误请求
*/
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {
	c.TplName = "404.tpl"
}
func (c *ErrorController) Error403() {
	c.TplName = "404.tpl"
}
func (c *ErrorController) Error404() {
	c.TplName = "404.tpl"
}

func (c *ErrorController) Error500() {
	c.TplName = "404.tpl"
}
func (c *ErrorController) Error503() {
	c.TplName = "404.tpl"
}
