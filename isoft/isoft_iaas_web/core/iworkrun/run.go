package iworkrun

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/core/iworknode"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

// args 为父流程遗传下来的参数
func Run(work iwork.Work, steps []iwork.WorkStep, args ...interface{}) {
	// 当前流程的 trackingId
	trackingId := stringutil.RandomUUID()
	if len(args) > 0{
		// 拼接父流程的 trackingId 信息,作为链式 trackingId
		trackingId = fmt.Sprintf("%s.%s",args[0].(string), trackingId)
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

	defer func() {
		if err := recover(); err != nil {
			iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("internal error:%s", err))
		}
	}()
	// 记录日志详细
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute work:%s", work.WorkName))
	// 申请数据中心存储中间数据
	datastore.RegistDataStore(trackingId)
	// 逐步执行步骤
	for _, step := range steps {
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute workstep:%s", step.WorkStepName))
		// 由工厂代为执行步骤
		factory := &iworknode.WorkStepFactory{WorkStep: &step, RunFunc: Run}
		factory.Execute(trackingId)
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute workstep:%s", step.WorkStepName))
	}
	// 注销数据中心
	datastore.UnRegistDataStore(trackingId)
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute work:%s", work.WorkName))
}
