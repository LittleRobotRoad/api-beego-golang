package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
	"errors"
)

type DriverUser struct {
	Id       int                `orm:"auto"`
	Phone    string
	WechatId string
	Password string
	Address  string
	//经度
	Latitude float64  `orm:"digits(9);decimals(6)"`
	//纬度
	Longitude  float64  `orm:"digits(9);decimals(6)"`
	NickName   string
	UserStatus int //0代表空闲  1代表忙碌
	BossUser   []*BossUser  `orm:"reverse(many)"`
	Created    time.Time        `orm:"index"`
}

type BossUser struct {
	Id       int            `orm:"auto"`
	Phone    string
	WechatId string
	NickName string
	Password string
	Address  string
	//经度
	Latitude float64  `orm:"digits(9);decimals(6)"`
	//纬度
	Longitude  float64  `orm:"digits(9);decimals(6)"`
	DriverUser []*DriverUser  `orm:"rel(m2m);rel_through(github.com/LittleRobotRoad/api-beego-golang/models.BossToDriver)"`
	Created    time.Time      `orm:"index"`
}

type BossToDriver struct {
	Id         int                                `orm:"auto"`
	BossUser   *BossUser                        `orm:"rel(fk)"`
	DriverUser *DriverUser                `orm:"rel(fk)"`
}

func init() {

}

// bool 是否成功    int64 插入user表的行数   error错误
func AddUser(u ReqRegisterUser) (bool, int64, error) {

	o := orm.NewOrm()
	var isSuccess bool
	var id int64
	var err error

	if u.UserType == 0 {
		driverUser := DriverUser{
			WechatId:   u.WechatId,
			Password:   u.Password,
			Address:    u.Address,
			Latitude:   u.Latitude,
			Longitude:  u.Longitude,
			Phone:      u.Phone,
			NickName:   u.NickName,
			UserStatus: u.UserStatus,
			Created:    time.Now(),
		}
		isSuccess, id, err = o.ReadOrCreate(&driverUser, "Phone")
	} else if u.UserType == 1 {
		bossUser := BossUser{
			Phone:     u.Phone,
			NickName:  u.NickName,
			Created:   time.Now(),
			WechatId:  u.WechatId,
			Password:  u.Password,
			Address:   u.Address,
			Latitude:  u.Latitude,
			Longitude: u.Longitude,
		}
		isSuccess, id, err = o.ReadOrCreate(&bossUser, "Phone")
	}

	//用户存在直接返回
	if !isSuccess {
		beego.Error(err)
		beego.Error("注册用户已经存在")
		return isSuccess, id, err
	}

	return isSuccess, id, err
}

func Login(reqBody ReqLoginUser) (int, interface{}, error) {

	var err error
	var code int

	o := orm.NewOrm()
	if reqBody.UserType == 0 {
		driver := DriverUser{Phone: reqBody.Phone}
		err = o.Read(&driver, "Phone")

		//判断为空不存入数据库
		if reqBody.Address != "" {
			driver.Address = reqBody.Address
		}
		if reqBody.Latitude != 0 {
			driver.Latitude = reqBody.Latitude
		}
		if reqBody.Longitude != 0 {
			driver.Longitude = reqBody.Longitude
		}

		if driver.Id != 0 {
			if reqBody.Password == driver.Password {
				code = 200
				_, err = o.Update(&driver)
				return code, driver, err
			} else {
				code = 302
				err = errors.New("password error")
			}
		} else {
			code = 303
			err = errors.New("user not exist")
		}

	} else if reqBody.UserType == 1 {
		boss := BossUser{Phone: reqBody.Phone}
		err = o.Read(&boss, "Phone")

		//判断为空不存入数据库
		if reqBody.Address != "" {
			boss.Address = reqBody.Address
		}
		if reqBody.Latitude != 0 {
			boss.Latitude = reqBody.Latitude
		}
		if reqBody.Longitude != 0 {
			boss.Longitude = reqBody.Longitude
		}

		if boss.Id != 0 {
			if reqBody.Password == boss.Password {
				code = 200
				_, err = o.Update(&boss)
				return code, boss, err
			} else {
				code = 302
				err = errors.New("password error")
			}
		} else {
			code = 303
			err = errors.New("user not exist")
		}
	}

	return code, nil, err
}

