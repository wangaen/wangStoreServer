package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	translator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
)

var vp translator.Translator

func initTrans(locale string) (err error) {
	// 修改gin框架中的validator属性，实现自定义定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		enLang := en.New() // 英文翻译器
		zhLang := zh.New() // 中文翻译器

		universalTranslator := translator.New(enLang, enLang, zhLang)

		// locale 通常取决于http请求头的 'Accept-Language'
		var ok bool
		vp, ok = universalTranslator.GetTranslator(locale)
		if !ok {
			return errors.New("无法获取参数校验器中文转换器")
		}

		// 注册一个获取参数的json标签的自定义方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			return field.Tag.Get("json")
		})

		// 注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, vp)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, vp)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, vp)
		}
		return
	}
	return
}

func ValidParams(c *gin.Context, form interface{}) bool {
	if err := initTrans("zh"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
		return false
	}
	if err := c.ShouldBindBodyWith(form, binding.JSON); err != nil {
		transErr, ok := err.(validator.ValidationErrors)
		if !ok {
			return jsonUnmarshalTypeError(c, err)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "fail",
				"msg":    removeStructName(transErr.Translate(vp)),
			})
		}
		return false
	}
	return true
}

func jsonUnmarshalTypeError(c *gin.Context, err error) bool {
	if jsonTypeRrr, is := err.(*json.UnmarshalTypeError); is {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": gin.H{
				jsonTypeRrr.Field: jsonTypeRrr.Field + " 必须是 " + jsonTypeRrr.Type.String() + "类型",
			},
			"status": "fail",
		})
	} else if err.Error() == "EOF" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"msg":    "参数不能为空",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"msg":    err.Error(),
		})
	}
	return false
}

// LoginForm.username -> username
func removeStructName(fields map[string]string) map[string]string {
	var errMap = make(map[string]string)
	for field, err := range fields {
		errMap[field[strings.Index(field, ".")+1:]] = err
	}
	return errMap
}
