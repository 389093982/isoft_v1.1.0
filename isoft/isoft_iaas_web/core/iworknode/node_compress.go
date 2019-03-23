package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/compressutil"
	"isoft/isoft_iaas_web/models/iwork"
)

type TarGzUnCompressNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *TarGzUnCompressNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	targz_file_path := tmpDataMap[iworkconst.STRING_PREFIX+"targz_file_path"].(string)
	dest_path := tmpDataMap[iworkconst.STRING_PREFIX+"dest_path"].(string)
	if err := compressutil.DeCompress(targz_file_path, dest_path); err != nil {
		panic(err)
	}
}

func (this *TarGzUnCompressNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "targz_file_path", "targz 文件路径"},
		2: {iworkconst.STRING_PREFIX + "dest_path", "解压后的路径"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *TarGzUnCompressNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *TarGzUnCompressNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *TarGzUnCompressNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *TarGzUnCompressNode) ValidateCustom() (checkResult []string) {
	return
}

type TarGzCompressNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *TarGzCompressNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, this.DataStore)
	dir_file_path := tmpDataMap[iworkconst.STRING_PREFIX+"dir_file_path"].(string)
	dest_file_path := tmpDataMap[iworkconst.STRING_PREFIX+"dest_file_path"].(string)
	if err := compressutil.CompressDir(dir_file_path, dest_file_path); err != nil {
		panic(err)
	}
}

func (this *TarGzCompressNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "dir_file_path", "待压缩的文件夹路径"},
		2: {iworkconst.STRING_PREFIX + "dest_file_path", "压缩后的targz文件路径"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *TarGzCompressNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *TarGzCompressNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *TarGzCompressNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *TarGzCompressNode) ValidateCustom() (checkResult []string) {
	return
}
