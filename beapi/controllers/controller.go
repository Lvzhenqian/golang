package controllers

import (
	"beapi/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func (this *Controller) Post(){
	//Name := this.GetString("name")
	//obid := models.IntoTable(Name)
	//this.Data["json"] = map[string]interface{}{
	//	"id": obid,
	//	"name": Name,
	//}
	var ob2 models.User
	ob2.Name = "test"
	var ob interface{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody,&ob); err == nil {
		beego.Debug(ob.(map[string]interface{}))
	}
	//if err := json.Unmarshal(this.Ctx.Input.RequestBody,&ob);err == nil{
	//	objectid := models.AddObject(ob)
	//	this.Data["json"] = map[string]interface{}{
	//		"ObjectId": objectid,
	//		"name": ob.Name,
	//	}
	//}else {
	//	this.Data["json"] = err.Error()
	//}
	this.ServeJSON()
}

func (this *Controller)Get()  {
	//this.Ctx.WriteString("Hello World!!")
	this.Data["json"] = map[string]string{
		"status": "ok",
	}
	this.ServeJSON()
	return
}