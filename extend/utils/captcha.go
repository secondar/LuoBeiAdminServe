package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
	"time"
)

type Captcha struct {
	Store  base64Captcha.Store
	Driver base64Captcha.Driver
}
type Result struct {
	Id  string `json:"id"`
	Src string `json:"src"`
}

// 初始化  验证码存储个数   存储时间
func (s *Captcha) Init(collectNum int, expiration time.Duration) {
	s.Store = base64Captcha.NewMemoryStore(collectNum, expiration*time.Minute)
}

// 生成图形化算术验证码配置
func (s *Captcha) UseMathConfig(Height int, Width int) {
	var (
		NoiseCount      = 0
		ShowLineOptions = base64Captcha.OptionShowHollowLine
		BgColor         = &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}
		Fonts []string
	)
	s.Driver = base64Captcha.NewDriverMath(Height, Width, NoiseCount, ShowLineOptions, BgColor, Fonts)
}

// 生成图形化字符串验证码配置
func (s *Captcha) UseStringConfig(Height int, Width int, Length int) {
	var (
		NoiseCount      = 0
		ShowLineOptions = base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine
		Source          = "123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
		BgColor         = &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}
	)
	s.Driver = base64Captcha.NewDriverString(Height, Width, NoiseCount, ShowLineOptions, Length, Source, BgColor, nil)
}

// 生成图形化汉字验证码配置
func (s *Captcha) UseChineseConfig(Height int, Width int, Length int) {
	var (
		Source  = "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,不想要,的值"
		BgColor = &color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}
	)
	s.Driver = base64Captcha.NewDriverChinese(Height, Width, 0, base64Captcha.OptionShowSlimeLine, Length, Source, BgColor, nil)
}

// 生成图形化数字音频验证码配置Language "zh"
func (s *Captcha) UseAutoConfig(Length int, Language string) {
	s.Driver = base64Captcha.NewDriverAudio(Length, Language)
}

// 生成
func (s *Captcha) Generate() (Result, error) {
	c := base64Captcha.NewCaptcha(s.Driver, s.Store)
	id, b64s, err := c.Generate()
	return Result{id, b64s}, err
}

// 验证 验证码ID 答案 是否清除
func (s *Captcha) VerifyCaptcha(id, VerifyValue string, clear bool) bool {
	return s.Store.Verify(id, VerifyValue, clear)
}

// 获取验证码答案 验证码ID 是否清除
func (s *Captcha) GetCodeAnswer(codeId string, clear bool) string {
	return s.Store.Get(codeId, clear)
}
