package iworkprotocol

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
)

type IWorkStep interface {
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

type OrmerProvider interface {
	GetOrmer() orm.Ormer
}
