package iworknode

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type JsonRenderNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *JsonRenderNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	json_object := tmpDataMap[iworkconst.COMPLEX_PREFIX+"json_data"].([]map[string]interface{})
	bytes, err := json.Marshal(json_object)
	if err == nil {
		this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"json_data", string(bytes))
	}
}

func (this *JsonRenderNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.COMPLEX_PREFIX + "json_data", "需要传入json对象"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *JsonRenderNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *JsonRenderNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "json_data"})
}

func (this *JsonRenderNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *JsonRenderNode) ValidateCustom() (checkResult []string) {
	return []string{}
}

type JsonParserNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *JsonParserNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	json_str := tmpDataMap[iworkconst.STRING_PREFIX+"json_data"].(string)
	json_objects := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(json_str), &json_objects)
	if err == nil {
		this.DataStore.CacheData(this.WorkStep.WorkStepName, "rows", json_objects)
		for index, json_object := range json_objects {
			for paramName, paramValue := range json_object {
				this.DataStore.CacheData(this.WorkStep.WorkStepName, fmt.Sprintf("rows[%d].%s", index, paramName), paramValue)
				if index == 0 {
					this.DataStore.CacheData(this.WorkStep.WorkStepName, fmt.Sprintf("rows.%s", paramName), paramValue)
				}
			}
		}
	}
}

func (this *JsonParserNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "json_data", "需要转换成json对象的字符串"},
		2: {"json_fields", "json对象的字段列表"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *JsonParserNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *JsonParserNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *JsonParserNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	items := []iworkmodels.ParamOutputSchemaItem{}
	if json_fields := param.GetStaticParamValue("json_fields", this.WorkStep).(string); strings.TrimSpace(json_fields) != "" {
		jsonArr := strings.Split(json_fields, ",")
		for _, paramName := range jsonArr {
			if _paramName := strings.TrimSpace(paramName); _paramName != "" {
				items = append(items, iworkmodels.ParamOutputSchemaItem{
					ParentPath: "rows",
					ParamName:  _paramName,
				})
			}
		}
	}
	return &iworkmodels.ParamOutputSchema{ParamOutputSchemaItems: items}
}

func (this *JsonParserNode) ValidateCustom() (checkResult []string) {
	return
}