package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/models/iwork"
	"os"
)

type GetEnvNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *GetEnvNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	env_var_value := os.Getenv(tmpDataMap[iworkconst.STRING_PREFIX+"env_var_name"].(string))
	dataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"env_var_value", env_var_value)
}

func (this *GetEnvNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "env_var_name", "环境变量名称"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *GetEnvNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *GetEnvNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "env_var_value"})
}

func (this *GetEnvNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *GetEnvNode) ValidateCustom() (checkResult []string) {
	return
}

type SetEnvNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SetEnvNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	env_var_name := tmpDataMap[iworkconst.STRING_PREFIX+"env_var_name"].(string)
	env_var_value := tmpDataMap[iworkconst.STRING_PREFIX+"env_var_value"].(string)
	if err := os.Setenv(env_var_name, env_var_value); err != nil {
		panic(err)
	}
}

func (this *SetEnvNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "env_var_name", "环境变量名称"},
		2: {iworkconst.STRING_PREFIX + "env_var_value", "环境变量值"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SetEnvNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *SetEnvNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *SetEnvNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *SetEnvNode) ValidateCustom() (checkResult []string) {
	return
}
