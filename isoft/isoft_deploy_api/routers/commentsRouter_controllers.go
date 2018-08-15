package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["isoft/isoft_deploy_api/controllers:EnvController"] = append(beego.GlobalControllerRouter["isoft/isoft_deploy_api/controllers:EnvController"],
		beego.ControllerComments{
			Method:           "ConnectonTest",
			Router:           `/connect_test`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["isoft/isoft_deploy_api/controllers:EnvController"] = append(beego.GlobalControllerRouter["isoft/isoft_deploy_api/controllers:EnvController"],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["isoft/isoft_deploy_api/controllers:EnvController"] = append(beego.GlobalControllerRouter["isoft/isoft_deploy_api/controllers:EnvController"],
		beego.ControllerComments{
			Method:           "SyncDeployHome",
			Router:           `/sync_deploy_home`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
