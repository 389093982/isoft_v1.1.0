package iworknode

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkSubNode struct {
	BaseNode
	WorkStep       *iwork.WorkStep
	WorkSubRunFunc func(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
}

func (this *WorkSubNode) Execute(trackingId string) {
	// 获取子流程流程名称
	workSubName := this.checkAndGetWorkSubName()
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	// 运行子流程
	work, _ := iwork.QueryWorkByName(workSubName, orm.NewOrm())
	steps, _ := iwork.QueryAllWorkStepByWorkName(workSubName, orm.NewOrm())
	this.RunOnceSubWork(work, steps, trackingId, tmpDataMap, this.DataStore)
}

func (this *WorkSubNode) checkAndGetWorkSubName() string {
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(
		schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep}))
	if strings.TrimSpace(workSubName) == "" {
		panic(errors.New("invalid workSubName"))
	}
	return workSubName
}

func (this *WorkSubNode) RunOnceSubWork(work iwork.Work, steps []iwork.WorkStep, trackingId string,
	tmpDataMap map[string]interface{}, dataStore *datastore.DataStore) {
	receiver := this.WorkSubRunFunc(work, steps, &entry.Dispatcher{TrackingId: trackingId, TmpDataMap: tmpDataMap})
	// 接收子流程数据存入 dataStore
	for paramName, paramValue := range receiver.TmpDataMap {
		dataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{paramName: paramValue})
	}
}

func (this *WorkSubNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "work_sub", "子流程信息"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

// 获取动态输入值
func (this *WorkSubNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	items := make([]iworkmodels.ParamInputSchemaItem, 0)
	// 获取子流程信息
	workSubName := this.getWorkSubName()
	if strings.TrimSpace(workSubName) != "" {
		// 获取子流程所有步骤
		subSteps, err := iwork.QueryAllWorkStepByWorkName(workSubName, this.GetOrmer())
		if err != nil {
			panic(err)
		}
		for _, subStep := range subSteps {
			// 找到子流程起始节点
			if strings.ToUpper(subStep.WorkStepType) == "WORK_START" {
				// 子流程起始节点输入参数
				subItems := schema.GetCacheParamInputSchema(&subStep, &WorkStepFactory{WorkStep: &subStep})
				for _, subItem := range subItems.ParamInputSchemaItems {
					items = append(items, iworkmodels.ParamInputSchemaItem{ParamName: subItem.ParamName})
				}
			}
		}
	}
	return &iworkmodels.ParamInputSchema{ParamInputSchemaItems: items}
}

func (this *WorkSubNode) getWorkSubName() string {
	// 读取历史输入值
	paramInputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	// 从历史输入值中获取子流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
	return workSubName
}

func (this *WorkSubNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := make([]iworkmodels.ParamOutputSchemaItem, 0)
	// 读取静态输入值
	paramInputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
	// 从静态输入值中获取子流程名称
	workSubName := iworkutil.GetWorkSubNameForWorkSubNode(paramInputSchema)
	if strings.TrimSpace(workSubName) != "" {
		// 获取子流程所有步骤
		subSteps, err := iwork.QueryAllWorkStepByWorkName(workSubName, orm.NewOrm())
		if err != nil {
			panic(err)
		}
		for _, subStep := range subSteps {
			// 找到子流程结束节点
			if strings.ToUpper(subStep.WorkStepType) == "WORK_END" {
				// 子流程结束节点输出参数
				subItems := schema.GetCacheParamOutputSchema(&subStep)
				for _, subItem := range subItems.ParamOutputSchemaItems {
					items = append(items, iworkmodels.ParamOutputSchemaItem{ParamName: subItem.ParamName})
				}
			}
		}
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *WorkSubNode) ValidateCustom() (checkResult []string) {
	workSubName := this.getWorkSubName()
	if workSubName == "" {
		checkResult = append(checkResult, fmt.Sprintf("Empty workSubName was found!"))
		return
	}
	work, err := iwork.QueryWorkByName(workSubName, orm.NewOrm())
	if err != nil {
		checkResult = append(checkResult, fmt.Sprintf("WorkSubName for %s was not found!", workSubName))
		return
	}
	if startStep, err := iwork.QueryWorkStepByStepName(work.Id, "start", orm.NewOrm()); err == nil {
		workSubInputSchema := schema.GetCacheParamInputSchema(this.WorkStep, &WorkStepFactory{WorkStep: this.WorkStep})
		inputSchema := schema.GetCacheParamInputSchema(&startStep, &WorkStepFactory{WorkStep: &startStep})
		for _, item := range inputSchema.ParamInputSchemaItems {
			exist := false
			for _, workSubItem := range workSubInputSchema.ParamInputSchemaItems {
				if item.ParamName == workSubItem.ParamName {
					exist = true
					break
				}
			}
			if !exist {
				checkResult = append(checkResult, fmt.Sprintf("Miss paramName for %s was found!", item.ParamName))
			}
		}
	} else {
		checkResult = append(checkResult, fmt.Sprintf("Miss start node for worksub %s was found!", work.WorkName))
	}
	return
}
