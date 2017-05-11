// @APIVersion 1.0.2
// @Title 司机预约 API
// @Description 主要是司机用户和老板用户注册查询等接口的示例
// @Contact daisukeayanami@gmail.com
// @TermsOfServiceUrl http://niconico.lol/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html

package routers

import (
	"github.com/LittleRobotRoad/api-beego-golang/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Debug("router init")

	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/object", beego.NSInclude(&controllers.ObjectController{})),
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/wechat", beego.NSInclude(&controllers.WeChatController{})))
	beego.AddNamespace(ns)
}
