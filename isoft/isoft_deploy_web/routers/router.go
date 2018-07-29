package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_deploy_web/controllers"
)

func init() {
	// 指定默认首页
	beego.Router("/", &controllers.EnvController{}, "get:List")

	beego.Router("/env/list", &controllers.EnvController{}, "get:List;post:PostList")
	beego.Router("/env/edit", &controllers.EnvController{}, "post:PostEdit")
	beego.Router("/env/connect_test", &controllers.EnvController{}, "post:ConnectonTest")
	beego.Router("/env/sync_deploy_home", &controllers.EnvController{}, "post:SyncDeployHome")

	beego.Router("/service/list", &controllers.ServiceController{}, "get:List;post:PostList")
	beego.Router("/service/edit", &controllers.ServiceController{}, "get:Edit;post:PostEdit")
	beego.Router("/service/modify", &controllers.ServiceController{}, "post:PostModify")
	beego.Router("/service/fileupload", &controllers.ServiceController{}, "post:FileUpload")
	beego.Router("/service/runDeployTask", &controllers.ServiceController{}, "post:RunDeployTask")
	beego.Router("/service/queryLastDeployStatus", &controllers.ServiceController{}, "post:QueryLastDeployStatus")
	beego.Router("/service/monitor", &controllers.ServiceController{}, "get:Monitor;post:PostMonitor")
	beego.Router("/service/editormonitor", &controllers.ServiceController{}, "post:EditorMonitor")
	beego.Router("/service/showServiceTrackingLogDetail", &controllers.ServiceController{}, "get:ShowServiceTrackingLogDetail")

	beego.Router("/docker/dockerInfoCheck", &controllers.DockerController{}, "post:DockerInfoCheck")
}
