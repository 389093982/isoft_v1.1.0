package iworkservice

import (
	"encoding/json"
	"errors"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

func EditWorkStepBaseInfoService(step *iwork.WorkStep) error {
	oldStep, err := iwork.QueryOneWorkStep(step.WorkId, step.WorkStepId)
	if err != nil {
		return err
	}
	// 变更类型需要置空 input 和 output 参数
	if step.WorkStepType != oldStep.WorkStepType {
		step.WorkStepInput = ""
		step.WorkStepOutput = ""
	}
	if _, err := iwork.InsertOrUpdateWorkStep(step); err == nil {
		// 级联更改相关联的步骤名称
		if err := ChangeReferencesWorkStepName(step.WorkId, oldStep.WorkStepName, step.WorkStepName); err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

func DeleteWorkStepByWorkStepIdService(work_id, work_step_id int64) error {
	return iwork.DeleteWorkStepByWorkStepId(work_id, work_step_id)
}

func ChangeReferencesWorkStepName(work_id int64, oldWorkStepName, workStepName string) error {
	if oldWorkStepName == workStepName {
		return nil
	}
	steps, err := iwork.QueryAllWorkStepInfo(work_id)
	if err != nil {
		return err
	}
	for _, step := range steps {
		step.WorkStepInput = strings.Replace(step.WorkStepInput, "$"+oldWorkStepName, "$"+workStepName, -1)
		_, err := iwork.InsertOrUpdateWorkStep(&step)
		if err != nil {
			return err
		}
	}
	return nil
}

func RefactorWorkStepInfoService(serviceArgs map[string]interface{}) error {
	// 获取参数
	work_id := serviceArgs["work_id"].(int64)
	refactor_worksub_name := serviceArgs["refactor_worksub_name"].(string)
	refactor_work_step_ids := serviceArgs["refactor_work_step_ids"].(string)
	var refactor_work_step_id_arr []int
	json.Unmarshal([]byte(refactor_work_step_ids), &refactor_work_step_id_arr)
	// 校验 refactor_work_step_id_arr 是否连续
	if refactor_work_step_id_arr[len(refactor_work_step_id_arr)-1]-refactor_work_step_id_arr[0] != len(refactor_work_step_id_arr)-1 {
		return errors.New("refactor workStepId 必须是连续的!")
	}
	// 创建子流程
	subWork := &iwork.Work{
		WorkName:        refactor_worksub_name,
		WorkDesc:        "refactor worksub",
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	iwork.InsertOrUpdateWork(subWork)
	// 循环移动子步骤
	for index, work_step_id := range refactor_work_step_id_arr {
		step, err := iwork.QueryWorkStepInfo(work_id, int64(work_step_id))
		if err == nil {
			if step.WorkStepType == "work_start" || step.WorkStepType == "work_end" {
				return errors.New("start 和 end 节点不能重构！")
			}
			InsertStartEndWorkStepNode(subWork.Id)
			newStep := iwork.CopyWorkStepInfo(step)
			newStep.WorkId = subWork.Id
			newStep.WorkStepId = int64(index + 2)
			iwork.InsertOrUpdateWorkStep(newStep)
			// 当前流程循环删除该节点
			DeleteWorkStepByWorkStepIdService(work_id, int64(work_step_id))
		}
	}
	return nil
}
