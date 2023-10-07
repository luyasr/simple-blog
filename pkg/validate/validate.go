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
	enTranslator := en.New()
	// 中文翻译器
	zhTranslator := zh.New()
	uni := ut.New(enTranslator, zhTranslator)
	// 获取需要的语言
	trans, _ := uni.GetTranslator("zhTranslator")
	// 注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})

	// 验证器注册翻译器
	_ = zhTrans.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(obj)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return errors.New(err.Error())
		}
		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(e.Translate(trans))
		}
	}
	return nil
}
