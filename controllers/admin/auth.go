package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/service"
	beego "github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	beego.Controller
}

var Captcha = utils.Captcha{}

func init() {
	Captcha.Init(1024, 3)
}
func (_this *AuthController) Login() {
	var ResultJson = utils.ResultJson{}
	Account := _this.GetString("account")
	Password := _this.GetString("password")
	CaptchaId := _this.GetString("captcha_id")
	CaptchaVal := _this.GetString("captcha_val")
	if utils.TrimSpace(CaptchaId) == "" || utils.TrimSpace(CaptchaVal) == "" {
		ResultJson.Code = 501
		ResultJson.Msg = "请填写验证码"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if !Captcha.VerifyCaptcha(CaptchaId, CaptchaVal, true) {
		ResultJson.Code = 501
		ResultJson.Msg = "验证码错误"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}

	Account = utils.TrimSpace(Account)
	Password = utils.TrimSpace(Password)
	if Account == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写用户名"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	if Password == "" {
		ResultJson.Code = 401
		ResultJson.Msg = "请填写用户密码"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	AdminService := service.AdminService{}
	AdminService.Init()
	AdminInfo, err := AdminService.Login(Account, Password, _this.Ctx.Input.IP())
	if err != nil {
		ResultJson.Code = 501
		ResultJson.Msg = err.Error()
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	ResultJson.Code = 200
	ResultJson.Msg = "登录成功"
	ResultJson.Data = AdminInfo
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
func (_this *AuthController) GetCaptcha() {
	var ResultJson = utils.ResultJson{}
	Captcha.UseMathConfig(50, 150)
	Result, err := Captcha.Generate()
	if err != nil {
		ResultJson.Code = 503
		ResultJson.Msg = "验证码生成失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	ResultJson.Code = 200
	ResultJson.Msg = "success"
	ResultJson.Data = Result
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
	return
}
