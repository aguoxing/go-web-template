package captcha

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type Captcha struct{}

// 设置自带的store
// var store = base64Captcha.DefaultMemStore
var store base64Captcha.Store = RedisStore{}

// CaptMake 生成验证码
func CaptMake() (id, b64s string, err error) {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	// 配置验证码信息
	//dight 数字验证码
	//audio 语音验证码
	//string 字符验证码
	//math 数学验证码(加减乘除)
	//chinese中文验证码
	captchaConfig := base64Captcha.DriverString{
		Length:     4,                                      // 验证码长度
		Height:     38,                                     // 验证码图片高度
		Width:      105,                                    // 验证码图片宽度
		Source:     "abcdefghijklmnopqrstuvwxyz1234567890", // 会出现字符
		NoiseCount: 0,                                      // 干扰词数量
		//DotCount:        0,      // 干扰点数量
		//MaxSkew:         1,      // 倾斜角度 0-1
		ShowLineOptions: 2 | 4, // 线条数量
		BgColor: &color.RGBA{ // 背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	lid, lb64s, lerr := captcha.Generate()
	return lid, lb64s, lerr
}

// CaptVerify 验证captcha是否正确
func CaptVerify(id string, capt string) bool {
	return store.Verify(id, capt, false)
}
