package controllers

import (
	"github.com/astaxie/beego"
)

type IFileController struct {
	beego.Controller
}

func (this *IFileController) FileUpload() {
	//service_id, _ := this.GetInt64("service_id")
	//serviceInfo, _ := models.QueryServiceInfoById(service_id)
	//_, h, err := this.GetFile("file")
	//if err != nil {
	//	this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
	//} else {
	//	// 根据 service_id 创建分级文件夹
	//	os.MkdirAll("static/uploadfile/"+serviceInfo.ServiceName, os.ModePerm)
	//	// 保存文件
	//	err := this.SaveToFile("file", path.Join(SFTP_SRC_DIR+"/static/uploadfile/"+serviceInfo.ServiceName, h.Filename))
	//	if err != nil {
	//		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
	//	} else {
	//		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	//	}
	//}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	this.ServeJSON()
}
