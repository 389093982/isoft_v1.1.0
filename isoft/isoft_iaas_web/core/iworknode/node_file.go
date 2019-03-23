package iworknode

import (
	"io/ioutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/fileutil"
	"isoft/isoft_iaas_web/models/iwork"
	"os"
	"strings"
)

type FileReadNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *FileReadNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	file_path := tmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"file_path", file_path)
	if bytes, err := ioutil.ReadFile(file_path); err == nil {
		this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"data", string(bytes))
	} else {
		panic(err)
	}
}

func (this *FileReadNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "file_path", "读取文件的绝对路径"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *FileReadNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *FileReadNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "file_path", "data"})
}

func (this *FileReadNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *FileReadNode) ValidateCustom() (checkResult []string) {
	return
}

type FileWriteNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func checkAppend(tmpDataMap map[string]interface{}) bool {
	if append, ok := tmpDataMap[iworkconst.BOOL_PREFIX+"append?"].(string); ok && strings.TrimSpace(append) != "" {
		return true
	}
	return false
}

func (this *FileWriteNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	file_path := tmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	var strdata string
	// 写字符串
	if data, ok := tmpDataMap[iworkconst.STRING_PREFIX+"data?"].(string); ok {
		strdata = data
	}
	// 写字节数组
	if bytes, ok := tmpDataMap[iworkconst.BYTE_ARRAY_PREFIX+"data?"].([]byte); ok {
		strdata = string(bytes)
	}
	// 判断是否需要添加行分隔符
	if linesep, ok := tmpDataMap[iworkconst.BOOL_PREFIX+"linesep?"].(string); ok && strings.TrimSpace(linesep) != "" {
		strdata += "\n"
	}
	if err := fileutil.WriteFile(file_path, []byte(strdata), checkAppend(tmpDataMap)); err != nil {
		panic(err)
	}
	this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"file_path", file_path)
}

func (this *FileWriteNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "file_path", "写入文件的绝对路径,文件不存在时会自动创建"},
		2: {iworkconst.STRING_PREFIX + "data?", "可选参数,写入文件的字符数据"},
		3: {iworkconst.BYTE_ARRAY_PREFIX + "data?", "可选参数,写入文件的二进制字节数据"},
		4: {iworkconst.BOOL_PREFIX + "append?", "可选参数,文件追加模式,值为空表示覆盖,有值表示追加"},
		5: {iworkconst.BOOL_PREFIX + "linesep?", "可选参数,行分隔符,默认没有分割符,有值表示使用换行符进行分割"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *FileWriteNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *FileWriteNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "file_path"})
}

func (this *FileWriteNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *FileWriteNode) ValidateCustom() (checkResult []string) {
	return
}

type FileRenameNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *FileRenameNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	file_path := tmpDataMap[iworkconst.STRING_PREFIX+"file_path"].(string)
	new_file_path := tmpDataMap[iworkconst.STRING_PREFIX+"new_file_path"].(string)
	if err := os.Rename(file_path, new_file_path); err == nil {
		this.DataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"file_path", new_file_path)
	} else {
		panic(err)
	}
}

func (this *FileRenameNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "file_path", "需要进行移动重命名的文件路径"},
		2: {iworkconst.STRING_PREFIX + "new_file_path", "移动重命名后的文件路径"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *FileRenameNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *FileRenameNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *FileRenameNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *FileRenameNode) ValidateCustom() (checkResult []string) {
	return
}
