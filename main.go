package main

import (
	_ "github.com/LittleRobotRoad/api-beego-golang/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/LittleRobotRoad/api-beego-golang/models"
	"github.com/astaxie/beego/context"
)

func init() {
	initDB()
}

//bee run -gendoc=true -downdoc=true
//bee pack -be GOOS=linux
func initDB() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.99.100:3306)/goweb?charset=utf8")
	// register model
	orm.RegisterModel(new(models.Object), new(models.DriverUser), new(models.BossUser), new(models.BossToDriver))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	//开启自动化文档
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//加入中间件拦截
	beego.InsertFilter("/v2/*", beego.BeforeExec, FilterBefore, false)
	beego.InsertFilter("/v2/*", beego.AfterExec, FilterAfter, false)

	beego.Run()

}

var FilterBefore = func(ctx *context.Context) {
	beego.Debug(ctx.Input.Method())

	////Post请求解析对应报文头抽取内容装入body
	//if ctx.Input.IsPost() {
	//	result := new(models.Result)
	//	json.Unmarshal(ctx.Input.RequestBody, &result)
	//	resultObject, _ := json.Marshal(result.Object)
	//	ctx.Input.RequestBody = resultObject
	//}
}

var FilterAfter = func(ctx *context.Context) {
}
