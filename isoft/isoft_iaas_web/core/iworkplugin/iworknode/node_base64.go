package iworknode

import (
	"encoding/base64"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
)

type Base64EncodeNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *Base64EncodeNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	input := tmpDataMap[iworkconst.STRING_PREFIX+"input"].(string)
	encodeString := base64.StdEncoding.EncodeToString([]byte(input))
	this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"encode_data", encodeString)
}

func (this *Base64EncodeNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "input", "待编码的输入字符串"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *Base64EncodeNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *Base64EncodeNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "encode_data"})
}

func (this *Base64EncodeNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *Base64EncodeNode) ValidateCustom() (checkResult []string) {
	return
}

type Base64DecodeNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *Base64DecodeNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	input := tmpDataMap[iworkconst.STRING_PREFIX+"input"].(string)
	bytes, err := base64.StdEncoding.DecodeString(input)
	if err == nil {
		this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"decode_data", string(bytes))
		this.DataStore.CacheByteData(this.WorkStep.WorkStepName, iworkconst.BYTE_ARRAY_PREFIX+"decode_data", bytes)
	} else {
		panic(err)
	}
}

func (this *Base64DecodeNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "input", "待解码的输入字符串"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *Base64DecodeNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *Base64DecodeNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "decode_data", iworkconst.BYTE_ARRAY_PREFIX + "decode_data"})
}

func (this *Base64DecodeNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *Base64DecodeNode) ValidateCustom() (checkResult []string) {
	return
}