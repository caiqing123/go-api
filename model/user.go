package model

type User struct {
	Username string `json:"username" form:"username" validate:"required,min=3,max=10" label:"用户名"`
	Password string  `json:"password" form:"password" validate:"required,gte=1,lte=10" label:"密码"`
}
