package iworkservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/schema"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

func LoadResourceInfo() *schema.ParamOutputSchema {
	pos := &schema.ParamOutputSchema{
		ParamOutputSchemaItems: []schema.ParamOutputSchemaItem{},
	}
	resources := iwork.QueryAllResource()
	for _, resource := range resources {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, schema.ParamOutputSchemaItem{
			ParamName: resource.ResourceName,
		})
	}
	return pos
}

func LoadWorkInfo() *schema.ParamOutputSchema {
	pos := &schema.ParamOutputSchema{
		ParamOutputSchemaItems: []schema.ParamOutputSchemaItem{},
	}
	works := iwork.QueryAllWorkInfo(orm.NewOrm())
	for _, work := range works {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, schema.ParamOutputSchemaItem{
			ParamName: work.WorkName,
		})
	}
	return pos
}

func LoadEntityInfo() *schema.ParamOutputSchema {
	pos := &schema.ParamOutputSchema{
		ParamOutputSchemaItems: []schema.ParamOutputSchemaItem{},
	}
	entities := iwork.QueryAllEntityInfo()
	for _, entity := range entities {
		pos.ParamOutputSchemaItems = append(pos.ParamOutputSchemaItems, schema.ParamOutputSchemaItem{
			ParamName: entity.EntityName,
		})
	}
	return pos
}

func LoadPreNodeOutputService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	preParamOutputSchemaTreeNodeArr := make([]*schema.TreeNode, 0)
	// 加载 resource 参数
	pos := LoadResourceInfo()
	preParamOutputSchemaTreeNodeArr = append(preParamOutputSchemaTreeNodeArr, pos.RenderToTreeNodes("$RESOURCE"))
	// 加载 work 参数
	pos = LoadWorkInfo()
	preParamOutputSchemaTreeNodeArr = append(preParamOutputSchemaTreeNodeArr, pos.RenderToTreeNodes("$WORK"))
	// 加载 entity 参数
	pos = LoadEntityInfo()
	preParamOutputSchemaTreeNodeArr = append(preParamOutputSchemaTreeNodeArr, pos.RenderToTreeNodes("$Entity"))
	// 加载前置步骤输出
	if steps, err := iwork.QueryAllPreStepInfo(work_id, work_step_id, o); err == nil {
		for _, step := range steps {
			pos := schema.GetCacheParamOutputSchema(&step)
			preParamOutputSchemaTreeNodeArr = append(preParamOutputSchemaTreeNodeArr, pos.RenderToTreeNodes("$"+step.WorkStepName))
		}
	}
	// 返回结果
	result["preParamOutputSchemaTreeNodeArr"] = preParamOutputSchemaTreeNodeArr
	return
}

func GetAllWorkStepInfoService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	steps, err := iwork.QueryAllWorkStepInfo(work_id, o)
	if err != nil {
		return nil, err
	}
	result["steps"] = steps
	return
}

func LoadWorkStepInfoService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	// 读取 work_step 信息
	step, err := iwork.QueryWorkStepInfo(work_id, work_step_id, o)
	if err != nil {
		return nil, err
	}
	var paramMappingsArr []string
	json.Unmarshal([]byte(step.WorkStepParamMapping), &paramMappingsArr)
	result["step"] = step
	result["paramInputSchema"] = schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep: &step})
	result["paramOutputSchema"] = schema.GetCacheParamOutputSchema(&step)
	result["paramOutputSchemaTreeNode"] = schema.GetCacheParamOutputSchema(&step).RenderToTreeNodes("output")
	result["paramMappings"] = paramMappingsArr
	return
}

func DeleteWorkStepByWorkStepIdService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	if step, err := iwork.QueryWorkStepInfo(work_id, work_step_id, o); err == nil {
		if step.WorkStepType == "work_start" || step.WorkStepType == "work_end" {
			return errors.New("start 节点和 end 节点不能被删除!")
		}
	}
	if err := iwork.DeleteWorkStepByWorkStepId(work_id, work_step_id, o); err != nil {
		return err
	}
	return nil
}

