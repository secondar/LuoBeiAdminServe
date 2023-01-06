package jwt_admin

import (
	"LuoBeiAdminServeForGolang/extend/lib"
	"LuoBeiAdminServeForGolang/extend/utils"
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	context "github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
)

// JwtAuth 中间件，检查token
func JwtAuth(next beego.FilterFunc) beego.FilterFunc {
	return func(ctx *context.Context) {
		token := ctx.Request.Header.Get("X-Token")
		if token == "" {
			OutJson(ctx, utils.ResultJson{Code: 401, Msg: "无权访问"})
			return
		}
		Jwt := NewJWT()
		claims, err := Jwt.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				OutJson(ctx, utils.ResultJson{Code: -1, Msg: "授权已过期"})
			} else {
				OutJson(ctx, utils.ResultJson{Code: -1, Msg: "非法的授权"})
			}
			return
		}
		ctx.Input.SetData("admin_token_claims", claims)
		// 基础权限，所有账户都得给通行
		if "/admin/api/v1/menu/getadminmenurouter" == ctx.Input.URL() {
			next(ctx)
			return
		}
		//
		// 获取账户信息
		AdminInfo := claims.AdminInfo.(map[string]interface{})
		// 如果权限组不是超级管理员那就校验权限
		if fmt.Sprintf("%v", AdminInfo["role"]) != "1" {
			tablePrefix, _ := beego.AppConfig.String("mysql::tableprefix")
			type AdminMenu struct {
				Id             int         `orm:"pk;auto;size(11)" json:"id"`
				Pid            int         `orm:"size(11)" json:"pid"`
				Title          string      `orm:"size(255)" json:"title"`
				Type           int8        `orm:"size(1)" json:"type"`
				Icon           *string     `orm:"size(255)" json:"icon"`
				Show           int8        `orm:"size(1)" json:"show"`
				Link           int8        `orm:"size(1)" json:"link"`
				ApiPath        string      `orm:"size(255)" json:"api_path"`
				Characteristic string      `orm:"size(255)" json:"characteristic"`
				Router         *string     `orm:"size(255)" json:"router"`
				Sort           *int        `orm:"size(11)" json:"sort"`
				Component      *string     `orm:"size(255)" json:"component"`
				Path           *string     `orm:"size(255)" json:"path"`
				Addtime        lib.Time    `orm:"auto_now_add" json:"addtime"`
				Children       []AdminMenu `orm:"-" json:"children"`
			}
			AdminMenuInfo := []AdminMenu{}
			Tsql := fmt.Sprintf("SELECT admin_menu.* FROM %sadmin_menu admin_menu RIGHT JOIN %sadmin_router admin_router ON admin_router.menu=admin_menu.id WHERE admin_router.role = ?", tablePrefix, tablePrefix)
			_, err = orm.NewOrm().Raw(Tsql, AdminInfo["role"]).QueryRows(&AdminMenuInfo)
			if err != nil {
				logs.Error(err)
				OutJson(ctx, utils.ResultJson{Code: 0, Msg: "权限验证失败，并非是验证不通过，而是代码出现错误，如果您是系统管理员，您可以通过错误日志查看详细信息"})
				return
			}
			for _, item := range AdminMenuInfo {
				if utils.TrimSpace(item.ApiPath) == ctx.Input.URL() {
					next(ctx)
					return
				}
			}
			// 到这自然就无权了
			OutJson(ctx, utils.ResultJson{Code: 401, Msg: "无权访问"})
			return
		} else {
			next(ctx)
		}
	}
}

// 这是权限输出给权限的
func OutJson(ctx *context.Context, OutData utils.ResultJson) {
	ctx.Output.Status = 200
	ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	ctx.Output.Header("Access-Control-Allow-Headers", "x-token,X-Token")
	ctx.Output.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
	ctx.Output.JSON(OutData, true, true)

}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "www.bugquit.com"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID        int         `json:"id"`
	AdminInfo interface{} `json:"admin_info"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
