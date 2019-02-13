package iworknode

import (
	"encoding/json"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
)

type JsonRenderNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}


func (this *JsonRenderNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	json_object := tmpDataMap["json_object"].([]map[string]interface{})
	bytes, err := json.Marshal(json_object)
	if err == nil{
		dataStore.CacheData(this.WorkStep.WorkStepName, "json_str", string(bytes))
	}
}

func (this *JsonRenderNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"json_object"})
}

func (this *JsonRenderNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"json_str"})
}

func (this *JsonRenderNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}



type JsonParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *JsonParserNode) Execute(trackingId string) {
}

func (this *JsonParserNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"json_str","json_fields"})
}

func (this *JsonParserNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *JsonParserNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}