func FilterWorkStepService(serviceArgs map[string]interface{}) (result map[string]interface{}, err error) {
	result = make(map[string]interface{}, 0)
	condArr := make(map[string]interface{})
	condArr["work_id"] = serviceArgs["work_id"].(int64)
	o := serviceArgs["o"].(orm.Ormer)
	worksteps, err := iwork.QueryWorkStep(condArr, o)
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
	o := serviceArgs["o"].(orm.Ormer)
	// end 节点之后不能添加节点
	if step, err := iwork.QueryWorkStepInfo(work_id, work_step_id, o); err == nil {
		if step.WorkStepType == "work_end" {
			return errors.New("不能再 end 节点后面添加节点!")
		}
	}
	// 将 work_step_id 之后的所有节点后移一位
	err := iwork.BatchChangeWorkStepIdOrder(work_id, work_step_id, "+", o)
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
	if _, err := iwork.InsertOrUpdateWorkStep(step, o); err != nil {
		return err
	}
	return nil
}

// 更改邻近两个节点的顺序
func changeNearWorkStepOrder(work_id, work_step_id int64, o orm.Ormer, nearStepLength int64) error {
	// 获取当前步骤
	step, err := iwork.QueryOneWorkStep(work_id, work_step_id, o)
	if err != nil {
		return err
	}
	// 获取邻近步骤
	nearStep, err := iwork.QueryOneWorkStep(work_id, work_step_id-nearStepLength, o)
	if err != nil {
		return err
	}
	if nearStep.WorkStepType == "work_start" || nearStep.WorkStepType == "work_end" {
		return errors.New("start 节点和 end 节点不能移动位置!")
	}
	nearStep.WorkStepId = nearStep.WorkStepId + nearStepLength
	step.WorkStepId = step.WorkStepId - nearStepLength
	// 更新邻近步骤
	if _, err := iwork.InsertOrUpdateWorkStep(&nearStep, o); err != nil {
		return err
	}
	// 更新当前步骤
	if _, err := iwork.InsertOrUpdateWorkStep(&step, o); err != nil {
		return err
	}
	return nil
}

func ChangeWorkStepOrderService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	_type := serviceArgs["_type"].(string)
	o := serviceArgs["o"].(orm.Ormer)
	if _type == "up" {
		return changeNearWorkStepOrder(work_id, work_step_id, o, 1)
	} else {
		return changeNearWorkStepOrder(work_id, work_step_id, o, -1)
	}
	return nil
}

func EditWorkStepBaseInfoService(serviceArgs map[string]interface{}) error {
	step := serviceArgs["step"].(*iwork.WorkStep)
	o := serviceArgs["o"].(orm.Ormer)
	oldStep, err := iwork.QueryOneWorkStep(step.WorkId, step.WorkStepId, o)
	if err != nil {
		return err
	}
	// 变更类型需要置空 input 和 output 参数
	if step.WorkStepType != oldStep.WorkStepType {
		step.WorkStepInput = ""
		step.WorkStepOutput = ""
	}
	if _, err := iwork.InsertOrUpdateWorkStep(step, o); err != nil {
		return err
	}
	// 级联更改相关联的步骤名称
	if err := ChangeReferencesWorkStepName(step.WorkId, oldStep.WorkStepName, step.WorkStepName, o); err != nil {
		return err
	}
	return nil
}

