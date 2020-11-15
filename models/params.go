package models

//定义请求参数的结构体

type ParamSignUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}