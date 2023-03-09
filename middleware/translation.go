package middleware

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	"github.com/go-playground/locales/fr"
	"github.com/go-playground/locales/id"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	esTranslations "github.com/go-playground/validator/v10/translations/es"
	frTranslations "github.com/go-playground/validator/v10/translations/fr"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/raven0520/wallet/params"
)

func GetCurrentLang(lang string) string {
	var curr string
	switch lang {
	case "en_US":
		curr = "en"
	case "zh_CN", "zh-CN", "zh_Hans":
		curr = "zh"
	default:
		curr = "en"
	}
	return curr
}

// TranslationMiddleware 设置 Translation
func TranslationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 参照：https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go

		// 设置支持语言
		_en := en.New()
		_zh := zh.New()
		_id := id.New()
		_fr := fr.New()
		_es := es.New()
		// 设置国际化翻译器
		uni := ut.New(_zh, _zh, _en, _es, _id, _fr)
		val := validator.New()

		// 根据参数取翻译器实例
		locale := context.GetHeader("Accept-Language")
		curr := GetCurrentLang(locale)
		trans, _ := uni.GetTranslator(curr)

		//翻译器注册到validator
		switch curr {
		case "zh": // 中文
			err := zhTranslations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})
		case "es": // 西班牙
			err := esTranslations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("es_comment")
			})
		case "id": // 印度尼西亚
			err := idTranslations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("it_comment")
			})
		case "fr": // 法语
			err := frTranslations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("fr_comment")
			})
		case "en": // 英语
		case "en_US": // 美式英语
			err := enTranslations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("en_comment")
			})
		default: // 默认英语
			err := enTranslations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				return
			}
			val.RegisterTagNameFunc(func(fld reflect.StructField) string {
				return fld.Tag.Get("comment")
			})

			//自定义验证方法
			//https://github.com/go-playground/validator/blob/v9/_examples/custom-validation/main.go
			//val.RegisterValidation("is-validuser", func(fl validator.FieldLevel) bool {
			//	return fl.Field().String() == "admin"
			//})

			//自定义验证器
			//https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
			//val.RegisterTranslation("is-validuser", trans, func(ut ut.Translator) error {
			//	return ut.Add("is-validuser", "{0} 填写不正确哦", true)
			//}, func(ut ut.Translator, fe validator.FieldError) string {
			//	t, _ := ut.T("is-validuser", fe.Field())
			//	return t
			//})
		}
		context.Set(params.ValidatorKey, val)
		context.Next()
	}
}
