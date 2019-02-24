package main

import (
	"beapi/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
)

func init() {
	// 注册Driver
	if err := orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		beego.Emergency(err.Error())
	}
	// Url拼接
	drivers := strings.Join([]string{
		beego.AppConfig.String("mysql::username"),":",
		beego.AppConfig.String("mysql::password"),"@tcp(",
		beego.AppConfig.String("mysql::host"),":",
		beego.AppConfig.String("mysql::port"),")/",
		beego.AppConfig.String("mysql::db"),"?charset=utf8&parseTime=true&loc=Local",
	},"")
	// 注册 default
	if base := orm.RegisterDataBase("default", "mysql", drivers); base != nil {
		beego.Emergency(base.Error())
	}
	orm.SetMaxIdleConns("default",30)
	orm.SetMaxOpenConns("default",30)
	// 同步数据表
	if err := orm.RunSyncdb("default", false, true); err != nil{
		 beego.Error(err.Error())
	}

}
func main() {
	orm.Debug = true
	beego.BConfig.CopyRequestBody = true
	beego.Router("/hello",&controllers.Controller{})
	beego.Router("/test",&controllers.Controller{})
	beego.Run()
	//orm.RunCommand()

}