func ChangeReferencesWorkStepName(work_id int64, oldWorkStepName, workStepName string, o orm.Ormer) error {
	if oldWorkStepName == workStepName {
		return nil
	}
	steps, err := iwork.QueryAllWorkStepInfo(work_id, o)
	if err != nil {
		return err
	}
	for _, step := range steps {
		step.WorkStepInput = strings.Replace(step.WorkStepInput, "$"+oldWorkStepName, "$"+workStepName, -1)
		_, err := iwork.InsertOrUpdateWorkStep(&step, o)
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
	o := serviceArgs["o"].(orm.Ormer)
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
	if _, err := iwork.InsertOrUpdateWork(subWork, o); err != nil {
		return err
	}
	// 为子流程添加开始和结束节点
	if err := InsertStartEndWorkStepNode(subWork.Id, o); err != nil {
		return err
	}
	// 循环移动子步骤
	for index, work_step_id := range refactor_work_step_id_arr {
		step, err := iwork.QueryWorkStepInfo(work_id, int64(work_step_id), o)
		if err != nil {
			return err
		}
		if step.WorkStepType == "work_start" || step.WorkStepType == "work_end" {
			return errors.New("start 和 end 节点不能重构！")
		}
		newStep := iwork.CopyWorkStepInfo(step)
		newStep.WorkId = subWork.Id
		newStep.WorkStepId = int64(index + 2)
		if _, err := iwork.InsertOrUpdateWorkStep(newStep, o); err != nil {
			return err
		}
		// 当前流程循环删除该节点
		_serviceArgs := map[string]interface{}{"work_id": work_id, "work_step_id": int64(work_step_id)}
		if err := DeleteWorkStepByWorkStepIdService(_serviceArgs); err != nil {
			return err
		}
	}
	return nil
}

func EditWorkStepParamInfoService(serviceArgs map[string]interface{}) error {
	work_id := serviceArgs["work_id"].(int64)
	work_step_id := serviceArgs["work_step_id"].(int64)
	paramInputSchemaStr := serviceArgs["paramInputSchemaStr"].(string)
	paramMappingsStr := serviceArgs["paramMappingsStr"].(string)
	o := serviceArgs["o"].(orm.Ormer)
	var paramInputSchema schema.ParamInputSchema
	json.Unmarshal([]byte(paramInputSchemaStr), &paramInputSchema)
	step, err := iwork.QueryOneWorkStep(work_id, work_step_id, o)
	if err != nil {
		return err
	}
	step.WorkStepInput = paramInputSchema.RenderToXml()
	step.WorkStepParamMapping = paramMappingsStr
	step.CreatedBy = "SYSTEM"
	step.CreatedTime = time.Now()
	step.LastUpdatedBy = "SYSTEM"
	step.LastUpdatedTime = time.Now()
	_, err = iwork.InsertOrUpdateWorkStep(&step, o)
	if err != nil {
		return err
	}
	// 保存完静态参数后自动构建获动态参数并保存
	BuildDynamic(work_id, work_step_id, o)
	return nil
}

// 构建动态值
func BuildDynamic(work_id int64, work_step_id int64, o orm.Ormer) {
	// 自动创建子流程
	BuildAutoCreateSubWork(work_id, work_step_id, o)
	// 构建动态输入值
	BuildDynamicInput(work_id, work_step_id, o)
	// 构建动态输出值
	BuildDynamicOutput(work_id, work_step_id, o)
}

// 构建动态输入值
func BuildDynamicInput(work_id int64, work_step_id int64, o orm.Ormer) {
	// 读取 work_step 信息
	step, err := iwork.QueryWorkStepInfo(work_id, work_step_id, o)
	if err != nil {
		panic(err)
	}
	// 获取默认数据
	defaultParamInputSchema := schema.GetDefaultParamInputSchema(&iworknode.WorkStepFactory{WorkStep: &step})
	// 获取动态数据
	runtimeParamInputSchema := schema.GetRuntimeParamInputSchema(&iworknode.WorkStepFactory{WorkStep: &step})
	// 合并默认数据和动态数据作为新数据
	newInputSchemaItems := append(defaultParamInputSchema.ParamInputSchemaItems, runtimeParamInputSchema.ParamInputSchemaItems...)
	// 获取历史数据
	historyParamInputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep: &step})
	for index, newInputSchemaItem := range newInputSchemaItems {
		// 存在则不添加且沿用旧值
		if exist, paramValue := CheckAndGetParamValueByInputSchemaParamName(historyParamInputSchema.ParamInputSchemaItems, newInputSchemaItem.ParamName); exist {
			newInputSchemaItems[index].ParamValue = paramValue
		}
	}
	paramInputSchema := &schema.ParamInputSchema{ParamInputSchemaItems: newInputSchemaItems}
	step.WorkStepInput = paramInputSchema.RenderToXml()
	if _, err = iwork.InsertOrUpdateWorkStep(&step, o); err != nil {
		panic(err)
	}
}

