// @APIVersion 1.0.2
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact daisukeayanami@gmail.com
// @TermsOfServiceUrl http://niconico.lol/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html

package routers

import (
	"niconico.lol/driver/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Debug("router init")

	ns := beego.NewNamespace("/v2",
		beego.NSNamespace("/object", beego.NSInclude(&controllers.ObjectController{})),
		beego.NSNamespace("/user", beego.NSInclude(&controllers.UserController{})),
		beego.NSNamespace("/wechat", beego.NSInclude(&controllers.WeChatController{})),
	)
	beego.AddNamespace(ns)
}
