package iworknode

import (
	"github.com/pkg/errors"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkSub struct {
	BaseNode
	WorkStep *iwork.WorkStep
	RunFunc  func(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
}

func (this *WorkSub) Execute(trackingId string) {
	// 获取子流程流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(
		schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep:this.WorkStep}))
	if strings.TrimSpace(workSubName) == ""{
		panic(errors.New("invalid workSubName"))
	}
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)

	// 运行子流程
	work, _ := iwork.QueryWorkByName(workSubName)
	steps, _ := iwork.GetAllWorkStepByWorkName(workSubName)
	receiver := this.RunFunc(work, steps, &entry.Dispatcher{TrackingId:trackingId, TmpDataMap:tmpDataMap})
	// 接收子流程数据存入 dataStore
	for paramName, paramValue := range receiver.TmpDataMap{
		dataStore.CacheData(this.WorkStep.WorkStepName, paramName, paramValue)
	}
}

func (this *WorkSub) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"work_sub"})
}
func (this *WorkSub) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}
func (this *WorkSub) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}
