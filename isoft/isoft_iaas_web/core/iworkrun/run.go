package iworkrun

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkcomponent"
	"isoft/isoft_iaas_web/core/iworkdata/datastore"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func Run(work iwork.Work, steps []iwork.WorkStep) {
	trackingId := stringutil.RandomUUID()
	// 记录日志
	iwork.InsertRunLogRecord(&iwork.RunLogRecord{
		TrackingId:trackingId,
		WorkName:work.WorkName,
		CreatedBy:"SYSTEM",
		CreatedTime:time.Now(),
		LastUpdatedBy:"SYSTEM",
		LastUpdatedTime:time.Now(),
	})

	defer func() {
		if err := recover(); err != nil {
			iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("internal error:%s",err))
		}
	}()
	// 记录日志详细
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute work:%s",work.WorkName))
	// 申请数据中心存储中间数据
	datastore.RegistDataStore(trackingId)
	// 逐步执行步骤
	for _, step := range steps {
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute workstep:%s",step.WorkStepName))
		// 由工厂代为执行步骤
		factory := &iworkcomponent.WorkStepFactory{WorkStep: &step}
		factory.Execute(trackingId)
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute workstep:%s",step.WorkStepName))
	}
	// 注销数据中心
	datastore.UnRegistDataStore(trackingId)
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute work:%s",work.WorkName))
}


