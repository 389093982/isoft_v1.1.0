package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_deploy_web/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/env/edit", &controllers.EnvController{}, "post:PostEdit")
	beego.Router("/api/env/all", &controllers.EnvController{}, "post:All")
	beego.Router("/api/env/list", &controllers.EnvController{}, "post:List")
	beego.Router("/api/env/connect_test", &controllers.EnvController{}, "post:ConnectonTest")
	beego.Router("/api/env/sync_deploy_home", &controllers.EnvController{}, "post:SyncDeployHome")

	beego.Router("/api/service/getServiceTrackingLogDetail", &controllers.ServiceController{}, "post:GetServiceTrackingLogDetail")
	beego.Router("/api/service/list", &controllers.ServiceController{}, "post:List")
	beego.Router("/api/service/edit", &controllers.ServiceController{}, "post:Edit")
	beego.Router("/api/service/fileDownload", &controllers.ServiceController{}, "get:FileDownload")
	beego.Router("/api/service/fileUpload", &controllers.ServiceController{}, "post:FileUpload")
	beego.Router("/api/service/runDeployTask", &controllers.ServiceController{}, "post:RunDeployTask")
	beego.Router("/api/service/queryLastDeployStatus", &controllers.ServiceController{}, "post:QueryLastDeployStatus")

	beego.Router("/api/config/edit", &controllers.ConfigController{}, "post:Edit")
	beego.Router("/api/config/list", &controllers.ConfigController{}, "post:List")
	beego.Router("/api/config/fileDownload", &controllers.ConfigController{}, "get:FileDownload")
	beego.Router("/api/config/fileUpload", &controllers.ConfigController{}, "post:FileUpload")

}
