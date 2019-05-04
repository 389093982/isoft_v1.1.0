package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/core/iworkutil/compressutil"
	"isoft/isoft_iaas_web/models/iwork"
)

type TarGzUnCompressNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *TarGzUnCompressNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	targz_file_path := tmpDataMap[iworkconst.STRING_PREFIX+"targz_file_path"].(string)
	dest_path := tmpDataMap[iworkconst.STRING_PREFIX+"dest_dir_path"].(string)
	if err := compressutil.DeCompress(targz_file_path, dest_path); err == nil {
		this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "dest_dir_path": dest_path})
	} else {
		panic(err)
	}
}

func (this *TarGzUnCompressNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "targz_file_path", "targz 文件路径"},
		2: {iworkconst.STRING_PREFIX + "dest_dir_path", "解压后的文件夹路径"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *TarGzUnCompressNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "dest_dir_path"})
}

type TarGzCompressNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *TarGzCompressNode) Execute(trackingId string) {
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep)
	dir_file_path := tmpDataMap[iworkconst.STRING_PREFIX+"src_dir_path"].(string)
	dest_file_path := tmpDataMap[iworkconst.STRING_PREFIX+"dest_file_path"].(string)
	if err := compressutil.CompressDir(dir_file_path, dest_file_path); err == nil {
		this.DataStore.CacheDatas(this.WorkStep.WorkStepName, map[string]interface{}{iworkconst.STRING_PREFIX + "dest_file_path": dest_file_path})
	} else {
		panic(err)
	}
}

func (this *TarGzCompressNode) GetDefaultParamInputSchema() *iworkmodels.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "src_dir_path", "待压缩的文件夹路径"},
		2: {iworkconst.STRING_PREFIX + "dest_file_path", "压缩后的targz文件路径"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *TarGzCompressNode) GetDefaultParamOutputSchema() *iworkmodels.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "dest_file_path"})
}
