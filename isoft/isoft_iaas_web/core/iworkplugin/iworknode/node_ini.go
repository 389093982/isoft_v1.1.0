package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkutil/fileutil"
	"isoft/isoft_iaas_web/models/iwork"
)

type IniReadNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *IniReadNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	file_path := tmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	var section_name string
	if _section_name, ok := tmpDataMap[iworkconst.STRING_PREFIX+"section_name?"].(string); ok {
		section_name = _section_name
	}
	key := tmpDataMap[iworkconst.STRING_PREFIX+"key"].(string)
	value, err := fileutil.ReadBeegoIniFile(file_path, section_name, key)
	if err == nil {
		this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"value", value)
	} else {
		panic(err)
	}
}

func (this *IniReadNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "file_path", "需要读取的文件路径"},
		2: {iworkconst.STRING_PREFIX + "section_name?", "section 名称,可为空"},
		3: {iworkconst.STRING_PREFIX + "key", "key 值"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *IniReadNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *IniReadNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "value"})
}

func (this *IniReadNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *IniReadNode) ValidateCustom() (checkResult []string) {
	return
}

type IniWriteNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *IniWriteNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	file_path := tmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	var section_name string
	if _section_name, ok := tmpDataMap[iworkconst.STRING_PREFIX+"section_name?"].(string); ok {
		section_name = _section_name
	}
	key := tmpDataMap[iworkconst.STRING_PREFIX+"key"].(string)
	value := tmpDataMap[iworkconst.STRING_PREFIX+"value"].(string)
	err := fileutil.WriteBeegoIniFile(file_path, section_name, key, value)
	if err != nil {
		panic(err)
	}
}

func (this *IniWriteNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "file_path", "需要读取的文件路径"},
		2: {iworkconst.STRING_PREFIX + "section_name?", "section 名称,可为空"},
		3: {iworkconst.STRING_PREFIX + "key", "key 值"},
		4: {iworkconst.STRING_PREFIX + "value", "value 值"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *IniWriteNode) GetRuntimeParamInputSchema() *iworkmodels.ParamInputSchema {
	return &iworkmodels.ParamInputSchema{}
}

func (this *IniWriteNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *IniWriteNode) GetRuntimeParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return &iworkmodels.ParamOutputSchema{}
}

func (this *IniWriteNode) ValidateCustom() (checkResult []string) {
	return
}
