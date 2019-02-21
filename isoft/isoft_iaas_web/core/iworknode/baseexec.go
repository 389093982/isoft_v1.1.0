package iworknode

import (
	"errors"
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkStepFactory struct {
	Work     iwork.Work
	WorkStep *iwork.WorkStep
	// 执行 Execute 方法时遇到子流程时的回调函数
	RunFunc    func(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver)
	Dispatcher *entry.Dispatcher
	Receiver   *entry.Receiver // 代理了 Receiver,值从 work_end 节点获取
}

type IStandardWorkStep interface {
	// 节点执行的方法
	Execute(trackingId string,skipFunc func(tmpDataMap map[string]interface{}) bool)
	// 获取默认输入参数
	GetDefaultParamInputSchema() *schema.ParamInputSchema
	// 获取动态输入参数
	GetRuntimeParamInputSchema() *schema.ParamInputSchema
	// 获取默认输出参数
	GetDefaultParamOutputSchema() *schema.ParamOutputSchema
	// 获取动态输出参数
	GetRuntimeParamOutputSchema() *schema.ParamOutputSchema
	// 节点定制化校验函数,校验不通过会触发 panic
	ValidateCustom()
}

func (this *WorkStepFactory) Execute(trackingId string) {
	skipFunc := func(tmpDataMap map[string]interface{}) bool {
		// if 节点判断, if条件判断为 false 时跳过
		if checkif, ok := tmpDataMap[iworkconst.BOOL_PREFIX + "if?"].(bool); ok && checkif == false{
			iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("The step for %s was skipped!", this.WorkStep.WorkStepName))
			return true
		}
		return false
	}

	proxy := this.getProxy()
	proxy.Execute(trackingId, skipFunc)
	if endNode, ok := proxy.(*WorkEndNode); ok {
		this.Receiver = endNode.Receiver
	}
}

func (this *WorkStepFactory) getProxy() IStandardWorkStep {
	switch strings.ToUpper(this.WorkStep.WorkStepType) {
	case "WORK_START":
		return &WorkStartNode{WorkStep: this.WorkStep, Dispatcher: this.Dispatcher}
	case "WORK_END":
		return &WorkEndNode{WorkStep: this.WorkStep, Receiver: this.Receiver}
	case "WORK_SUB":
		return &WorkSub{WorkStep: this.WorkStep, RunFunc: this.RunFunc}
	case "SQL_EXECUTE":
		return &SQLExecuteNode{WorkStep: this.WorkStep}
	case "SQL_QUERY":
		return &SQLQueryNode{WorkStep: this.WorkStep}
	case "SQL_QUERY_PAGE":
		return &SQLQueryPageNode{WorkStep: this.WorkStep}
	case "JSON_RENDER":
		return &JsonRenderNode{WorkStep: this.WorkStep}
	case "JSON_PARSER":
		return &JsonParserNode{WorkStep: this.WorkStep}
	case "HTTP_REQUEST":
		return &HttpRequestNode{WorkStep: this.WorkStep}
	case "MAPPER":
		return &MapperNode{WorkStep: this.WorkStep}
	case "FILE_READ":
		return &FileReadNode{WorkStep: this.WorkStep}
	case "FILE_WRITE":
		return &FileWriteNode{WorkStep: this.WorkStep}
	case "HREF_PARSER":
		return &HrefParserNode{WorkStep: this.WorkStep}
	case "ENTITY_PARSER":
		return &EntityParserNode{WorkStep: this.WorkStep}
	case "DB_PARSER":
		return &DBParserNode{WorkStep: this.WorkStep}
	case "MEMORYMAP_CACHE":
		return &MemoryMapCacheNode{WorkStep: this.WorkStep}
	case "GOTO_CONDITION":
		return &GotoConditionNode{WorkStep: this.WorkStep}
	case "EMPTY":
		return &EmptyNode{WorkStep: this.WorkStep}
	}
	panic(errors.New(fmt.Sprintf("[%v-%v]unsupport workStepType:%s", this.WorkStep.WorkId, this.WorkStep.WorkStepName, this.WorkStep.WorkStepType)))
}

func (this *WorkStepFactory) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	var inputSchema *schema.ParamInputSchema
	if _schema := this.getProxy().GetDefaultParamInputSchema(); _schema != nil {
		inputSchema = _schema
	}else{
		inputSchema = &schema.ParamInputSchema{}
	}
	// 不支持 if 判断的节点
	if !stringutil.CheckContains(this.WorkStep.WorkStepType, []string{"work_start","work_end"}){
		appendDefaultParamInputSchemaItem(schema.ParamInputSchemaItem{
			ParamName:iworkconst.BOOL_PREFIX + "if?",ParamDesc:"if指令,只有满足条件时才会执行,不填时必定会执行!"}, inputSchema)
	}

	// 不支持 redirect 跳转的节点
	if !stringutil.CheckContains(this.WorkStep.WorkStepType, []string{"goto_condition","redirect","work_start","work_end"}){
		appendDefaultParamInputSchemaItem(schema.ParamInputSchemaItem{
			ParamName:iworkconst.STRING_PREFIX + "redirect?",ParamDesc:"redirect指令,当前节点执行完成后,会跳往匹配上的节点继续执行!"}, inputSchema)

	}
	return inputSchema
}

func (this *WorkStepFactory) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	if schema := this.getProxy().GetRuntimeParamInputSchema(); schema != nil {
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

func (this *WorkStepFactory) ValidateCustom() {
	this.getProxy().ValidateCustom()
}

func appendDefaultParamInputSchemaItem(item schema.ParamInputSchemaItem, inputSchema *schema.ParamInputSchema)  {
	for _, _item := range inputSchema.ParamInputSchemaItems{
		if _item.ParamName == item.ParamName{
			return
		}
	}
	inputSchema.ParamInputSchemaItems = append(inputSchema.ParamInputSchemaItems, item)
}
