package iworknode

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkStepFactory struct {
	Work     	iwork.Work
	WorkStep 	*iwork.WorkStep
	// 执行 Execute 方法时遇到子流程时的回调函数
	RunFunc  	func(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
	Dispatcher 	*entry.Dispatcher
	Receiver   	*entry.Receiver				// 代理了 Receiver,值从 work_end 节点获取
}

type IStandardWorkStep interface {
	Execute(trackingId string)
	GetDefaultParamInputSchema() *schema.ParamInputSchema
	GetDefaultParamOutputSchema() *schema.ParamOutputSchema
	GetRuntimeParamOutputSchema() *schema.ParamOutputSchema
}

func (this *WorkStepFactory) Execute(trackingId string) {
	proxy := this.getProxy()
	proxy.Execute(trackingId)
	if endNode, ok := proxy.(*WorkEndNode); ok{
		this.Receiver = endNode.Receiver
	}
}

func (this *WorkStepFactory) getProxy() IStandardWorkStep {
	switch strings.ToUpper(this.WorkStep.WorkStepType) {
	case "WORK_START":
		return &WorkStartNode{WorkStep: this.WorkStep, Dispatcher:this.Dispatcher}
	case "WORK_END":
		return &WorkEndNode{WorkStep: this.WorkStep, Receiver:this.Receiver}
	case "WORK_SUB":
		return &WorkSub{WorkStep: this.WorkStep, RunFunc: this.RunFunc}
	case "SQL_EXECUTE":
		return &SQLExecuteNode{WorkStep: this.WorkStep}
	case "SQL_QUERY":
		return &SQLQueryNode{WorkStep: this.WorkStep}
	}
	panic(fmt.Sprintf("unsupport workStepType:%s", this.WorkStep.WorkStepType))
}

func (this *WorkStepFactory) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	if schema := this.getProxy().GetDefaultParamInputSchema(); schema != nil {
		return schema
	}
	return &schema.ParamInputSchema{}
}

func (this *WorkStepFactory) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	if schema := this.getProxy().GetDefaultParamOutputSchema(); schema != nil {
		return schema
	}
	return &schema.ParamOutputSchema{}
}

func (this *WorkStepFactory) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	if schema := this.getProxy().GetRuntimeParamOutputSchema(); schema != nil {
		return schema
	}
	return &schema.ParamOutputSchema{}
}
