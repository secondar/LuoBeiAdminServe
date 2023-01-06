package admin

import (
	"LuoBeiAdminServeForGolang/extend/utils"
	"crypto/md5"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"math/rand"
	"os"
	"path"
	"time"
)

type UpFileController struct {
	beego.Controller
}

func (_this *UpFileController) UploadImg() {
	ResultJson := utils.ResultJson{}
	f, h, _ := _this.GetFile("file") //获取上传的文件
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".ico":  true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		ResultJson.Code = 401
		ResultJson.Msg = "上传文件不符合标准"
		_this.Data["json"] = ResultJson
		return
	}
	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll(uploadDir, 777)
	if err != nil {
		logs.Error(err)
		ResultJson.Code = 503
		ResultJson.Msg = "文件夹创建失败"
		_this.Data["json"] = ResultJson
		_ = _this.ServeJSON()
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ))
	fileName := fmt.Sprintf("%x", hashName) + ext
	fpath := uploadDir + fileName
	defer f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = _this.SaveToFile("file", fpath)
	if err != nil {
		logs.Error(err)
		ResultJson.Code = 503
		ResultJson.Msg = "文件上传失败"
	} else {
		web_domain, _ := beego.AppConfig.String("default::web_domain")
		ResultJson.Code = 200
		ResultJson.Msg = web_domain + "/" + fpath
	}
	_this.Data["json"] = ResultJson
	_ = _this.ServeJSON()
}
