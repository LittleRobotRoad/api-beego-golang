package main

import (
	_ "niconico.lol/driver/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"niconico.lol/driver/models"
	"github.com/astaxie/beego/context"
)

func init() {
	initDB()
}

//bee run -gendoc=true -downdoc=true
//bee pack -be GOOS=linux
func initDB() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/goweb?charset=utf8")
	// register model
	orm.RegisterModel(new(models.Object), new(models.User))

	// create table
	orm.RunSyncdb("default", false, true)
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

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
