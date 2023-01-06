package service

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"LuoBeiAdminServeForGolang/extend/utils"
	"LuoBeiAdminServeForGolang/middleware/jwt_admin"
	"LuoBeiAdminServeForGolang/models"
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

type AdminService struct {
	Models models.AdminModel
}

func (s *AdminService) Init() {
	s.Models.NewAdminQs()
}

// 新增用户
func (s *AdminService) Add(AdminData models.Admin) (int64, error) {
	err := AdminData.ValidAdd()
	if err != nil {
		return 0, err
	}
	if s.Models.Qs == nil {
		s.Models.NewAdminQs()
	}
	AdminDataChek := models.Admin{}
	err = s.Models.Qs.Filter("account", AdminData.Account).One(&AdminDataChek)
	if err != orm.ErrNoRows {
		return 0, errors.New("账户已存在")
	}
	_ = AdminDataChek
	row, err := s.Models.Orm.Insert(&AdminData)
	if err != nil {
		logs.Error(err)
		return row, errors.New("新增用户失败，如果您是系统管理员您可以通过错误日志查看详细错误信息")
	}
	return row, err
}

// 编辑用户
func (s *AdminService) Edit(AdminData models.Admin) (int64, error) {
	err := AdminData.ValidEdit()
	if err != nil {
		return 0, err
	}
	if s.Models.Qs == nil {
		s.Models.NewAdminQs()
	}
	AdminDataChek := models.Admin{}
	err = s.Models.Qs.Filter("id", AdminData.Id).One(&AdminDataChek)
	if err == orm.ErrNoRows {
		return 0, errors.New("账户不存在")
	} else if err != nil {
		logs.Error(err)
		return 0, errors.New("查询用户是否存在时出现错误，如果您是系统管理员您可以通过错误日志查看详细错误信息")
	}
	AdminData.Id = AdminDataChek.Id
	if AdminData.Password == "" {
		AdminData.Password = AdminDataChek.Password
		AdminData.Interfere = AdminDataChek.Interfere
	}
	row, err := s.Models.Orm.Update(&AdminData)
	if err != nil {
		logs.Error(err)
		return row, errors.New("编辑用户失败，如果您是系统管理员您可以通过错误日志查看详细错误信息")
	}
	return row, err
}

// 删除用户
func (s *AdminService) Delete(id int) error {
	if s.Models.Qs == nil {
		s.Models.NewAdminQs()
	}
	_, err := s.Models.Qs.Filter("id", id).Delete()
	if err != nil {
		logs.Error(err)
		return errors.New("删除用户失败，如果您是系统管理员您可以通过错误日志查看详细错误信息")
	}
	return err
}

//登录
func (s *AdminService) Login(Account string, password string, ip string) (models.Admin, error) {
	if s.Models.Qs == nil {
		s.Models.NewAdminQs()
	}
	AdminInfo := models.Admin{}
	err := s.Models.Qs.Filter("account", Account).One(&AdminInfo)
	if err != nil {
		if err == orm.ErrNoRows {
			return AdminInfo, errors.New("账户不存在")
		} else {
			logs.Error("账户%s在登录时数据库出现错误，错误信息：%s", Account, err.Error())
			return AdminInfo, errors.New("数据库出现错误，如果您的管理员，请前往查看错误日志，如果您不是管理员，您可以尝试再次登录")
		}
	}
	if AdminInfo.Password != utils.Password(password, AdminInfo.Interfere) {
		return AdminInfo, errors.New("密码错误")
	}
	if AdminInfo.State != 1 {
		return models.Admin{}, errors.New("账户禁用！")
	}
	AdminRoleService := AdminRoleService{}
	AdminRoleService.Init()
	role, err := AdminRoleService.Details(AdminInfo.Role)
	if err != nil {
		return models.Admin{}, errors.New("权限组不存在！")
	}
	if role.State != 1 {
		return models.Admin{}, errors.New("权限组禁用！")
	}
	// 签发身份
	AdminJwtKey, err := beego.AppConfig.String("jwt::admin_key")
	if err != nil || AdminJwtKey == "" {
		logs.Error("没有配置admin_jwt_key，将使用默认www.bugquit.com")
		AdminJwtKey = "www.bugquit.com"
	}
	AdminJwt := &jwt_admin.JWT{
		[]byte(AdminJwtKey),
	}
	ExpirationTime := time.Now()
	TimeTmp, err := time.ParseDuration("86400s")
	if err != nil {
		logs.Error("定义的JWT过期时间转换失败，使用默认时间86400s")
		TimeTmp, _ = time.ParseDuration("86400s")
	}
	ExpirationTime = ExpirationTime.Add(TimeTmp)
	_ = TimeTmp
	Claims := jwt_admin.CustomClaims{
		AdminInfo.Id,
		AdminInfo,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix()), // 签名生效时间
			ExpiresAt: ExpirationTime.Unix(),    // 过期时间
			Issuer:    AdminJwtKey,              //签名的发行者
		},
	}
	token, err := AdminJwt.CreateToken(Claims)
	if err != nil {
		logs.Error("签名签发错误:%s", err.Error())
		return AdminInfo, errors.New("签名签发错误!")
	}
	AdminInfo.Token = token
	// 登录日志
	LogData := models.AdminLog{}
	LogData.Aid = AdminInfo.Id
	LogData.Content = "使用IP" + ip + "在" + time.Now().Format("2006-01-02 15:04:05") + "登录"
	LogData.Type = 1
	AdminLogService := AdminLogService{}
	AdminLogService.Init()
	_, err = AdminLogService.Add(LogData)
	if err != nil {
		logs.Error("写入登录日志时出错：%s", err.Error())
	}
	// 记录在线用户
	AdminOnLineData := models.AdminOnLine{}
	AdminOnLineService := AdminOnLineService{}
	AdminOnLineService.Init()
	AdminOnLineData.Aid = AdminInfo.Id
	AdminOnLineData.Account = AdminInfo.Account
	AdminOnLineData.Token = AdminInfo.Token
	LibTime := lib.Time{ExpirationTime}
	AdminOnLineData.ExpirationTime = LibTime
	_, err = AdminOnLineService.Add(AdminOnLineData)
	if err != nil {
		logs.Error("写入在线账户时出错：%s", err.Error())
	}
	return AdminInfo, nil
}

// 通过token获取账户信息
func (s *AdminService) CtxTokenGetAdminInfo(claims *jwt_admin.CustomClaims) (models.Admin, error) {
	AdminInfo := models.Admin{}
	c, err := json.Marshal(claims.AdminInfo)
	if err != nil {
		logs.Error(err)
		return AdminInfo, errors.New("从token中获取账户信息时失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	err = json.Unmarshal(c, &AdminInfo)
	if err != nil {
		logs.Error(err)
		return AdminInfo, errors.New("从token中获取账户信息时失败，如果您是系统管理员，您可以通过错误日志查看详细信息")
	}
	return AdminInfo, err
}
