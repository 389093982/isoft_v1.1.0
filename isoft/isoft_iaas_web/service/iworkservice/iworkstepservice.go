package iworkservice

import (
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

func EditWorkStepBaseInfoService(step *iwork.WorkStep) error {
	oldStep, err := iwork.QueryOneWorkStep(step.WorkId, step.WorkStepId)
	if err != nil{
		return err
	}
	// 变更类型需要置空 input 和 output 参数
	if step.WorkStepType != oldStep.WorkStepType {
		step.WorkStepInput = ""
		step.WorkStepOutput = ""
	}
	if _, err := iwork.InsertOrUpdateWorkStep(step); err == nil {
		// 级联更改相关联的步骤名称
		if err := ChangeReferencesWorkStepName(step.WorkId, oldStep.WorkStepName, step.WorkStepName); err != nil{
			return err
		}
	}else{
		return err
	}
	return nil
}

func DeleteWorkStepByWorkStepIdService(work_id, work_step_id int64) error {
	return iwork.DeleteWorkStepByWorkStepId(work_id, work_step_id)
}

func ChangeReferencesWorkStepName(work_id int64, oldWorkStepName, workStepName string) error {
	if oldWorkStepName == workStepName{
		return nil
	}
	steps, err := iwork.QueryAllWorkStepInfo(work_id)
	if err != nil{
		return err
	}
	for _, step := range steps{
		step.WorkStepInput = strings.Replace(step.WorkStepInput, "$" + oldWorkStepName, "$" + workStepName, -1)
		_, err := iwork.InsertOrUpdateWorkStep(&step)
		if err != nil{
			return err
		}
	}
	return nil
}