package iworkrun

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

// dispatcher 为父流程遗传下来的参数
func Run(work iwork.Work, steps []iwork.WorkStep, dispatcher *entry.Dispatcher) (receiver *entry.Receiver) {
	// 当前流程的 trackingId
	trackingId := stringutil.RandomUUID()
	if dispatcher != nil && dispatcher.TrackingId != "" {
		// 拼接父流程的 trackingId 信息,作为链式 trackingId
		trackingId = fmt.Sprintf("%s.%s", dispatcher.TrackingId, trackingId)
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

	start := time.Now()
	defer func() {
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("task total cost time :%v ms",time.Now().Sub(start).Nanoseconds() / 1e6))
	}()

	defer func() {
		if err := recover(); err != nil {
			iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("internal error:%s", err))
		}
	}()
	// 记录日志详细
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("~~~~~~~~~~start execute work:%s~~~~~~~~~~", work.WorkName))
	// 申请数据中心存储中间数据
	datastore.RegistDataStore(trackingId)
	// 逐步执行步骤
	for _, step := range steps {
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute workstep: >> [[%s]]", step.WorkStepName))
		// 由工厂代为执行步骤
		factory := &iworknode.WorkStepFactory{WorkStep: &step, RunFunc: Run, Dispatcher: dispatcher}
		factory.Execute(trackingId)
		// factory 节点如果代理的是 work_end 节点,则传递 Receiver 出去
		if factory.Receiver != nil {
			receiver = factory.Receiver
		}
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute workstep: >> [[%s]]", step.WorkStepName))
	}
	// 注销数据中心
	datastore.UnRegistDataStore(trackingId)
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("~~~~~~~~~~end execute work:%s~~~~~~~~~~", work.WorkName))
	return
}
