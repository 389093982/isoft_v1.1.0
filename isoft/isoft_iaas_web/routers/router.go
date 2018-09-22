package routers

import (
	"isoft/isoft_iaas_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/ifile/fileUpload/", &controllers.IFileController{}, "post:FileUpload")
	beego.Router("/api/ifile/filterPageMetadatas/", &controllers.IFileController{}, "post:FilterPageMetadatas")
	beego.Router("/api/ifile/locateShards/", &controllers.IFileController{}, "post:LocateShards")
	beego.Router("/api/ifile/fileDownload/", &controllers.IFileController{}, "get:FileDownload")
}
