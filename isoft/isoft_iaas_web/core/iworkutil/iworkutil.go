package iworkutil

import (
	"encoding/base64"
	"github.com/astaxie/beego/orm"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkconst"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkmodels"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
)

// 用正则表达式匹配出相对变量值
func GetRelativeValueWithReg(s string) []string {
	return stringutil.GetNoRepeatSubStringWithRegexp(s, `\$[a-zA-Z0-9_]+\.[a-zA-Z0-9_]+`)
}

// 用正则表达式匹配出相对变量值,多个值只返回第一个
func GetSingleRelativeValueWithReg(s string) string {
	values := GetRelativeValueWithReg(s)
	if len(values) > 0 {
		return values[0]
	}
	return s
}

func GetWorkSubNameFromParamValue(paramValue string) string {
	value := strings.TrimSpace(paramValue)
	value = strings.Replace(value, "$WORK.", "", -1)
	value = strings.Replace(value, ";", "", -1)
	value = strings.Replace(value, "\n", "", -1)
	value = strings.TrimSpace(value)
	return value
}

func GetWorkSubNameForWorkSubNode(paramInputSchema *iworkmodels.ParamInputSchema) string {
	for _, item := range paramInputSchema.ParamInputSchemaItems {
		if item.ParamName == iworkconst.STRING_PREFIX+"work_sub" && strings.HasPrefix(strings.TrimSpace(item.ParamValue), "$WORK.") {
			// 找到 work_sub 字段值
			return GetWorkSubNameFromParamValue(strings.TrimSpace(item.ParamValue))
		}
	}
	return ""
}

func EncodeToBase64String(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func DecodeBase64String(encodeString string) (bytes []byte) {
	if bytes, err := base64.StdEncoding.DecodeString(encodeString); err != nil {
		return bytes
	}
	return
}

func GetParamValueForWorkVars(paramValue string, store *datastore.DataStore) string {
	paramValue = strings.Replace(paramValue, "$WorkVars.", "", -1)
	paramValue = strings.Replace(paramValue, ";", "", -1)
	return store.GetData("__workVars__", paramValue).(string)
}

func GetParamValueForEntity(paramValue string) string {
	paramValue = strings.TrimSpace(paramValue)
	paramValue = strings.Replace(paramValue, ";", "", -1)
	if !strings.HasPrefix(paramValue, "$Entity.") {
		return paramValue
	}
	entity_name := strings.Replace(paramValue, "$Entity.", "", -1)
	if entity, err := iwork.QueryEntityByEntityName(entity_name); err == nil {
		return entity.EntityFieldStr
	}
	return ""
}

func GetAllPreStepNodeName(work_id, work_step_id int64) []string {
	result := make([]string, 0)
	steps, err := iwork.QueryAllPreStepInfo(work_id, work_step_id, orm.NewOrm())
	if err == nil {
		// 当前步骤信息
		currentWorkStep, _ := iwork.QueryWorkStepInfo(work_id, work_step_id, orm.NewOrm())
		// 所有步骤信息
		allSteps, _ := iwork.QueryAllWorkStepInfo(work_id, orm.NewOrm())
		parser := block.BlockParser{Steps: allSteps}
		currentBlockStep, allBlockSteps := parser.ParseAndGetCurrentBlockStep(&currentWorkStep)
		for _, step := range steps {
			// 判断前置 step 在块范围内是否是可访问的
			if block.CheckBlockAccessble(allBlockSteps, currentBlockStep, step.WorkStepId) {
				result = append(result, step.WorkStepName)
			}
		}
	}
	return result
}
