package iworkrun

import (
	"fmt"
	"strings"
)

// 对 trakingId 进行优化,避免过长的 trackingId
func optimizeTrackingId(pTrackingId, trackingId string) string {
	if strings.Count(pTrackingId, ".") <= 1 {
		return fmt.Sprintf("%s.%s", pTrackingId, trackingId)
	}
	// a.~.b.c
	trackingId = strings.Join(
		[]string{
			pTrackingId[:strings.Index(pTrackingId, ".")], // 顶级 trackingId
			"~", // 过渡级 trackingId
			pTrackingId[strings.LastIndex(pTrackingId, ".")+1:], // 父级 trackingId
			trackingId, // 当前级 trackingId
		}, ".")
	return trackingId
}
