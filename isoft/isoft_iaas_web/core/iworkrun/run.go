package iworkrun

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/memory"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

// dispatcher 为父流程遗传下来的参数
func Run(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) {
	trackingId := createNewTrackingIdForWork(dispatcher, work)
	defer recordCostTimeLog("execute work", trackingId, time.Now())
	defer func() {
		if err := recover(); err != nil {
			iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("internal error:%s", err))
		}
	}()
	// 记录日志详细
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("~~~~~~~~~~start execute work:%s~~~~~~~~~~", work.WorkName))
	// 逐步执行步骤
	for _, step := range steps {
		if step.WorkStepType == "empty" {continue}
		// 获取数据中心
		store := datastore.GetDataStore(trackingId)
		if redirectNodeName, ok := store.GetData("__goto_condition__","__redirect__").(string);
			ok && strings.TrimSpace(redirectNodeName) != ""{
			if step.WorkStepName == redirectNodeName{
				// 相等代表刚好调到 redirect 节点,此时要将 store 里面的跳转信息置空
				store.CacheData("__goto_condition__","__redirect__", "")
			}else{
				iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("The step for %s was skipped!", step.WorkStepName))
				// 不相等代表还没有调到 redirect 节点,此时直接跳过, redirect 节点值为 __out__ 时,所用节点都匹配不上,刚好表示为跳出流程
				continue
			}
		}

		_receiver := runOneStep(trackingId, &step, dispatcher)
		if _receiver != nil {
			receiver = _receiver
		}
	}
	// 注销数据中心,无需注册,不存在时会自动注册
	datastore.UnRegistDataStore(trackingId)
	// 注销 MemoryCache,无需注册,不存在时会自动注册
	memory.UnRegistMemoryCache(trackingId)
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("~~~~~~~~~~end execute work:%s~~~~~~~~~~", work.WorkName))
	return
}

func optimizeTrackingId(pTrackingId, trackingId string) string {
	if strings.Count(pTrackingId, ".") > 1{
	// a.~.b.c
		trackingId = strings.Join(
			[]string{
				pTrackingId[:strings.Index(pTrackingId, ".")],			// 顶级 trackingId
				"~",															// 过渡级 trackingId
				pTrackingId[strings.LastIndex(pTrackingId, ".") + 1 :],	// 父级 trackingId
				trackingId,														// 当前级 trackingId
			},".")
	}else{
		trackingId = fmt.Sprintf("%s.%s", pTrackingId, trackingId)
	}
	return trackingId
}

func createNewTrackingIdForWork(dispatcher *entry.Dispatcher, work iwork.Work) string {
	// 生成当前流程的 trackingId
	trackingId := stringutil.RandomUUID()
	// 调度者不为空时代表有父级流程
	if dispatcher != nil && dispatcher.TrackingId != "" {
		// 拼接父流程的 trackingId 信息,作为链式 trackingId
		// 同时优化 trackingId,防止递归调用时 trackingId 过长
		trackingId = optimizeTrackingId(dispatcher.TrackingId, trackingId)
	}
	// 记录日志
	iwork.InsertRunLogRecord(&iwork.RunLogRecord{
		TrackingId:      trackingId,
		WorkName:        work.WorkName,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	})
	return trackingId
}

// 执行单个步骤
func runOneStep(trackingId string, step *iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) {
	defer recordCostTimeLog(step.WorkStepName, trackingId, time.Now())
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute workstep: >>>>>>>>>> [[%s]]", step.WorkStepName))
	// 由工厂代为执行步骤
	factory := &iworknode.WorkStepFactory{WorkStep: step, RunFunc: Run, Dispatcher: dispatcher}
	factory.Execute(trackingId)
	// factory 节点如果代理的是 work_end 节点,则传递 Receiver 出去
	if factory.Receiver != nil {
		receiver = factory.Receiver
	}
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute workstep: >>>>>>>>>> [[%s]]", step.WorkStepName))
	return
}

// 统计操作所花费的时间方法
func recordCostTimeLog(operateName, trackingId string, start time.Time) {
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf(
		"%s total cost time :%v ms",operateName, time.Now().Sub(start).Nanoseconds() / 1e6))
}