package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/controllers/iquartz"
	"isoft/isoft_iaas_web/controllers/iresource"
	"isoft/isoft_iaas_web/controllers/iwork"
	"strings"
)

func loadIWorkerRouter()  {
	if strings.Contains(beego.AppConfig.String("open.moudles"), "iwork") {
		loadloadIWorkerRouterDetail()
	}
}

func loadloadIWorkerRouterDetail()  {
	beego.Router("/api/iquartz/addQuartz", &iquartz.QuartzController{}, "post:AddQuartz")
	beego.Router("/api/iquartz/filterPageQuartz", &iquartz.QuartzController{}, "post:FilterPageQuartz")
	beego.Router("/api/iresource/addResource", &iresource.ResourceController{}, "post:AddResource")
	beego.Router("/api/iresource/filterPageResource", &iresource.ResourceController{}, "post:FilterPageResource")

	beego.Router("/api/iwork/filterPageWork", &iwork.WorkController{}, "post:FilterPageWork")
	beego.Router("/api/iwork/editWork", &iwork.WorkController{}, "post:EditWork")
	beego.Router("/api/iwork/deleteWorkById", &iwork.WorkController{}, "post:DeleteWorkById")
	beego.Router("/api/iwork/filterPageWorkStep", &iwork.WorkController{}, "post:FilterPageWorkStep")
	beego.Router("/api/iwork/addWorkStep", &iwork.WorkController{}, "post:AddWorkStep")
	beego.Router("/api/iwork/editWorkStepBaseInfo", &iwork.WorkController{}, "post:EditWorkStepBaseInfo")
	beego.Router("/api/iwork/editWorkStepParamInfo", &iwork.WorkController{}, "post:EditWorkStepParamInfo")
	beego.Router("/api/iwork/deleteWorkStepById", &iwork.WorkController{}, "post:DeleteWorkStepById")
	beego.Router("/api/iwork/loadWorkStepInfo", &iwork.WorkController{}, "post:LoadWorkStepInfo")
	beego.Router("/api/iwork/getAllWorkStepInfo", &iwork.WorkController{}, "post:GetAllWorkStepInfo")
	beego.Router("/api/iwork/changeWorkStepOrder", &iwork.WorkController{}, "post:ChangeWorkStepOrder")
	beego.Router("/api/iwork/runWork", &iwork.WorkController{}, "post:RunWork")
	beego.Router("/api/iwork/loadPreNodeOutput", &iwork.WorkController{}, "post:LoadPreNodeOutput")
	beego.Router("/api/iwork/filterPageLogRecord", &iwork.WorkController{}, "post:FilterPageLogRecord")
	beego.Router("/api/iwork/getLastRunLogDetail", &iwork.WorkController{}, "post:GetLastRunLogDetail")
	beego.Router("/api/iwork/httpservice/:work_name", &iwork.WorkController{}, "get,post:PublishAsSerivce")
	beego.Router("/api/iwork/getRelativeWork", &iwork.WorkController{}, "post:GetRelativeWork")
	beego.Router("/api/iwork/filterPageEntity", &iwork.WorkController{}, "post:FilterPageEntity")
	beego.Router("/api/iwork/editEntity", &iwork.WorkController{}, "post:EditEntity")
	beego.Router("/api/iwork/deleteEntity", &iwork.WorkController{}, "post:DeleteEntity")
}
