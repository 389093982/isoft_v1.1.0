package iworkrun

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworkdata/block"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworkdata/memory"
	"isoft/isoft_iaas_web/core/iworklog"
	"isoft/isoft_iaas_web/core/iworkplugin/iworknode"
	"isoft/isoft_iaas_web/core/iworkutil/datatypeutil"
	"isoft/isoft_iaas_web/core/iworkutil/errorutil"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

// dispatcher 为父流程遗传下来的参数
func Run(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) {
	logwriter := new(iworklog.CacheLoggerWriter)
	defer logwriter.Close()

	// 为当前流程创建新的 trackingId
	trackingId := createNewTrackingIdForWork(dispatcher, work)
	defer logwriter.RecordCostTimeLog("execute work", trackingId, time.Now())
	defer func() {
		if err := recover(); err != nil {
			// 记录 4 kb大小的堆栈信息
			logwriter.Write(trackingId, "~~~~~~~~~~~~~~~~~~~~~~~~ internal error trace stack ~~~~~~~~~~~~~~~~~~~~~~~~~~")
			logwriter.Write(trackingId, string(errorutil.PanicTrace(4)))
			logwriter.Write(trackingId, fmt.Sprintf("<span style='color:red;'>internal error:%s</span>", err))
			logwriter.Write(trackingId, "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}
	}()
	// 记录日志详细
	logwriter.Write(trackingId, fmt.Sprintf("~~~~~~~~~~start execute work:%s~~~~~~~~~~", work.WorkName))
	// 获取数据中心
	store := datastore.InitDataStore(trackingId, logwriter)

	// 将 steps 转换成 BlockSteps
	// 逐个 block 依次执行
	for _, blockStep := range getExecuteOrder(steps) {
		if blockStep.Step.WorkStepType != "empty" {
			receiver = RunOneStep(trackingId, logwriter, blockStep, store, dispatcher)
		}
	}

	// 注销 MemoryCache,无需注册,不存在时会自动注册
	memory.UnRegistMemoryCache(trackingId)
	logwriter.Write(trackingId, fmt.Sprintf("~~~~~~~~~~end execute work:%s~~~~~~~~~~", work.WorkName))
	return
}

func getExecuteOrder(steps []iwork.WorkStep) []*block.BlockStep {
	order := make([]*block.BlockStep, 0)
	deferOrder := make([]*block.BlockStep, 0)
	var end *block.BlockStep
	for _, blockStep := range block.ParseToBlockStep(steps) {
		if blockStep.Step.IsDefer == "true" {
			deferOrder = append(deferOrder, blockStep)
		} else if blockStep.Step.WorkStepType == "work_end" {
			end = blockStep
		} else {
			order = append(order, blockStep)
		}
	}
	// is_defer 和 work_end 都是需要延迟执行
	order = append(order, datatypeutil.ReverseSlice(deferOrder).([]*block.BlockStep)...)
	order = append(order, end)
	return order
}

// 执行单个 BlockStep
func RunOneStep(trackingId string, logwriter *iworklog.CacheLoggerWriter, blockStep *block.BlockStep,
	datastore *datastore.DataStore, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) {
	// 统计耗费时间
	defer logwriter.RecordCostTimeLog(blockStep.Step.WorkStepName, trackingId, time.Now())
	// 记录开始执行日志
	logwriter.Write(trackingId, fmt.Sprintf("start execute blockStep: >>>>>>>>>> [[<span style='color:blue;'>%s<span>]]", blockStep.Step.WorkStepName))
	// 由工厂代为执行步骤
	factory := &iworknode.WorkStepFactory{
		WorkStep:         blockStep.Step,
		WorkSubRunFunc:   Run,
		Dispatcher:       dispatcher,
		Receiver:         receiver,
		BlockStep:        blockStep,
		BlockStepRunFunc: RunOneStep,
		DataStore:        datastore,
		LogWriter:        logwriter,
	}
	factory.Execute(trackingId)
	// 记录结束执行日志
	logwriter.Write(trackingId, fmt.Sprintf("end execute blockStep: >>>>>>>>>> [[%s]]", blockStep.Step.WorkStepName))
	// factory 节点如果代理的是 work_end 节点,则传递 Receiver 出去
	return factory.Receiver
}
