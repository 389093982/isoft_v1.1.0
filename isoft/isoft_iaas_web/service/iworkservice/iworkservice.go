package iworkservice

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/core/iworkrun"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

func GetRelativeWorkService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	subWorks := make([]iwork.Work, 0)
	parentWorks, _, err := iwork.QueryParentWorks(work_id, o)
	if err != nil {
		return nil, err
	}
	steps, err := iwork.QueryAllWorkStepInfo(work_id, o)
	if err != nil {
		return nil, err
	}
	for _, step := range steps {
		if step.WorkSubId > 0 {
			subwork, err := iwork.QueryWorkById(step.WorkSubId, o)
			if err != nil {
				return nil, err
			}
			subWorks = append(subWorks, subwork)
		}
	}
	result["parentWorks"] = parentWorks
	result["subworks"] = subWorks
	return
}

func RunWorkService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	work, err := iwork.QueryWorkById(work_id, o)
	if err != nil {
		return err
	}
	steps, err := iwork.QueryAllWorkStepInfo(work_id, o)
	if err != nil {
		return err
	}
	go iworkrun.Run(work, steps, nil)
	return nil
}

func FilterPageLogRecord(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	offset := serviceArgs["offset"].(int)
	current_page := serviceArgs["current_page"].(int)
	ctx := serviceArgs["ctx"].(*context.Context)
	runLogRecords, count, err := iwork.QueryRunLogRecord(work_id, current_page, offset)
	if err != nil {
		return nil, err
	}
	paginator := pagination.SetPaginator(ctx, offset, count)
	if err != nil {
		return nil, err
	}
	result["runLogRecords"] = runLogRecords
	result["paginator"] = pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
	return
}

func FilterPageWorkService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	condArr := serviceArgs["condArr"].(map[string]string)
	offset := serviceArgs["offset"].(int)
	current_page := serviceArgs["current_page"].(int)
	ctx := serviceArgs["ctx"].(*context.Context)
	o := serviceArgs["o"].(orm.Ormer)
	works, count, err := iwork.QueryWork(condArr, current_page, offset, o)
	if err != nil {
		return nil, err
	}
	paginator := pagination.SetPaginator(ctx, offset, count)
	if err != nil {
		return nil, err
	}
	result["works"] = works
	result["paginator"] = pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
	return
}

// 根据 id 信息查找旧的 work 信息
func getOldWorkInfoById(id int64, o orm.Ormer) (oldWorkName string, oldWorkId int64) {
	if id <= 0 {
		return
	}
	if work1, err := iwork.QueryWorkById(id, o); err == nil {
		oldWorkName = work1.WorkName
		oldWorkId = work1.Id
	}
	return
}

func EditWorkService(serviceArgs map[string]interface{}) error {
	work := serviceArgs["work"].(iwork.Work)
	o := serviceArgs["o"].(orm.Ormer)
	oldWorkName, oldWorkId := getOldWorkInfoById(work.Id, o)
	// 插入或者更新 work 信息
	if _, err := iwork.InsertOrUpdateWork(&work, o); err != nil {
		return err
	}
	if oldWorkName == "" {
		// 新增 work 场景,自动添加开始和结束节点
		if err := InsertStartEndWorkStepNode(work.Id, o); err != nil {
			return err
		}
		if _, err := InsertOrUpdateAutoCronMeta(work.WorkName, -1); err != nil {
			return err
		}
	} else {
		// 修改 work 场景
		if err := ChangeReferencesWorkName(oldWorkId, oldWorkName, work.WorkName, o); err != nil {
			return err
		}
		var oldMetaId int64
		if meta, err := iwork.QueryCronMetaByName(oldWorkName); err != nil {
			oldMetaId = -1
		} else {
			oldMetaId = meta.Id
		}
		if _, err := InsertOrUpdateAutoCronMeta(work.WorkName, oldMetaId); err != nil {
			return err
		}
	}
	return nil
}

func InsertOrUpdateAutoCronMeta(task_name string, meta_id int64) (id int64, err error) {
	meta := &iwork.CronMeta{
		TaskName:        task_name,
		TaskType:        "iwork_quartz",
		CronStr:         "0 * * * * ?",
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	if meta_id > 0 {
		meta.Id = meta_id
	}
	id, err = iwork.InsertOrUpdateCronMeta(meta)
	return
}

func DeleteWorkByIdService(serviceArgs map[string]interface{}) error {
	id := serviceArgs["id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	return iwork.DeleteWorkById(id, o)
}

func ChangeReferencesWorkName(work_id int64, oldWorkName, workName string, o orm.Ormer) error {
	if oldWorkName == workName {
		return nil
	}
	parentWorks, _, err := iwork.QueryParentWorks(work_id, o)
	if err != nil {
		return nil
	}
	for _, parentWork := range parentWorks {
		steps, _ := iwork.QueryAllWorkStepInfo(parentWork.Id, o)
		for _, step := range steps {
			if step.WorkStepType != "work_sub" {
				continue
			}
			inputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep: &step})
			for index, item := range inputSchema.ParamInputSchemaItems {
				if item.ParamName == iworkconst.STRING_PREFIX+"work_sub" && strings.Contains(item.ParamValue, oldWorkName) {
					inputSchema.ParamInputSchemaItems[index].ParamValue = strings.Replace(item.ParamValue, oldWorkName, workName, -1)
				}
			}
			step.WorkStepInput = inputSchema.RenderToXml()
			iwork.InsertOrUpdateWorkStep(&step, o)
		}
	}
	return nil
}

func InsertStartEndWorkStepNode(work_id int64, o orm.Ormer) error {
	insertDefaultWorkStepNodeFunc := func(nodeName string, work_step_id int64) error {
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
		if _, err := iwork.InsertOrUpdateWorkStep(step, o); err != nil {
			return err
		}
		return nil
	}
	if err := insertDefaultWorkStepNodeFunc("start", 1); err != nil {
		return err
	}
	if err := insertDefaultWorkStepNodeFunc("end", 2); err != nil {
		return err
	}
	return nil
}
