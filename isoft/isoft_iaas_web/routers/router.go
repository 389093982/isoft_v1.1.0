package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/ifile/fileUpload/", &controllers.IFileController{}, "post:FileUpload")
	beego.Router("/api/ifile/filterPageMetadatas/", &controllers.IFileController{}, "post:FilterPageMetadatas")
	beego.Router("/api/ifile/locateShards/", &controllers.IFileController{}, "post:LocateShards")
	beego.Router("/api/ifile/fileDownload/", &controllers.IFileController{}, "get:FileDownload")
	beego.Router("/api/heartbeat/sendHeartBeat/", &controllers.HeartBeatController{}, "post:SendHeartBeat")
	beego.Router("/api/heartbeat/queryAllAliveHeartBeat/", &controllers.HeartBeatController{}, "get:QueryAllAliveHeartBeat")
}
