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

type ReqRegisterUser struct {
	Id         int
	Phone      string
	WechatId   string
	NickName   string
	Password   string
	UserType   int //0代表司机  1代表老板
	UserStatus int //0代表空闲  1代表忙碌
	Address    string
	//经度
	Latitude float64
	//纬度
	Longitude float64
	Age       int
}

type ReqLoginUser struct {
	Phone    string
	Password string
	UserType int

	Address string
	//经度
	Latitude float64
	//纬度
	Longitude float64
}

type ReqDriverStatus struct {
	DriverId     int
	DriverStatus int

	Address string
	//经度
	Latitude float64
	//纬度
	Longitude float64
}

type ReqAddFriends struct {
	Phone  string
	BossId int
}
