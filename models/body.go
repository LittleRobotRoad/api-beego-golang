package models


type Result struct {
	Code   int
	Object interface{}
}

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	LastPage   bool
	List       interface{}
}

type SwaggerRegisterUser struct {
	Id         int
	Phone   string
	WechatId   string
	NickName   string
	Password   string
	UserType   int //0代表司机  1代表老板
	UserStatus int //0代表空闲  1代表忙碌
	Email      string
	Age        int
}

type SwaggerLoginUser struct {
	Phone string
	Password string
}

type SwaggerDriverStatus struct {
	DriverId int
	DriverStatus int
}
