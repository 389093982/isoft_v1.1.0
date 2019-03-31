package iworkrun

import (
	"fmt"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

// 统计操作所花费的时间方法
func recordCostTimeLog(operateName, trackingId string, start time.Time) {
	iwork.InsertRunLogDetail(trackingId, fmt.Sprintf(
		"%s total cost time :%v ms", operateName, time.Now().Sub(start).Nanoseconds()/1e6))
}
