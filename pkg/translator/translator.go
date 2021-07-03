package translator

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func NewTrans(locale string) (ut.Translator, error) {
	// 修改gin框架的validator引擎属性，实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() // 中文翻译器
		enT := en.New() // 英文翻译器
		// 第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		trans, ok := uni.GetTranslator(locale)
		if !ok {
			return nil, fmt.Errorf("uni.GetTranslator(%s) error", locale)
		}

		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			_ = en_translations.RegisterDefaultTranslations(v, trans)
		}

		//// 注册一个获取json的tag的自定义方法
		//v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		//	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		//	if name == "-" {
		//		return ""
		//	}
		//	return name
		//})

		//// 注册自定义验证器
		//_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		//_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
		//	return ut.Add("mobile", "{0}非法的手机号码", true) // see universal-translator for details
		//}, func(ut ut.Translator, fe translator.FieldError) string {
		//	t, _ := ut.T("mobile", fe.Field())
		//	return t
		//})

		return trans, nil
	}

	return nil, nil
}
