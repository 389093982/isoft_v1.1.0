package iwork

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Work struct {
	Id              int64     `json:"id"`
	WorkName        string    `json:"work_name"`
	WorkDesc        string    `json:"work_desc" orm:"type(text)"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

type WorkStep struct {
	Id                   int64     `json:"id"`
	WorkId               int64    `json:"work_id"`
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

func GetAllWorkInfo() (works []Work) {
	o := orm.NewOrm()
	o.QueryTable("work").OrderBy("id").All(&works)
	return
}

func QueryWorkById(work_id int64) (work Work, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("work").Filter("id", work_id).One(&work)
	return
}

func QueryWorkByName(work_name string) (work Work, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("work").Filter("work_name", work_name).One(&work)
	return
}

func InsertOrUpdateWork(work *Work) (id int64, err error) {
	o := orm.NewOrm()
	if work.Id > 0 {
		id, err = o.Update(work)
	} else {
		id, err = o.Insert(work)
	}
	return
}


func QueryParentWorks(work_id int64) (works []Work, counts int64, err error) {
	works = make([]Work, 0)
	o := orm.NewOrm()
	params := make([]orm.Params, 0)
	_, err = o.QueryTable("work_step").Filter("work_sub_id", work_id).Distinct().Values(&params, "work_id")
	if err == nil{
		for _, param := range params{
			parent_work_id := param["work_id"].(int64)
			pWork, _ := QueryWorkById(parent_work_id)
			works = append(works, pWork)
		}
	}
	return
}

func QueryWork(condArr map[string]string, page int, offset int) (works []Work, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("work")
	var cond = orm.NewCondition()
	if search, ok := condArr["search"]; ok && strings.TrimSpace(search) != "" {
		subCond := orm.NewCondition()
		subCond = cond.And("work_name__contains", search)
		cond = cond.AndCond(subCond)
	}
	qs = qs.SetCond(cond)
	counts, _ = qs.Count()
	qs = qs.OrderBy("-last_updated_time").Limit(offset, (page-1)*offset)
	qs.All(&works)
	return
}

func DeleteWorkById(id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("work").Filter("id", id).Delete()
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

func GetNextWorkStepId(work_id int64) int64 {
	steps, _ := GetAllWorkStepInfo(work_id)
	for index := 1; index <= len(steps)+1; index++ {
		o := orm.NewOrm()
		if exist := o.QueryTable("work_step").Filter("work_id", work_id).Filter("work_step_id", index).Exist(); !exist {
			return int64(index)
		}
	}
	// 默认返回 1
	return 1
}

func QueryWorkStep(condArr map[string]interface{}, page int, offset int) (steps []WorkStep, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("work_step")
	if work_id, ok := condArr["work_id"]; ok {
		qs = qs.Filter("work_id", work_id)
	}
	qs = qs.OrderBy("work_step_id")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&steps)
	return
}

func GetOneWorkStep(work_id int64, work_step_id int64) (step WorkStep, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("work_step").Filter("work_id", work_id).Filter("work_step_id", work_step_id).One(&step)
	return
}

func GetAllWorkStepInfo(work_id int64) (steps []WorkStep, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("work_step").Filter("work_id", work_id).OrderBy("work_step_id").All(&steps)
	return
}

func LoadWorkStepInfo(work_id int64, work_step_id int64) (step WorkStep, err error) {
	o := orm.NewOrm()
	err = o.QueryTable("work_step").Filter("work_id", work_id).Filter("work_step_id", work_step_id).One(&step)
	return
}

func DeleteWorkStepById(id int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("work_step").Filter("id", id).Delete()
	return err
}

// 获取前置节点信息
func GetAllPreStepInfo(work_id int64, work_step_id int64) (steps []WorkStep, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("work_step").Filter("work_id", work_id).
		Filter("work_step_id__lt", work_step_id).OrderBy("work_step_id").All(&steps)
	return
}

func GetAllWorkStepByWorkName(work_name string) (steps []WorkStep, err error) {
	if work, err := QueryWorkByName(work_name); err == nil {
		steps, err = GetAllWorkStepInfo(work.Id)
	}
	return
}
