package iwork

import (
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/core/iworkrun"
	"isoft/isoft_iaas_web/models/iwork"
)

// 示例地址: http://localhost:8086/api/iwork/httpservice/test_iblog_table_migrate?author0=admin1234567
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
	mapData := this.ParseParam(steps)
	receiver := iworkrun.Run(work, steps, &entry.Dispatcher{TmpDataMap:mapData})
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "result": receiver.TmpDataMap}
	this.ServeJSON()
}

func (this *WorkController) ParseParam(steps []iwork.WorkStep) map[string]interface{} {
	mapData := map[string]interface{}{}
	for _, step := range steps{
		if step.WorkStepType == "work_start"{
			inputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep:&step})
			for _, item := range inputSchema.ParamInputSchemaItems{
				mapData[item.ParamName] = this.Input().Get(item.ParamName)
			}
		}
	}
	return mapData
}