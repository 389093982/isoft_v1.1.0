package iworkrun

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkcomponent"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

func Run(work iwork.Work, steps []iwork.WorkStep) {
	trackingId := stringutil.RandomUUID()
	// 插入 RunLogRecord
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
			fmt.Print(err)
		}
	}()

	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute work:%s",work.WorkName))

	for _, step := range steps {
		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("start execute workstep:%s",step.WorkStepName))

		factory := &iworkcomponent.WorkStepFactory{WorkStep: &step}
		factory.Execute()

		iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute workstep:%s",step.WorkStepName))
	}

	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf("end execute work:%s",work.WorkName))
}


