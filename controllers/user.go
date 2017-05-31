package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/LittleRobotRoad/api-beego-golang/models"
	"errors"
	"strconv"
	"github.com/LittleRobotRoad/api-beego-golang/utils"
)

var (
	err  error
	code int
)

// 关于司机用户和承运商用户的注册以及状态查询
type UserController struct {
	beego.Controller
}

func (c *UserController) URLMapping() {
	c.Mapping("logout", c.Logout)
	c.Mapping("login", c.Login)
	c.Mapping("add", c.Add)
	c.Mapping("driverstatus", c.DriverStatus)
	c.Mapping("getalldriver", c.GetAllDriver)
	c.Mapping("getidledriver", c.GetIdleDriver)
	c.Mapping("addfriends", c.AddFriends)
	c.Mapping("deletefriend", c.DeleteFriend)
	c.Mapping("getbossidledriver", c.GetBossIdleDriver)
	c.Mapping("getbossalldriver", c.GetBossAllDriver)
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {

	//初始化参数
	err = nil
	code = 200

	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = jsonPack(nil, err, code)
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	body		body 	models.ReqLoginUser	true		"body for user content"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {

	var reqBody models.ReqLoginUser
	//初始化参数
	err = nil
	code = 200

	json.Unmarshal(u.Ctx.Input.RequestBody, &reqBody)
	code, user, err := models.Login(reqBody)

	u.Data["json"] = jsonPack(user, err, code)
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {

	//初始化参数
	err = nil
	code = 200

	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// @Title GetAllDriver
// @Description 获取所有司机的信息
// @Param	pageSize		path 	string	true		"pageSize default 10"
// @Param	pageNo			path 	string	true		"pageNo default 1"
// @Success 200 {string} logout success
// @router /getalldriver/:pageNo/:pageSize [get]
func (u *UserController) GetAllDriver() {

	//初始化参数
	err = nil
	code = 200

	var users models.Page
	var pageNo, pageSize int

	pageNo = getParamInt(u, "pageNo", 1)
	pageSize = getParamInt(u, "pageSize", 10)

	users, err = models.GetAllDriver(pageNo, pageSize)

	u.Data["json"] = jsonPack(users, err, code)

	u.ServeJSON()
}

// @Title GetIdleDriver
// @Description 获取所有空闲司机的信息
// @Param	pageSize		path 	string	true		"pageSize default 10"
// @Param	pageNo			path 	string	true		"pageNo default 1"
// @Success 200 {string} logout success
// @router /getidledriver/:pageNo/:pageSize [get]
func (u *UserController) GetIdleDriver() {

	//初始化参数
	err = nil
	code = 200

	var users models.Page
	var pageNo, pageSize int

	pageNo = getParamInt(u, "pageNo", 1)
	pageSize = getParamInt(u, "pageSize", 10)

	users, err = models.GetIdleDriver(pageNo, pageSize)

	u.Data["json"] = jsonPack(users, err, code)

	u.ServeJSON()
}

// @Title GetIdleDriver
// @Description 获取所有空闲司机的信息
// @Param	pageSize		path 	string	true		"pageSize default 10"
// @Param	pageNo			path 	string	true		"pageNo default 1"
// @Param	bossId			path 	string	true		"pageNo default 1"
// @Success 200 {string} logout success
// @router /getbossidledriver/:bossId/:pageNo/:pageSize [get]
func (u *UserController) GetBossIdleDriver() {

	//初始化参数
	err = nil
	code = 200

	var users models.Page
	var pageNo, pageSize, bossId int

	pageNo = getParamInt(u, "pageNo", 1)
	pageSize = getParamInt(u, "pageSize", 10)
	bossId = getParamInt(u, "bossId", 10)

	users, err = models.GetBossIdleDriver(bossId, pageNo, pageSize)

	u.Data["json"] = jsonPack(users, err, code)

	u.ServeJSON()
}

// @Title GetIdleDriver
// @Description 获取所有空闲司机的信息
// @Param	pageSize		path 	string	true		"pageSize default 10"
// @Param	pageNo			path 	string	true		"pageNo default 1"
// @Param	bossId			path 	string	true		"pageNo default 1"
// @Success 200 {string} logout success
// @router /getbossalldriver/:bossId/:pageNo/:pageSize [get]
func (u *UserController) GetBossAllDriver() {

	//初始化参数
	err = nil
	code = 200

	var users models.Page
	var pageNo, pageSize, bossId int

	pageNo = getParamInt(u, "pageNo", 1)
	pageSize = getParamInt(u, "pageSize", 10)
	bossId = getParamInt(u, "bossId", 10)

	users, err = models.GetBossAllDriver(bossId, pageNo, pageSize)

	u.Data["json"] = jsonPack(users, err, code)

	u.ServeJSON()
}

// @Title setDriverStatus
// @Description 设置司机状态
// @Param	body		body 	models.ReqDriverStatus	true		"body for user content"
// @Success 200 {string} logout success
// @router /driverstatus [post]
func (u *UserController) DriverStatus() {

	//初始化参数
	err = nil
	code = 200
	var reqBody models.ReqDriverStatus
	var isSuccess bool

	err =json.Unmarshal(u.Ctx.Input.RequestBody,&reqBody)

	isSuccess, err = models.SetDriverStatus(reqBody)

	if !isSuccess {
		code = 300
	}
	u.Data["json"] = jsonPack(nil, err, code)
	u.ServeJSON()
}

// @Title AddFriends
// @Description 承运商添加司机好友
// @Param	body		body 	models.ReqAddFriends	true		"body for user content"
// @Success 200 {string} logout success
// @router /addfriends [post]
func (u *UserController) AddFriends() {

	//初始化参数
	err = nil
	code = 200
	var reqBody models.ReqAddFriends
	var isSuccess bool

	err = json.Unmarshal(u.Ctx.Input.RequestBody, &reqBody)

	isSuccess, err = models.AddFriend(reqBody.BossId, reqBody.Phone)

	if !isSuccess {
		code = 300
	}
	u.Data["json"] = jsonPack(nil, err, code)
	u.ServeJSON()
}

// @Title setDriverStatus
// @Description 设置司机状态
// @Param	body		body 	models.ReqAddFriends	true		"body for user content"
// @Success 200 {string} logout success
// @router /deletefriend [post]
func (u *UserController) DeleteFriend() {

	//初始化参数
	err = nil
	code = 200
	var jsonMap map[string]interface{}
	var isSuccess bool

	jsonMap, err = Json2map(u.Ctx.Input.RequestBody)

	isSuccess, err = models.DeleteFriend(utils.TurnInt(jsonMap["BossId"]), utils.TurnInt(jsonMap["DriverId"]))

	if !isSuccess {
		code = 300
	}
	u.Data["json"] = jsonPack(nil, err, code)
	u.ServeJSON()
}

// @Title CreateUser
// @Description CreatedTime不用输入，系统自动生成
// @Param	body		body 	models.ReqRegisterUser	true		"body for user content"
// @Success 200 {int} models.User.Uid
// @Failure 403 body is empty
// @router /add [post]
func (u *UserController) Add() {
	var user models.ReqRegisterUser

	//初始化参数
	err = nil
	code = 200
	errr := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		err = errr
		beego.Error(errr)
	}
	isExist, id, err := models.AddUser(user)
	if isExist {
		code = 200
	} else {
		code = 201
		err = errors.New("此手机号已经注册")
	}

	content := strconv.FormatInt(id, 10)
	u.Data["json"] = jsonPackString(content, err, code)
	u.ServeJSON()
}

func Json2map(jsonByte []byte) (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal(jsonByte, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func jsonPack(objects interface{}, err error, code int) map[string]interface{} {

	var jsonMap = make(map[string]interface{})

	jsonMap["data"] = objects
	jsonMap["code"] = code
	if err != nil {
		jsonMap["msg"] = err.Error()
	} else {
		jsonMap["msg"] = "success"
	}

	return jsonMap
}

func jsonPackString(jsonStr string, err error, code int) map[string]interface{} {

	var jsonMap = make(map[string]interface{})

	jsonMap["data"] = jsonStr
	jsonMap["code"] = code
	if err != nil {
		jsonMap["msg"] = err.Error()
	} else {
		jsonMap["msg"] = "success"
	}

	return jsonMap
}

func getParamInt(u *UserController, keyStr string, defaultValue int) int {
	paramStr := u.Ctx.Input.Param(":" + keyStr)
	if paramStr == "" {
		return defaultValue
	}
	param, _ := strconv.Atoi(paramStr)
	return param
}
