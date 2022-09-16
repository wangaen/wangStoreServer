package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"net/http"
	"wangStoreServer/common/controller"
	"wangStoreServer/common/utils"
)

type captchaBody struct {
	Id           string `json:"id"`
	verifyCode   string `json:"verifyCode"`
	DriverString *base64Captcha.DriverString
	DriverMath   *base64Captcha.DriverMath
}

var store = base64Captcha.DefaultMemStore

func (c *captchaBody) initDriverString() *base64Captcha.DriverString {
	c.DriverString = &base64Captcha.DriverString{
		Height:          40,
		Width:           80,
		Length:          4,
		NoiseCount:      64, // 数字干扰
		ShowLineOptions: 2,  // 波浪线干扰
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 255, G: 255, B: 255, A: 30},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	return c.DriverString.ConvertFonts()
}

func (c *captchaBody) initDriverMath() *base64Captcha.DriverMath {
	c.DriverMath = &base64Captcha.DriverMath{
		Height:          40,
		Width:           80,
		NoiseCount:      32, // 数字干扰
		ShowLineOptions: 2,  // 波浪线干扰
		BgColor:         &color.RGBA{R: 255, G: 255, B: 255, A: 30},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	return c.DriverMath.ConvertFonts()
}

// 获取验证码
func generateCaptchaHandler(c *gin.Context) {
	var param = &captchaBody{}
	var v *base64Captcha.Captcha
	if utils.GetRandomNum(2) == 0 {
		v = base64Captcha.NewCaptcha(param.initDriverString(), store)
	} else {
		v = base64Captcha.NewCaptcha(param.initDriverMath(), store)
	}
	id, b64s, err := v.Generate()
	if err != nil {
		controller.ErrResponseBody(c, http.StatusInternalServerError, "error", err.Error())
		return
	}
	controller.OkResponseBody(c, "ok", "获取验证码成功", gin.H{
		"id":         id,
		"verifyCode": b64s,
	})
}

// 验证验证码
func captchaVerifyHandle(c *gin.Context) {
	var param captchaBody
	param.Id = c.Query("id")
	param.verifyCode = c.Query("verifyCode")
	if param.Id == "" || param.verifyCode == "" {
		controller.OkResponseBody(c, "fail", "参数 id 或 verifyCode 不能为空", "")
		return
	}
	if !store.Verify(param.Id, param.verifyCode, true) {
		controller.OkResponseBody(c, "fail", "验证码错误", "")
		return
	}
	controller.OkResponseBody(c, "ok", "验证码正确", "")

}

func InitCaptcha(r *gin.Engine) {
	v := r.Group("/captcha")
	{
		v.GET("/get", generateCaptchaHandler)
		v.GET("/verify", captchaVerifyHandle)
	}
}
