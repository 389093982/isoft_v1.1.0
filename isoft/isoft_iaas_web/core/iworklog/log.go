package iworklog

import (
	"fmt"
	"isoft/isoft_iaas_web/models/iwork"
	"time"
)

const cacheLen = 10

type CacheLoggerWriter struct {
	caches []*iwork.RunLogDetail
}

func (this *CacheLoggerWriter) cleanCaches() {
	this.caches = make([]*iwork.RunLogDetail, 0)
}

func (this *CacheLoggerWriter) Write(trackingId, detail string) {
	if this.caches == nil {
		this.cleanCaches()
	}
	log := &iwork.RunLogDetail{
		TrackingId:      trackingId,
		Detail:          detail,
		CreatedBy:       "SYSTEM",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "SYSTEM",
		LastUpdatedTime: time.Now(),
	}
	this.caches = append(this.caches, log)
	if len(this.caches) >= cacheLen {
		this.Flush()
		this.cleanCaches()
	}
}

func (this *CacheLoggerWriter) Flush() {
	if _, err := iwork.InsertMultiRunLogDetail(this.caches); err != nil {
		fmt.Println(err.Error())
	}
}

func (this *CacheLoggerWriter) Close() {
	this.Flush()
}
