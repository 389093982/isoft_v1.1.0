package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/fileutil"
	"isoft/isoft_iaas_web/models/iwork"
)

type HrefParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}


func (this *HrefParserNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	if html_str, ok := tmpDataMap["html_str?"].(string); ok{
	}
	if html_bytes, ok := tmpDataMap["html_bytes?"].([]byte); ok{
	}
}

func (this *HrefParserNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"html_str?", "html_bytes?"})
}

func (this *HrefParserNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *HrefParserNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *HrefParserNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}
