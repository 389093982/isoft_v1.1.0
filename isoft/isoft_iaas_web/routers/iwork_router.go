package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_iaas_web/controllers/iwork"
	"strings"
)

func loadIWorkerRouter() {
	if strings.Contains(beego.AppConfig.String("open.modules"), "iwork") {
		loadloadIWorkerRouterDetail()
	}
}

func loadloadIWorkerRouterDetail() {
	beego.Router("/api/iwork/addQuartz", &iwork.WorkController{}, "post:AddQuartz")
	beego.Router("/api/iwork/filterPageQuartz", &iwork.WorkController{}, "post:FilterPageQuartz")
	beego.Router("/api/iwork/editQuartz", &iwork.WorkController{}, "post:EditQuartz")
	beego.Router("/api/iwork/addResource", &iwork.WorkController{}, "post:AddResource")
	beego.Router("/api/iwork/filterPageResource", &iwork.WorkController{}, "post:FilterPageResource")
	beego.Router("/api/iwork/deleteResource", &iwork.WorkController{}, "post:DeleteResource")
	beego.Router("/api/iwork/validateResource", &iwork.WorkController{}, "post:ValidateResource")

	beego.Router("/api/iwork/saveHistory", &iwork.WorkController{}, "post:SaveHistory")
	beego.Router("/api/iwork/filterPageWorkHistory", &iwork.WorkController{}, "post:FilterPageWorkHistory")
	beego.Router("/api/iwork/filterPageWork", &iwork.WorkController{}, "post:FilterPageWork")
	beego.Router("/api/iwork/editWork", &iwork.WorkController{}, "post:EditWork")
	beego.Router("/api/iwork/deleteWorkById", &iwork.WorkController{}, "post:DeleteWorkById")
	beego.Router("/api/iwork/filterWorkStep", &iwork.WorkController{}, "post:FilterWorkStep")
	beego.Router("/api/iwork/addWorkStep", &iwork.WorkController{}, "post:AddWorkStep")
	beego.Router("/api/iwork/editWorkStepBaseInfo", &iwork.WorkController{}, "post:EditWorkStepBaseInfo")
	beego.Router("/api/iwork/editWorkStepParamInfo", &iwork.WorkController{}, "post:EditWorkStepParamInfo")
	beego.Router("/api/iwork/deleteWorkStepByWorkStepId", &iwork.WorkController{}, "post:DeleteWorkStepByWorkStepId")
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
	beego.Router("/api/iwork/buildIWorkDL", &iwork.WorkController{}, "post:BuildIWorkDL")
	beego.Router("/api/iwork/validateAllWork", &iwork.WorkController{}, "post:ValidateAllWork")
	beego.Router("/api/iwork/loadValidateResult", &iwork.WorkController{}, "post:LoadValidateResult")
	beego.Router("/api/iwork/refactorWorkStepInfo", &iwork.WorkController{}, "post:RefactorWorkStepInfo")
	beego.Router("/api/iwork/batchChangeIndent", &iwork.WorkController{}, "post:BatchChangeIndent")

	beego.Router("/api/iwork/submitMigrate", &iwork.WorkController{}, "post:SubmitMigrate")
	beego.Router("/api/iwork/filterPageMigrate", &iwork.WorkController{}, "post:FilterPageMigrate")
	beego.Router("/api/iwork/getMigrateInfo", &iwork.WorkController{}, "post:GetMigrateInfo")
	beego.Router("/api/iwork/executeMigrate", &iwork.WorkController{}, "post:ExecuteMigrate")
	beego.Router("/api/iwork/buildInstanceSql", &iwork.WorkController{}, "post:BuildInstanceSql")

	beego.Router("/api/iwork/editWorkVar", &iwork.WorkController{}, "post:EditWorkVar")
	beego.Router("/api/iwork/loadAllWorkVar", &iwork.WorkController{}, "post:LoadAllWorkVar")

	beego.Router("/api/iwork/loadQuickSqlMeta", &iwork.WorkController{}, "post:LoadQuickSqlMeta")

}
