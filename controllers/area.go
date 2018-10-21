package controllers

import (
	"github.com/astaxie/beego"
	"loveHome/models"
	"github.com/astaxie/beego/orm"
)

type AreaController struct {
	beego.Controller
}


func (c *AreaController) retResp(resp map[string]interface{})  {
	c.Data["json"] = &resp
	c.ServeJSONP()
}
func (c *AreaController) GetArea()  {
	beego.Info("connect success")
	//从session中拿数据

	resp := make(map[string]interface{})
/*	area := models.Area{}
	o := orm.NewOrm()
	err := o.Read(&area)
	if err != nil {
		beego.Info("数据错误")
		resp["errno"] = 4001
		resp["errmsg"] = "查询失败"
		//c.Data["json"] = resp
		//c.ServeJSON()
		c.retResp(resp)
		return
	}

	resp["errno"] = 0
	resp["errmsg"] ="OK"
	resp["data"] = &area
	c.retResp(resp)*/





	//resp["errno"] = 4001
	//resp["errmsg"] = "查询失败"
	////resp["errno"] = models.RECODE_OK
	////resp["errmsg"] = models.RecodeText(models.RECODE_OK)
	defer c.retResp(resp)  //因为每一种判断都需要返回结果，执行c.retResp(resp)，所以直接用defer在最后执行一次
	//
	////从redis缓存中拿数据
	//cache_conn, err := cache.NewCache("redis",`{"key":"lovehome","conn":"6379","dbNum":"0"}`)
	//if areaData := cache_conn.Get("area");areaData != nil{
	//	beego.Info("get data from cache ======")
	//	resp["data"]=areaData
	//	return
	//}
	//
	//从mysql中拿到area数据
	var areas []models.Area
	o := orm.NewOrm()
	num, err := o.QueryTable("area").All(&areas)
	if err != nil {
		resp["errno"] = 4001 //models.RECODE_DBERR
		resp["errmsg"] = "查询错误" //modes.RecodeText(models.RECODE_DBERR)

		return
	}
	if num == 0 {
		resp["errno"] = 4002
		resp["errmsg"] = "没有查到数据"
		return
	}
	resp["data"]= areas

	////把数据转换成json格式存入缓存
	//json_str,err := json.Marshal(areas)
	//if err != nil {
	//	beego.Error("encoding err")
	//	return
	//}
	//cache_conn.Put("area",json_str,time.Second*3600)
	//
	////打包成json返回给前端
	//beego.Info("query data success,resp = ",resp,"mum = ",num)
	}