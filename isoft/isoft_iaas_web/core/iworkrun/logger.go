package iworkrun

import (
	"fmt"
	"isoft/isoft_iaas_web/core/iworklog"
	"time"
)

// 统计操作所花费的时间方法
func recordCostTimeLog(logwriter *iworklog.CacheLoggerWriter, operateName, trackingId string, start time.Time) {
	logwriter.Write(trackingId, fmt.Sprintf(
		"%s total cost time :%v ms", operateName, time.Now().Sub(start).Nanoseconds()/1e6))
}
