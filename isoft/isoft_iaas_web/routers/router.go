package routers

import (
	"isoft/isoft_iaas_web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/ifile/fileUpload/", &controllers.IFileController{}, "post:FileUpload")
	beego.Router("/api/ifile/filterPageMetadatas/", &controllers.IFileController{}, "post:FilterPageMetadatas")
}
