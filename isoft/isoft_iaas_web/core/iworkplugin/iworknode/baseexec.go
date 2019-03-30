package iworknode

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkplugin/iworkprotocol"
	"isoft/isoft_iaas_web/models/iwork"
	"reflect"
	"strings"
)

var typeMap map[string]reflect.Type

func init() {
	typeMap = make(map[string]reflect.Type, 0)
	vs := []interface{}{
		WorkStartNode{},
		WorkEndNode{},
		WorkSubNode{},
		SQLExecuteNode{},
		SQLQueryNode{},
		SQLQueryPageNode{},
		JsonRenderNode{},
		JsonParserNode{},
		HttpRequestNode{},
		MapperNode{},
		FileReadNode{},
		FileWriteNode{},
		FileSyncNode{},
		FileDeleteNode{},
		HrefParserNode{},
		EntityParserNode{},
		DBParserNode{},
		MemoryMapCacheNode{},
		GotoConditionNode{},
		CalHashNode{},
		SetEnvNode{},
		GetEnvNode{},
		RunCmdNode{},
		SftpUploadNode{},
		SSHShellNode{},
		TarGzUnCompressNode{},
		TarGzCompressNode{},
		IniReadNode{},
		IniWriteNode{},
		IFNode{},
		EmptyNode{},
		Base64EncodeNode{},
		Base64DecodeNode{},
	}
	for _, v := range vs {
		t := reflect.ValueOf(v).Type()
		typeMap[strings.ToUpper(t.Name())] = t
	}
}

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
	fillFieldValueToStruct(stepNode, fieldMap)
	return stepNode
}

// 将结构体里的成员按照字段名字来赋值
func fillFieldValueToStruct(ptr interface{}, fields map[string]interface{}) {
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		if value, ok := fields[fieldInfo.Name]; ok {
			//给结构体赋值
			//保证赋值时数据类型一致
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			}
		}
	}
	return
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
