package model

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhs "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

var (
	validate = validator.New()          // 實例化驗證器
	chinese  = zh.New()                 // 獲取中文翻譯器
	uni      = ut.New(chinese, chinese) // 設置成中文翻譯器
	trans, _ = uni.GetTranslator("zh")  // 獲取翻譯字典
)

// InitTrans 初始化翻译器
func InitTrans(model interface{}) validator.ValidationErrorsTranslations {
	// 注册翻译器
	_ = zhs.RegisterDefaultTranslations(validate, trans)
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("label")
	})
	// 使用验证器验证
	err := validate.Struct(model)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			// 翻译，並返回
			var errList = make(map[string]string, len(errors))
			//修改返回格式
			for _, err := range errors {
				// 获取原来的标签 - form
				fieldName := err.StructField()
				t := reflect.TypeOf(model)
				field, _ := t.FieldByName(fieldName)
				j := field.Tag.Get("form")
				// 修改
				errList[j] = err.Translate(trans)
			}
			return errList
		}
	}
	return nil
}
