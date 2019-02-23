package iworkservice

import (
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

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