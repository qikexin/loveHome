package controllers

import (
	"github.com/astaxie/beego"
	"loveHome/models"
)

type SessionController struct {
	beego.Controller
}
func (this SessionController)RetData(resp map[string]interface{}){
	this.Data["json"]= resp
	this.ServeJSON()
}
func (this SessionController)GetSessionData(){
	resp := make(map[string]interface{})
 	//如果登录成功，返回{"name":"123124234235表示session_id"}
	defer this.RetData(resp)
	user := models.User{}
	resp["errno"] = models.RECODE_DBERR
	resp["errmsg"] = models.RecodeText(models.RECODE_DBERR)

	name := this.GetSession("name")
	if name != nil {
		user.Name = name.(string)
		resp["errno"] = models.RECODE_OK
		resp["errmsg"] = models.RecodeText(models.RECODE_OK)
		resp["data"] = user
	}
}
func(this SessionController)DeleteSessionData(){
	resp := make(map[string]interface{})
	defer this.RetData(resp)
	this.DelSession("name")

	resp["errno"] =models.RECODE_OK
	resp["errmsg"] = models.RecodeText(models.RECODE_OK)
}