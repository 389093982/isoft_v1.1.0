package iworknode

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkplugin/iworkprotocol"
	"isoft/isoft_iaas_web/core/iworkutil/reflectutil"
	"isoft/isoft_iaas_web/models/iwork"
	"reflect"
	"strings"
)

type WorkStepFactory struct {
	Work           iwork.Work
	WorkStep       *iwork.WorkStep  // 普通步骤执行时使用的参数
	BlockStep      *block.BlockStep // 块步骤执行时使用的参数
	WorkSubRunFunc func(work iwork.Work, steps []iwork.WorkStep,
		dispatcher *entry.Dispatcher) (receiver *entry.Receiver) // 执行步骤时遇到子流程时的回调函数
	BlockStepRunFunc func(trackingId string, blockStep *block.BlockStep,
		datastore *datastore.DataStore, dispatcher *entry.Dispatcher, receiver *entry.Receiver) // 执行步骤时使用 BlockStep 时的回调函数
	Dispatcher *entry.Dispatcher
	Receiver   *entry.Receiver // 代理了 Receiver,值从 work_end 节点获取
	DataStore  *datastore.DataStore
	O          orm.Ormer
}

func (this *WorkStepFactory) Execute(trackingId string) {
	proxy := this.getProxy()
	proxy.Execute(trackingId)
	if endNode, ok := proxy.(*WorkEndNode); ok {
		this.Receiver = endNode.Receiver
	}
}

func GetIWorkStep(workStepType string) iworkprotocol.IWorkStep {
	// 调整 workStepType
	_workStepType := strings.ToUpper(strings.Replace(workStepType, "_", "", -1) + "NODE")
	if t, ok := typeMap[_workStepType]; ok {
		return reflect.New(t).Interface().(iworkprotocol.IWorkStep)
	}
	panic(fmt.Sprintf("invalid workStepType for %s", workStepType))
}

func (this *WorkStepFactory) getProxy() iworkprotocol.IWorkStep {
	fieldMap := map[string]interface{}{
		"WorkStep":         this.WorkStep,
		"BaseNode":         BaseNode{DataStore: this.DataStore, o: this.O},
		"Dispatcher":       this.Dispatcher,
		"Receiver":         this.Receiver,
		"WorkSubRunFunc":   this.WorkSubRunFunc,
		"BlockStep":        this.BlockStep,
		"BlockStepRunFunc": this.BlockStepRunFunc,
	}
	stepNode := GetIWorkStep(this.WorkStep.WorkStepType)
	if stepNode == nil {
		panic(errors.New(fmt.Sprintf("[%v-%v]unsupport workStepType:%s",
			this.WorkStep.WorkId, this.WorkStep.WorkStepName, this.WorkStep.WorkStepType)))
	}
	// 从 map 中找出属性值赋值给对象
	reflectutil.FillFieldValueToStruct(stepNode, fieldMap)
	return stepNode
}

func (this *WorkStepFactory) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	var inputSchema *iworkmodels.ParamInputSchema
	if _schema := this.getProxy().GetDefaultParamInputSchema(); _schema != nil {
		inputSchema = _schema
	} else {
		inputSchema = &iworkmodels.ParamInputSchema{}
	}
	return inputSchema
}

func (this *WorkStepFactory) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	if schema := this.getProxy().GetRuntimeParamInputSchema(); schema != nil {
		return schema
	}
	return &iworkmodels.ParamInputSchema{}
}

func (this *WorkStepFactory) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	if schema := this.getProxy().GetDefaultParamOutputSchema(); schema != nil {
		return schema
	}
	return &iworkmodels.ParamOutputSchema{}
}

func (this *WorkStepFactory) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	if schema := this.getProxy().GetRuntimeParamOutputSchema(); schema != nil {
		return schema
	}
	return &iworkmodels.ParamOutputSchema{}
}

func (this *WorkStepFactory) ValidateCustom() (checkResult []string) {
	return this.getProxy().ValidateCustom()
}
