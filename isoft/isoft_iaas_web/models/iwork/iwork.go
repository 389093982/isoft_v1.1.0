package iwork

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Work struct {
	Id              int64     `json:"id"`
	WorkName        string    `json:"work_name" orm:"unique"`
	WorkDesc        string    `json:"work_desc" orm:"type(text)"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time" orm:"auto_now_add;type(datetime)"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func QueryAllWorkInfo(o orm.Ormer) (works []Work) {
	o.QueryTable("work").OrderBy("id").All(&works)
	return
}

func QueryWorkById(work_id int64, o orm.Ormer) (work Work, err error) {
	err = o.QueryTable("work").Filter("id", work_id).One(&work)
	return
}

func QueryWorkByName(work_name string, o orm.Ormer) (work Work, err error) {
	err = o.QueryTable("work").Filter("work_name", work_name).One(&work)
	return
}

func InsertOrUpdateWork(work *Work, o orm.Ormer) (id int64, err error) {
	if work.Id > 0 {
		id, err = o.Update(work)
	} else {
		id, err = o.Insert(work)
	}
	return
}

func QueryParentWorks(work_id int64, o orm.Ormer) (workSteps []WorkStep, works []Work, counts int64, err error) {
	qs := o.QueryTable("work_step").Filter("work_sub_id", work_id)
	if _, err = qs.All(&workSteps); err != nil {
		return
	}
	params := make([]orm.Params, 0)
	if _, err = qs.Distinct().Values(&params, "work_id"); err != nil {
		return
	}
	works = make([]Work, 0)
	for _, param := range params {
		p_work_id := param["WorkId"].(int64)
		pWork, _err := QueryWorkById(p_work_id, o)
		if _err != nil {
			err = _err
			return
		}
		works = append(works, pWork)
	}
	return
}

func QueryWork(condArr map[string]string, page int, offset int, o orm.Ormer) (works []Work, counts int64, err error) {
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

func DeleteWorkById(id int64, o orm.Ormer) error {
	if err := DeleteAllWorkStep(id, o); err != nil {
		return err
	}
	_, err := o.QueryTable("work").Filter("id", id).Delete()
	return err
}
