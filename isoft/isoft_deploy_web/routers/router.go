package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_deploy_web/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/api/v1/env/edit", &controllers.EnvController{}, "post:PostEdit")
	beego.Router("/api/v1/env/all", &controllers.EnvController{}, "post:All")
	beego.Router("/api/v1/env/list", &controllers.EnvController{}, "post:List")
	beego.Router("/api/v1/env/connect_test", &controllers.EnvController{}, "post:ConnectonTest")
	beego.Router("/api/v1/env/sync_deploy_home", &controllers.EnvController{}, "post:SyncDeployHome")
	beego.Router("/api/v1/service/getServiceTrackingLogDetail", &controllers.ServiceController{}, "post:GetServiceTrackingLogDetail")
	beego.Router("/api/v1/service/list", &controllers.ServiceController{}, "post:List")
	beego.Router("/api/v1/service/edit", &controllers.ServiceController{}, "post:Edit")
	beego.Router("/api/v1/service/fileDownload", &controllers.ServiceController{}, "get:FileDownload")
	beego.Router("/api/v1/service/fileUpload", &controllers.ServiceController{}, "post:FileUpload")
	beego.Router("/api/v1/service/runDeployTask", &controllers.ServiceController{}, "post:RunDeployTask")
	beego.Router("/api/v1/service/queryLastDeployStatus", &controllers.ServiceController{}, "post:QueryLastDeployStatus")

}
