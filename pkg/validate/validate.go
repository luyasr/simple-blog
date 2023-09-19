package validate

import (
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Struct(obj any) error {
	validate := validator.New()
	// 英文翻译器
	en := en.New()
	// 中文翻译器
	zh := zh.New()
	uni := ut.New(en, zh)
	// 获取需要的语言
	trans, _ := uni.GetTranslator("zh")
	// 注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// 验证器注册翻译器
	_ = zhTrans.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
