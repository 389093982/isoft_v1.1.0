package iworknode

import (
	"io/ioutil"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/fileutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

type FileReadNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}


func (this *FileReadNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	file_path := tmpDataMap["file_path"].(string)
	dataStore.CacheData(this.WorkStep.WorkStepName, "file_path", file_path)
	if bytes, err := ioutil.ReadFile(file_path); err == nil{
		dataStore.CacheData(this.WorkStep.WorkStepName, "data", string(bytes))
	}else{
		panic(err)
	}
}

func (this *FileReadNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"file_path"})
}

func (this *FileReadNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *FileReadNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"file_path", "data"})
}

func (this *FileReadNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}



type FileWriteNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func checkAppend(tmpDataMap map[string]interface{}) bool {
	if append,ok := tmpDataMap["append?"].(string); ok && strings.TrimSpace(append) != ""{
		return true
	}
	return false
}

func (this *FileWriteNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataSource(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	file_path := tmpDataMap["file_path"].(string)
	// 写字符串
	if data, ok := tmpDataMap["data?"].(string); ok{
		if err := fileutil.WriteFile(file_path, []byte(data), checkAppend(tmpDataMap)); err != nil{
			panic(err)
		}
	}
	// 写字节数组
	if bytes, ok := tmpDataMap["bytes?"].([]byte); ok{
		if err := fileutil.WriteFile(file_path, bytes, checkAppend(tmpDataMap)); err != nil{
			panic(err)
		}
	}
	dataStore.CacheData(this.WorkStep.WorkStepName, "file_path", file_path)
}

func (this *FileWriteNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	return schema.BuildParamInputSchemaWithSlice([]string{"file_path", "data?", "bytes?", "append?"})
}

func (this *FileWriteNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *FileWriteNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{"file_path"})
}

func (this *FileWriteNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}
