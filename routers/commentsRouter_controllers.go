package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAllDriver",
			Router: `/getalldriver/:pageNo/:pageSize`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetIdleDriver",
			Router: `/getidledriver/:pageNo/:pageSize`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetBossIdleDriver",
			Router: `/getbossidledriver/:bossId/:pageNo/:pageSize`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetBossAllDriver",
			Router: `/getbossalldriver/:bossId/:pageNo/:pageSize`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "DriverStatus",
			Router: `/driverstatus`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "AddFriends",
			Router: `/addfriends`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "DeleteFriend",
			Router: `/deletefriend`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:UserController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:WeChatController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:WeChatController"],
		beego.ControllerComments{
			Method: "GetOpenId",
			Router: `/getOpenId/:loginCode`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:WeChatController"] = append(beego.GlobalControllerRouter["github.com/api-beego-golang/controllers:WeChatController"],
		beego.ControllerComments{
			Method: "WechatLogin",
			Router: `/wechatlogin/:code`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
