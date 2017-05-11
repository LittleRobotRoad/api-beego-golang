package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/httplib"
	"encoding/json"
	"errors"
	"strconv"
	"github.com/LittleRobotRoad/api-beego-golang/models"
)

// Operations about WeChat
type WeChatController struct {
	beego.Controller
}

func (c *WeChatController) URLMapping() {
	c.Mapping("getOpenId", c.GetOpenId)
	c.Mapping("wechatlogin", c.WechatLogin)
}

// @Title 微信获取唯一识别id
// @Description delete the user
// @Param	loginCode		path 	string	true		"loginCode WeChat"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /getOpenId/:loginCode [get]
func (u *WeChatController) GetOpenId() {

	var wxUrl string = "https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code"
	var appId string = "wx13895157528f4796"

	loginCode := u.Ctx.Input.Param(":loginCode")

	wxUrl = strings.Replace(wxUrl, "APPID", appId, -1)
	wxUrl = strings.Replace(wxUrl, "SECRET", "dfd6dfe7a2510b0a2f68a04f92cdcd18", -1)
	wxUrl = strings.Replace(wxUrl, "JSCODE", loginCode, -1)

	req := httplib.Get(wxUrl)
	str, err := req.String()
	if err != nil {
		beego.Error(err)
	}
	u.Data["json"] = str
	u.ServeJSON()
}


// @Title CreateUser
// @Description CreatedTime不用输入，系统自动生成
// @Param	body		body 	models.ReqRegisterUser	true		"body for user content"
// @Success 200 {int} models.User.Uid
// @Failure 403 body is empty
// @router /wechatlogin/:code [get]
func (u *WeChatController) WechatLogin() {
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
		err = errors.New("user is Exist")
	}

	content := strconv.FormatInt(id, 10)
	u.Data["json"] = jsonPackString(content, err, code)
	u.ServeJSON()
}
