package iworkrun

import (
	"fmt"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_iaas_web/core/iworkdata/entry"
	"isoft/isoft_iaas_web/models/iwork"
	"strings"
	"time"
)

// 获取当前 work 需要的 trakingId
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
