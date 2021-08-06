package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/go-playground/locales/zh"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    zhs "github.com/go-playground/validator/v10/translations/zh"
)

type TestController struct {
}

var (
    validate = validator.New()          // 實例化驗證器
    chinese  = zh.New()                 // 獲取中文翻譯器
    uni      = ut.New(chinese, chinese) // 設置成中文翻譯器
    trans, _ = uni.GetTranslator("zh")  // 獲取翻譯字典
)
type User struct {
    Name  string `form:"name" validate:"required,min=3,max=5"`
    Email string `form:"email" validate:"email"`
    Age   int8   `form:"age" validate:"gte=18,lte=20"`
}


func (t *TestController) Index(context *gin.Context) {
        var user User
        err := context.ShouldBindQuery(&user)
        if err != nil {
            context.JSON(500, gin.H{"msg": err})
            return
        }
        // 註冊翻譯器
        _ = zhs.RegisterDefaultTranslations(validate, trans)
        // 使用驗證器驗證
        err = validate.Struct(user)
        if err != nil {
            if errors, ok := err.(validator.ValidationErrors); ok {
                // 翻譯，並返回
                context.JSON(500, gin.H{
                    "翻譯前": errors.Error(),
                    "翻譯後": errors.Translate(trans),
                })
                return
            }
        }
        context.JSON(200,gin.H{"msg":"success"})
}
