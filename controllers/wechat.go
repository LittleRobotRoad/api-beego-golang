package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/astaxie/beego/httplib"
)

// Operations about WeChat
type WeChatController struct {
	beego.Controller
}

func (c *WeChatController) URLMapping() {
	c.Mapping("openId", c.getOpenId)
}

// @Title 微信获取OpenId的接口
// @Description delete the user
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /getOpenId/:loginCode [get]
func (u *WeChatController) getOpenId() {

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
	beego.Debug(str)
	u.Data["json"] = str
	u.ServeJSON()
}
