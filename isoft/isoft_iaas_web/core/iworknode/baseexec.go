package iworknode

import (
	"errors"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type WorkStepFactory struct {
	Work           iwork.Work
	WorkStep       *iwork.WorkStep  // 普通步骤执行时使用的参数
	BlockStep      *block.BlockStep // 块步骤执行时使用的参数
	WorkSubRunFunc func(work iwork.Work, steps []iwork.WorkStep,
		dispatcher *entry.Dispatcher) (receiver *entry.Receiver) // 执行步骤时遇到子流程时的回调函数
	BlockStepRunFunc func(trackingId string, blockStep *block.BlockStep,
		datastore *datastore.DataStore, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) // 执行步骤时使用 BlockStep 时的回调函数
	Dispatcher *entry.Dispatcher
	Receiver   *entry.Receiver // 代理了 Receiver,值从 work_end 节点获取
	DataStore  *datastore.DataStore
}

type IStandardWorkStep interface {
	// 节点执行的方法
	Execute(trackingId string)
	// 获取默认输入参数
	GetDefaultParamInputSchema() *schema.ParamInputSchema
	// 获取动态输入参数
	GetRuntimeParamInputSchema() *schema.ParamInputSchema
	// 获取默认输出参数
	GetDefaultParamOutputSchema() *schema.ParamOutputSchema
	// 获取动态输出参数
	GetRuntimeParamOutputSchema() *schema.ParamOutputSchema
	// 节点定制化校验函数,校验不通过会触发 panic
	ValidateCustom() (checkResult []string)
}

func (this *WorkStepFactory) Execute(trackingId string) {
	proxy := this.getProxy()
	proxy.Execute(trackingId)
	if endNode, ok := proxy.(*WorkEndNode); ok {
		this.Receiver = endNode.Receiver
	}
}

func (this *WorkStepFactory) getProxy() IStandardWorkStep {
	switch strings.ToUpper(this.WorkStep.WorkStepType) {
	case "WORK_START":
		return &WorkStartNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}, Dispatcher: this.Dispatcher}
	case "WORK_END":
		return &WorkEndNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}, Receiver: this.Receiver}
	case "WORK_SUB":
		return &WorkSub{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}, WorkSubRunFunc: this.WorkSubRunFunc}
	case "SQL_EXECUTE":
		return &SQLExecuteNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "SQL_QUERY":
		return &SQLQueryNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "SQL_QUERY_PAGE":
		return &SQLQueryPageNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "JSON_RENDER":
		return &JsonRenderNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "JSON_PARSER":
		return &JsonParserNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "HTTP_REQUEST":
		return &HttpRequestNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "MAPPER":
		return &MapperNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "FILE_READ":
		return &FileReadNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "FILE_WRITE":
		return &FileWriteNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "FILE_RENAME":
		return &FileRenameNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "HREF_PARSER":
		return &HrefParserNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "ENTITY_PARSER":
		return &EntityParserNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "DB_PARSER":
		return &DBParserNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "MEMORYMAP_CACHE":
		return &MemoryMapCacheNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "GOTO_CONDITION":
		return &GotoConditionNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "CAL_HASH":
		return &CalHashNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "SET_ENV":
		return &SetEnvNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "GET_ENV":
		return &GetEnvNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "RUN_CMD":
		return &RunCmd{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "SFTP_UPLOAD":
		return &SftpUploadNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "SSH_SHELL":
		return &SSHShellNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "TARGZ_UNCOMPRESS":
		return &TarGzUnCompressNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "TARGZ_COMPRESS":
		return &TarGzCompressNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	case "IF":
		return &IFNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}, BlockStep: this.BlockStep, BlockStepRunFunc: this.BlockStepRunFunc}
	case "EMPTY":
		return &EmptyNode{WorkStep: this.WorkStep, BaseNode: BaseNode{DataStore: this.DataStore}}
	}
	panic(errors.New(fmt.Sprintf("[%v-%v]unsupport workStepType:%s", this.WorkStep.WorkId, this.WorkStep.WorkStepName, this.WorkStep.WorkStepType)))
}

func (this *WorkStepFactory) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	var inputSchema *schema.ParamInputSchema
	if _schema := this.getProxy().GetDefaultParamInputSchema(); _schema != nil {
		inputSchema = _schema
	} else {
		inputSchema = &schema.ParamInputSchema{}
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

func (this *WorkStepFactory) ValidateCustom() (checkResult []string) {
	return this.getProxy().ValidateCustom()
}
