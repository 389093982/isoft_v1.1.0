package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) RenderJsonErrorWithInvalidParam() {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "参数不合法!"}
	this.ServeJSON()
}

func (this *BaseController) RenderJsonErrorWithInvalidParamDetail(detail string) {
	this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": detail}
	this.ServeJSON()
}

func (this *BaseController) RenderJsonSuccessWithResultMap(jsonMap map[string]interface{}) {
	jsonMap["status"] = "SUCCESS"
	this.Data["json"] = &jsonMap
	this.ServeJSON()
}
