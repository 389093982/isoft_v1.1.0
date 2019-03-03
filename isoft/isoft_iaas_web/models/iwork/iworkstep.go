package iwork

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type WorkStep struct {
	Id                   int64     `json:"id"`
	WorkId               int64     `json:"work_id"`
	WorkStepId           int64     `json:"work_step_id"`
	WorkSubId            int64     `json:"work_sub_id"` // 子流程 id
	WorkStepName         string    `json:"work_step_name"`
	WorkStepDesc         string    `json:"work_step_desc" orm:"type(text)"`
	WorkStepType         string    `json:"work_step_type"`
	WorkStepInput        string    `json:"work_step_input" orm:"type(text)"`
	WorkStepOutput       string    `json:"work_step_output" orm:"type(text)"`
	WorkStepParamMapping string    `json:"work_step_param_mapping" orm:"type(text)"`
	CreatedBy            string    `json:"created_by"`
	CreatedTime          time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy        string    `json:"last_updated_by"`
	LastUpdatedTime      time.Time `json:"last_updated_time"`
}

// 多字段唯一键
func (u *WorkStep) TableUnique() [][]string {
	return [][]string{
		[]string{"WorkId", "WorkStepName"},
	}
}

func DeleteAllWorkStep(work_id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("work_step").Filter("work_id", work_id).Delete()
	return err
}

func InsertOrUpdateWorkStep(step *WorkStep) (id int64, err error) {
	o := orm.NewOrm()
	if step.Id > 0 {
		id, err = o.Update(step)
	} else {
		id, err = o.Insert(step)
	}
	return
}

func QueryWorkStep(condArr map[string]interface{}) (steps []WorkStep, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("work_step")
	if work_id, ok := condArr["work_id"]; ok {
		qs = qs.Filter("work_id", work_id)
	}
	qs = qs.OrderBy("work_step_id")
	qs.All(&steps)
	return
}

func QueryOneWorkStep(work_id int64, work_step_id int64) (step WorkStep, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("work_step").Filter("work_id", work_id).Filter("work_step_id", work_step_id).One(&step)
	return
}

func QueryAllWorkStepInfo(work_id int64) (steps []WorkStep, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("work_step").Filter("work_id", work_id).OrderBy("work_step_id").All(&steps)
	return
}

func QueryWorkStepInfo(work_id int64, work_step_id int64) (step WorkStep, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("work_step").Filter("work_id", work_id).Filter("work_step_id", work_step_id).One(&step)
	return
}

// mod 只支持 +、- 符号
func BatchChangeWorkStepIdOrder(work_id, work_step_id int64, mod string) error {
	o := orm.NewOrm()
	query := fmt.Sprintf("UPDATE work_step SET work_step_id = work_step_id %s 1 WHERE work_id = ? and work_step_id > ?", mod)
	_, err := o.Raw(query, work_id, work_step_id).Exec()
	return err
}

func DeleteWorkStepByWorkStepId(work_id, work_step_id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("work_step").Filter("work_step_id", work_step_id).Delete()
	if err == nil {
		err = BatchChangeWorkStepIdOrder(work_id, work_step_id, "-")
	}
	return err
}

// 获取前置节点信息
func QueryAllPreStepInfo(work_id int64, work_step_id int64) (steps []WorkStep, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("work_step").Filter("work_id", work_id).
		Filter("work_step_id__lt", work_step_id).OrderBy("work_step_id").All(&steps)
	return
}

func QueryAllWorkStepByWorkName(work_name string) (steps []WorkStep, err error) {
	if work, err := QueryWorkByName(work_name); err == nil {
		steps, err = QueryAllWorkStepInfo(work.Id)
	}
	return
}

func CopyWorkStepInfo(step WorkStep) *WorkStep {
	newStep := &WorkStep{
		WorkStepName:         step.WorkStepName,
		WorkStepType:         step.WorkStepType,
		WorkStepDesc:         step.WorkStepDesc,
		WorkStepInput:        step.WorkStepInput,
		WorkStepOutput:       step.WorkStepOutput,
		WorkStepParamMapping: step.WorkStepParamMapping,
		CreatedBy:            step.CreatedBy,
		CreatedTime:          step.CreatedTime,
		LastUpdatedBy:        step.LastUpdatedBy,
		LastUpdatedTime:      step.LastUpdatedTime,
	}
	return newStep
}
