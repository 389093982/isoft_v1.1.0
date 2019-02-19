package iworknode

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
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
	if err == nil {
		dataStore.CacheData(this.WorkStep.WorkStepName, "json_str", string(bytes))
	}
}

func (this *JsonRenderNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1:[]string{"json_object","需要传入json对象"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *JsonRenderNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
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
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	json_str := tmpDataMap["json_str"].(string)
	json_objects := make([]map[string]interface{},0)
	err := json.Unmarshal([]byte(json_str), &json_objects)
	if err == nil {
		dataStore.CacheData(this.WorkStep.WorkStepName, "rows", json_objects)
		for index, json_object := range json_objects {
			for paramName, paramValue := range json_object {
				dataStore.CacheData(this.WorkStep.WorkStepName, fmt.Sprintf("rows[%d].%s", index, paramName), paramValue)
				if index == 0 {
					dataStore.CacheData(this.WorkStep.WorkStepName, fmt.Sprintf("rows.%s", paramName), paramValue)
				}
			}
		}
	}
}

func (this *JsonParserNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1:[]string{"json_str","需要转换成json对象的字符串"},
		2:[]string{"json_fields","json对象的字段列表"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *JsonParserNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *JsonParserNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *JsonParserNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	items := []schema.ParamOutputSchemaItem{}
	if json_fields := param.GetStaticParamValue("json_fields", this.WorkStep); strings.TrimSpace(json_fields) != "" {
		jsonArr := strings.Split(json_fields, ",")
		for _, paramName := range jsonArr {
			if _paramName := strings.TrimSpace(paramName); _paramName != "" {
				items = append(items, schema.ParamOutputSchemaItem{
					ParentPath: "rows",
					ParamName:  _paramName,
				})
			}
		}
	}
	return &schema.ParamOutputSchema{ParamOutputSchemaItems: items}
}