package iworkservice

import (
	"encoding/json"
	"errors"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

func DeleteWorkStepByWorkStepIdService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	if err := iwork.DeleteWorkStepByWorkStepId(work_id, work_step_id); err != nil {
		return err
	}
	return nil
}

func FilterWorkStepService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	condArr := make(map[string]interface{})
	condArr["work_id"] = serviceArgs["work_id"].(int64)
	worksteps, err := iwork.QueryWorkStep(condArr)
	if err != nil {
		return nil, err
	}
	result["worksteps"] = worksteps
	return
}

func AddWorkStepService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	work_step_type := serviceArgs["default_work_step_type"].(string)
	// 将 work_step_id 之后的所有节点后移一位
	err := iwork.BatchChangeWorkStepIdOrder(work_id, work_step_id, "+")
	if err != nil {
		return err
	}
	step := &iwork.WorkStep{
		WorkId:          work_id,
		WorkStepName:    "random_" + stringutil.RandomUUID(),
		WorkStepType:    work_step_type,
		WorkStepId:      work_step_id + 1,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	if _, err := iwork.InsertOrUpdateWorkStep(step); err != nil {
		return err
	}
	return nil
}

func ChangeWorkStepOrderService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	_type := serviceArgs["_type"].(string)
	// 获取当前步骤
	step, err := iwork.QueryOneWorkStep(work_id, work_step_id)
	if err != nil {
		return err
	}
	if _type == "up" {
		prevStep, err := iwork.QueryOneWorkStep(work_id, work_step_id-1)
		if err != nil {
			return err
		}
		prevStep.WorkStepId = prevStep.WorkStepId + 1
		step.WorkStepId = step.WorkStepId - 1
		if _, err := iwork.InsertOrUpdateWorkStep(&prevStep); err != nil {
			return err
		}
		if _, err := iwork.InsertOrUpdateWorkStep(&step); err != nil {
			return err
		}
	} else {
		nextStep, err := iwork.QueryOneWorkStep(work_id, work_step_id+1)
		if err != nil {
			return err
		}
		nextStep.WorkStepId = nextStep.WorkStepId + 1
		step.WorkStepId = step.WorkStepId + 1
		if _, err := iwork.InsertOrUpdateWorkStep(&nextStep); err != nil {
			return err
		}
		if _, err := iwork.InsertOrUpdateWorkStep(&step); err != nil {
			return err
		}
	}
	return nil
}

func EditWorkStepBaseInfoService(serviceArgs map[string]interface{}) error {
	step := serviceArgs["step"].(*iwork.WorkStep)
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
	if _, err := iwork.InsertOrUpdateWork(subWork); err != nil {
		return err
	}
	// 为子流程添加开始和结束节点
	if err := InsertStartEndWorkStepNode(subWork.Id); err != nil {
		return err
	}
	// 循环移动子步骤
	for index, work_step_id := range refactor_work_step_id_arr {
		step, err := iwork.QueryWorkStepInfo(work_id, int64(work_step_id))
		if err != nil {
			return err
		}
		if step.WorkStepType == "work_start" || step.WorkStepType == "work_end" {
			return errors.New("start 和 end 节点不能重构！")
		}
		newStep := iwork.CopyWorkStepInfo(step)
		newStep.WorkId = subWork.Id
		newStep.WorkStepId = int64(index + 2)
		if _, err := iwork.InsertOrUpdateWorkStep(newStep); err != nil {
			return err
		}
		// 当前流程循环删除该节点
		if err := DeleteWorkStepByWorkStepIdService(work_id, int64(work_step_id)); err != nil {
			return err
		}
	}
	return nil
}