func GetAllDriver(pageNo, pageSize int) (Page, error) {

	qr := orm.NewOrm().QueryTable("driver_user")

	page, err := PageAdd(qr, pageSize, pageNo, func() (interface{}, error) {
		var users []DriverUser
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

	qr := orm.NewOrm().QueryTable("driver_user").Filter("UserStatus", 0)

	page, err := PageAdd(qr, pageSize, pageNo, func() (interface{}, error) {
		var users []DriverUser
		_, err := qr.Limit(pageSize, (pageNo-1)*pageSize).All(&users)
		return users, err
	})

	if err != nil {
		beego.Error(err)
	}
	return page, err
}

func GetBossAllDriver(bossId, pageNo, pageSize int) (Page, error) {

	var err error

	o := orm.NewOrm()

	var userss []*DriverUser
	//BossUser__BossUser__Id  第一个是参数名 第二个是model名 第三个是属性名
	qr := o.QueryTable("driver_user").Filter("BossUser__BossUser__Id", bossId)

	beego.Debug(qr.All(&userss))

	page, err := PageAdd(qr, pageSize, pageNo, func() (interface{}, error) {
		var users []DriverUser
		_, err := qr.Limit(pageSize, (pageNo-1)*pageSize).All(&users)
		return users, err
	})

	if err != nil {
		beego.Error(err)
	}
	return page, err
}

func GetBossIdleDriver(bossId, pageNo, pageSize int) (Page, error) {

	var err error

	o := orm.NewOrm()

	var userss []*DriverUser
	//BossUser__BossUser__Id  第一个是参数名 第二个是model名 第三个是属性名
	qr := o.QueryTable("driver_user").Filter("BossUser__BossUser__Id", bossId).Filter("UserStatus", 0)

	beego.Debug(qr.All(&userss))

	page, err := PageAdd(qr, pageSize, pageNo, func() (interface{}, error) {
		var users []DriverUser
		_, err := qr.Limit(pageSize, (pageNo-1)*pageSize).All(&users)
		return users, err
	})

	if err != nil {
		beego.Error(err)
	}
	return page, err
}

func SetDriverStatus(reqBody ReqDriverStatus) (bool, error) {

	var err error
	o := orm.NewOrm()
	driver := DriverUser{Id: reqBody.DriverId}
	err = o.Read(&driver)
	if err == nil {
		driver.UserStatus = reqBody.DriverStatus
		//判断为空不存入数据库
		if reqBody.Address != "" {
			driver.Address = reqBody.Address
		}
		if reqBody.Latitude != 0 {
			driver.Latitude = reqBody.Latitude
		}
		if reqBody.Longitude != 0 {
			driver.Longitude = reqBody.Longitude
		}

		_, err = o.Update(&driver)
		if err == nil {
			return true, nil
		}
	}
	return false, err
}

func DeleteUser(uid string) bool {
	user := orm.NewOrm().QueryTable("user").Filter("Id", uid)
	_, err := user.Delete()
	if err != nil {
		beego.Error(err)
		return false
	}
	return true
}

func AddFriend(bossuid int, driverPhone string) (bool, error) {
	o := orm.NewOrm()

	driver := DriverUser{Phone: driverPhone}
	err := o.Read(&driver,"Phone")

	if err != nil {
		beego.Error(err)
		return false, err
	}

	boss := BossUser{Id: bossuid}
	err = o.Read(&boss)
	if err != nil {
		beego.Error(err)
		return false, err
	}

	//insert into post_tags
	m2m := o.QueryM2M(&boss, "DriverUser")

	exist := m2m.Exist(driver)
	if !exist {
		_, err = m2m.Add(driver)
	} else {
		return true, errors.New("已经是好友")
	}
	return true, err
}

func DeleteFriend(bossuid, driveruid int) (bool, error) {
	o := orm.NewOrm()

	driver := DriverUser{Id: driveruid}
	err := o.Read(&driver)

	if err != nil {
		beego.Error(err)
		return false, err
	}

	boss := BossUser{Id: bossuid}
	err = o.Read(&boss)
	if err != nil {
		beego.Error(err)
		return false, err
	}

	//insert into post_tags
	m2m := o.QueryM2M(&boss, "DriverUser")

	exist := m2m.Exist(driver)
	if !exist {
		return true, errors.New("不是好友关系")
	} else {
		_, err = m2m.Remove(driver)
	}
	return true, err
}
