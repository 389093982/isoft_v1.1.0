package iworkrun

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/memory"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/core/iworkutil/errorutil"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

// dispatcher 为父流程遗传下来的参数
func Run(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) {
	// 为当前流程创建新的 trackingId
	trackingId := createNewTrackingIdForWork(dispatcher, work)
	defer recordCostTimeLog("execute work", trackingId, time.Now())
	defer func() {
		if err := recover(); err != nil {
			// 记录 4 kb大小的堆栈信息
			iwork.InsertRunLogDetail(trackingId, "~~~~~~~~~~~~~~~~~~~~~~~~ internal error trace stack ~~~~~~~~~~~~~~~~~~~~~~~~~~")
			iwork.InsertRunLogDetail(trackingId, string(errorutil.PanicTrace(4)))
			iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("internal error:%s", err))
			iwork.InsertRunLogDetail(trackingId, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}
	}()
	// 记录日志详细
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("~~~~~~~~~~start execute work:%s~~~~~~~~~~", work.WorkName))
	// 获取数据中心
	store := datastore.InitDataStore(trackingId)
	// 将 steps 转换成 BlockSteps
	// 逐个 block 依次执行
	for _, blockStep := range block.ParseToBlockStep(steps) {
		if blockStep.Step.WorkStepType == "empty" {
			continue
		}
		if redirectNodeName, ok := store.GetData("__goto_condition__", "__redirect__").(string); ok && strings.TrimSpace(redirectNodeName) != "" {
			if blockStep.Step.WorkStepName == redirectNodeName {
				// 相等代表刚好调到 redirect 节点,此时要将 store 里面的跳转信息置空
				store.CacheData("__goto_condition__", "__redirect__", "")
			} else {
				iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("The step for %s was skipped!", blockStep.Step.WorkStepName))
				// 不相等代表还没有调到 redirect 节点,此时直接跳过, redirect 节点值为 __out__ 时,所用节点都匹配不上,刚好表示为跳出流程
				continue
			}
		}

		_receiver := RunOneStep(trackingId, blockStep, store, dispatcher)
		if _receiver != nil {
			receiver = _receiver
		}
	}
	// 注销 MemoryCache,无需注册,不存在时会自动注册
	memory.UnRegistMemoryCache(trackingId)
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("~~~~~~~~~~end execute work:%s~~~~~~~~~~", work.WorkName))
	return
}

func optimizeTrackingId(pTrackingId, trackingId string) string {
	if strings.Count(pTrackingId, ".") > 1 {
		// a.~.b.c
		trackingId = strings.Join(
			[]string{
				pTrackingId[:strings.Index(pTrackingId, ".")], // 顶级 trackingId
				"~", // 过渡级 trackingId
				pTrackingId[strings.LastIndex(pTrackingId, ".")+1:], // 父级 trackingId
				trackingId, // 当前级 trackingId
			}, ".")
	} else {
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

// 执行单个 BlockStep
func RunOneStep(trackingId string, blockStep *block.BlockStep, datastore *datastore.DataStore, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) {
	defer recordCostTimeLog(blockStep.Step.WorkStepName, trackingId, time.Now())
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute blockStep: >>>>>>>>>> [[%s]]", blockStep.Step.WorkStepName))
	// 由工厂代为执行步骤
	factory := &iworknode.WorkStepFactory{
		WorkStep:         blockStep.Step,
		WorkSubRunFunc:   Run,
		Dispatcher:       dispatcher,
		BlockStep:        blockStep,
		BlockStepRunFunc: RunOneStep,
		DataStore:        datastore,
	}
	factory.Execute(trackingId)
	// factory 节点如果代理的是 work_end 节点,则传递 Receiver 出去
	if factory.Receiver != nil {
		receiver = factory.Receiver
	}
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute blockStep: >>>>>>>>>> [[%s]]", blockStep.Step.WorkStepName))
	return
}

// 统计操作所花费的时间方法
func recordCostTimeLog(operateName, trackingId string, start time.Time) {
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf(
		"%s total cost time :%v ms", operateName, time.Now().Sub(start).Nanoseconds()/1e6))
}
