package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

type User struct {
	Id         int     `orm:"auto"`
	Phone      string
	WechatId   string
	NickName   string
	Password   string
	UserType   int //0代表司机  1代表老板
	UserStatus int //0代表空闲  1代表忙碌
	Email      string
	Age        int
	Created    time.Time `orm:"index"`
}

func init() {

}

func AddUser(u User) (bool, int64, error) {

	u.Created = time.Now()

	isExist, id, err := orm.NewOrm().ReadOrCreate(&u, "Phone")

	return isExist, id, err
}

func Login(username string) User {

	user := User{Phone: username}
	err := orm.NewOrm().Read(&user, "Phone")
	if err != nil {
		beego.Error(err)
	}
	return user
}

func GetAllDriver(pageNo, pageSize int) (Page, error) {

	qr := orm.NewOrm().QueryTable("user").Filter("UserType", 0)

	page, err := PageAdd(qr, pageSize, pageNo, func() (interface{}, error) {
		var users []User
		_, err := qr.Limit(pageSize, (pageNo-1)*pageSize).All(&users)
		return users, err
	})

	if err != nil {
		beego.Error(err)
	}

	return page, err
}



func GetIdleDriver(pageNo, pageSize int) (Page, error) {

	var err error

	qr := orm.NewOrm().QueryTable("user").Filter("UserType", 0).Filter("UserStatus", 0)

	page, err := PageAdd(qr, pageSize, pageNo, func() (interface{}, error) {
		var users []User
		_, err := qr.Limit(pageSize, (pageNo-1)*pageSize).All(&users)
		return users, err
	})

	if err != nil {
		beego.Error(err)
	}
	return page, err
}

func SetDriverStatus(id, status int) (bool, error) {

	var err error
	user := User{Id: id}
	o := orm.NewOrm()
	err = o.Read(&user)
	if err == nil {
		user.UserStatus = status
		_, err = o.Update(&user)
		if err == nil {
			return true, nil
		}
	}
	return false, err
}

func DeleteUser(uid string) bool {
	user := orm.NewOrm().QueryTable("user").Filter("Uid", uid)
	_, err := user.Delete()
	if err != nil {
		beego.Error(err)
		return false
	}
	return true
}
