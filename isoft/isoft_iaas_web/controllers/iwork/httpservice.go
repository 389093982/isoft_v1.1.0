package iwork

import (
	"isoft/isoft_iaas_web/core/iworkrun"
	"isoft/isoft_iaas_web/models/iwork"
)

// 示例地址: http://localhost:8086/api/iwork/httpservice/test_iblog_table_migrate
func (this *WorkController) PublishAsSerivce()  {
	defer func() {
		if err := recover(); err != nil{
			this.Data["json"] = &map[string]interface{}{"status":"ERROR", "errorMsg":err.(error).Error()}
			this.ServeJSON()
		}
	}()
	work_name := this.Ctx.Input.Param(":work_name")
	work, err := iwork.QueryWorkByName(work_name)
	if err != nil{
		panic(err)
	}
	steps, err := iwork.GetAllWorkStepByWorkName(work_name)
	if err != nil{
		panic(err)
	}
	receiver := iworkrun.Run(work, steps, nil)
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "result": receiver.TmpDataMap}
	this.ServeJSON()
}
