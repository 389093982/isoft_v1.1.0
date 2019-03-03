package iworkservice

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

// 根据 id 信息查找旧的 work 信息
func getOldWorkInfoById(id int64) (oldWorkName string, oldWorkId int64) {
	if id <= 0{
		return
	}
	if work1, err := iwork.QueryWorkById(id); err == nil{
		oldWorkName = work1.WorkName
		oldWorkId = work1.Id
	}
	return
}

func EditWorkService(work iwork.Work) error {
	oldWorkName,oldWorkId := getOldWorkInfoById(work.Id)
	// 插入或者更新 work 信息
	if _, err := iwork.InsertOrUpdateWork(&work); err != nil{
		return err
	}
	if oldWorkName == ""{
		// 新增 work 场景,自动添加开始和结束节点
		if err := InsertStartEndWorkStepNode(work.Id); err != nil{
			return err
		}
		if _, err := InsertOrUpdateAutoCronMeta(work.WorkName, -1); err != nil{
			return err
		}
	}else{
		// 修改 work 场景
		if err := ChangeReferencesWorkName(oldWorkId, oldWorkName, work.WorkName); err != nil{
			return err
		}
		var oldMetaId int64
		if meta, err := iwork.QueryCronMetaByName(oldWorkName); err != nil{
			oldMetaId = -1
		}else{
			oldMetaId = meta.Id
		}
		if _, err := InsertOrUpdateAutoCronMeta(work.WorkName, oldMetaId); err != nil{
			return err
		}
	}
	return nil
}

func InsertOrUpdateAutoCronMeta(task_name string, meta_id int64) (id int64, err error) {
	meta := &iwork.CronMeta{
		TaskName:task_name,
		TaskType:"iwork_quartz",
		CronStr:"0 * * * * ?",
		CreatedBy:"SYSTEM",
		CreatedTime:time.Now(),
		LastUpdatedBy:"SYSTEM",
		LastUpdatedTime:time.Now(),
	}
	if meta_id > 0{
		meta.Id = meta_id
	}
	id , err = iwork.InsertOrUpdateCronMeta(meta)
	return
}

func DeleteWorkByIdService(id int64) error {
	return iwork.DeleteWorkById(id)
}

func ChangeReferencesWorkName(work_id int64, oldWorkName,workName string) error {
	if oldWorkName == workName{
		return nil
	}
	parentWorks, _,err := iwork.QueryParentWorks(work_id)
	if err != nil {
		return nil
	}
	for _, parentWork := range parentWorks{
		steps, _ := iwork.QueryAllWorkStepInfo(parentWork.Id)
		for _, step := range steps{
			if step.WorkStepType != "work_sub"{
				continue
			}
			inputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep:&step})
			for index, item := range inputSchema.ParamInputSchemaItems{
				if item.ParamName == iworkconst.STRING_PREFIX + "work_sub" && strings.Contains(item.ParamValue, oldWorkName){
					inputSchema.ParamInputSchemaItems[index].ParamValue = strings.Replace(item.ParamValue, oldWorkName, workName, -1)
				}
			}
			step.WorkStepInput = inputSchema.RenderToXml()
			iwork.InsertOrUpdateWorkStep(&step)
		}
	}
	return nil
}

func InsertStartEndWorkStepNode(work_id int64) error {
	insertDefaultWorkStepNodeFunc := func(nodeName string, work_step_id int64) error{
		step := &iwork.WorkStep{
			WorkId:          work_id,
			WorkStepId:      work_step_id,
			WorkStepName:    nodeName,
			WorkStepDesc:    fmt.Sprintf("%s节点", nodeName),
			WorkStepType:    fmt.Sprintf("work_%s", nodeName),
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		if _, err := iwork.InsertOrUpdateWorkStep(step); err != nil{
			return err
		}
		return nil
	}
	if err := insertDefaultWorkStepNodeFunc("start", 1); err != nil{
		return err
	}
	if err := insertDefaultWorkStepNodeFunc("end", 2); err != nil{
		return err
	}
	return nil
}