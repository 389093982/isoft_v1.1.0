package iwork

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Work struct {
	Id              int64     	`json:"id"`
	WorkName		string    	`json:"work_step_id"`
	CreatedBy       string    	`json:"created_by"`
	CreatedTime     time.Time 	`json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    	`json:"last_updated_by"`
	LastUpdatedTime time.Time 	`json:"last_updated_time"`
}

type WorkStep struct {
	Id              int64     	`json:"id"`
	WorkId      	string    	`json:"work_id"`
	WorkStepId      int8    	`json:"work_step_id"`
	WorkStepInput   string    	`json:"work_step_input"`
	WorkStepOutput	string    	`json:"work_step_output"`
	CreatedBy       string    	`json:"created_by"`
	CreatedTime     time.Time 	`json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    	`json:"last_updated_by"`
	LastUpdatedTime time.Time 	`json:"last_updated_time"`
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

func QueryWorkStep(condArr map[string]string, page int, offset int) (steps []WorkStep, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("work_step").OrderBy("work_step_id")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&steps)
	return
}

func DeleteWorkStepById(id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("work_step").Filter("id", id).Delete()
	return err
}