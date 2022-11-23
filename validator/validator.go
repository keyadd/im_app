package validator

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// 定义一个全局翻译器T
var trans ut.Translator

// InitTrans 初始化翻译器
func InitTrans(locale string) (trans ut.Translator) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	//var ok bool
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//自定义验证方法
		//LoginV(v)

		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 为SignUpParam注册自定义校验方法
		//v.RegisterStructValidation(SignUpParamStructLevelValidation, model.ParamSignUp{})

		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			fmt.Printf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err := enTranslations.RegisterDefaultTranslations(v, trans)
			fmt.Printf("init validator failed,err:%v\n", err)
		}

	}
	return trans
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
//func SignUpParamStructLevelValidation(sl validator.StructLevel) {
//	su := sl.Current().Interface().(model.ParamSignUp)
//
//	if su.Password != su.RePassword {
//		// 输出错误提示信息，最后一个参数就是传递的param
//		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
//	}
//}