// 构建动态输出值
func BuildDynamicOutput(work_id int64, work_step_id int64, o orm.Ormer) {
	// 读取 work_step 信息
	step, err := iwork.QueryWorkStepInfo(work_id, work_step_id, o)
	if err != nil {
		panic(err)
	}
	runtimeParamOutputSchema := schema.GetRuntimeParamOutputSchema(&iworknode.WorkStepFactory{WorkStep: &step})
	defaultParamOutputSchema := schema.GetDefaultParamOutputSchema(&iworknode.WorkStepFactory{WorkStep: &step})
	defaultParamOutputSchema.ParamOutputSchemaItems = append(defaultParamOutputSchema.ParamOutputSchemaItems, runtimeParamOutputSchema.ParamOutputSchemaItems...)
	// 构建输出参数,使用全新值
	step.WorkStepOutput = defaultParamOutputSchema.RenderToXml()
	if _, err = iwork.InsertOrUpdateWorkStep(&step, o); err != nil {
		panic(err)
	}
}

func checkAndCreateSubWork(work_name string, o orm.Ormer) {
	if _, err := iwork.QueryWorkByName(work_name, orm.NewOrm()); err != nil {
		// 不存在 work 则直接创建
		work := &iwork.Work{
			WorkName:        work_name,
			WorkDesc:        fmt.Sprintf("自动创建子流程:%s", work_name),
			CreatedBy:       "SYSTEM",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "SYSTEM",
			LastUpdatedTime: time.Now(),
		}
		if _, err := iwork.InsertOrUpdateWork(work, o); err == nil {
			// 写入 DB 并自动创建开始和结束节点
			InsertStartEndWorkStepNode(work.Id, o)
		}
	}
}

func BuildAutoCreateSubWork(work_id int64, work_step_id int64, o orm.Ormer) {
	// 读取 work_step 信息
	step, err := iwork.QueryWorkStepInfo(work_id, work_step_id, o)
	if err != nil {
		panic(err)
	}
	if step.WorkStepType != "work_sub" {
		return
	}
	paramInputSchema := schema.GetCacheParamInputSchema(&step, &iworknode.WorkStepFactory{WorkStep: &step})
	for index, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == iworkconst.STRING_PREFIX+"work_sub" {
			paramValue := strings.TrimSpace(item.ParamValue)
			if !strings.HasPrefix(paramValue, "$WORK.") {
				// 修改值并同步到数据库
				paramInputSchema.ParamInputSchemaItems[index] = schema.ParamInputSchemaItem{
					ParamName:  item.ParamName,
					ParamValue: strings.Join([]string{"$WORK.", paramValue}, ""),
				}
				step.WorkStepInput = paramInputSchema.RenderToXml()
				// 自动创建子流程
				checkAndCreateSubWork(paramValue, o)
			}
			// 维护 work 的 WorkSubId 属性
			subWork, _ := iwork.QueryWorkByName(strings.Replace(paramValue, "$WORK.", "", -1), orm.NewOrm())
			step.WorkSubId = subWork.Id
			break
		}
	}
	iwork.InsertOrUpdateWorkStep(&step, o)
}

func CheckAndGetParamValueByInputSchemaParamName(items []schema.ParamInputSchemaItem, paramName string) (exist bool, paramValue string) {
	for _, item := range items {
		if item.ParamName == paramName {
			return true, item.ParamValue
		}
	}
	return false, ""
}
