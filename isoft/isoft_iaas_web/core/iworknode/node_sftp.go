package iworknode

import (
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/param"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworkutil/sftputil"
	"isoft/isoft_iaas_web/models/iwork"
)

type SftpUploadNode struct {
	BaseNode
	WorkStep *iwork.WorkStep
}

func (this *SftpUploadNode) Execute(trackingId string) {
	// 数据中心
	dataStore := datastore.GetDataStore(trackingId)
	// 节点中间数据
	tmpDataMap := this.FillParamInputSchemaDataToTmp(this.WorkStep, dataStore)
	sftpResource := param.GetStaticParamValue(iworkconst.STRING_PREFIX+"db_conn", this.WorkStep).(iwork.Resource)
	local_file_path := tmpDataMap[iworkconst.STRING_PREFIX+"local_file_path"].(string)
	remote_file_path := tmpDataMap[iworkconst.STRING_PREFIX+"remote_file_path"].(string)
	err := sftputil.SFTPFileCopy(sftpResource.ResourceUsername, sftpResource.ResourcePassword, sftpResource.ResourceDsn, 22, local_file_path, remote_file_path)
	if err == nil {
		dataStore.CacheData(this.WorkStep.WorkStepName, iworkconst.STRING_PREFIX+"remote_file_path", remote_file_path)
	} else {
		panic(err)
	}
}

func (this *SftpUploadNode) GetDefaultParamInputSchema() *schema.ParamInputSchema {
	paramMap := map[int][]string{
		1: {iworkconst.STRING_PREFIX + "sftp_conn", "sftp连接信息,需要使用 $RESOURCE 全局参数"},
		2: {iworkconst.STRING_PREFIX + "local_file_path", "本地文件路径"},
		3: {iworkconst.STRING_PREFIX + "remote_file_path", "远程文件路径"},
	}
	return schema.BuildParamInputSchemaWithDefaultMap(paramMap)
}

func (this *SftpUploadNode) GetRuntimeParamInputSchema() *schema.ParamInputSchema {
	return &schema.ParamInputSchema{}
}

func (this *SftpUploadNode) GetDefaultParamOutputSchema() *schema.ParamOutputSchema {
	return schema.BuildParamOutputSchemaWithSlice([]string{iworkconst.STRING_PREFIX + "remote_file_path"})
}

func (this *SftpUploadNode) GetRuntimeParamOutputSchema() *schema.ParamOutputSchema {
	return &schema.ParamOutputSchema{}
}

func (this *SftpUploadNode) ValidateCustom() (checkResult []string) {
	return
}
