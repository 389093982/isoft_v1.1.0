package utils

import (
	"isoft/isoft/common/logutil"
	"time"
)

func RecordTimeCostForMethod(label string, startTime time.Time) {
	logutil.Infoln(label, "cost time", time.Now().Sub(startTime).Nanoseconds()/1e6, "ms")
}
