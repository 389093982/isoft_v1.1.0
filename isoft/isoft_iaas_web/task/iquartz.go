package task

import (
	"fmt"
	"github.com/robfig/cron"
	"isoft/isoft_iaas_web/imodules"
	"isoft/isoft_iaas_web/models/iwork"
)

func StartIQuartzInitialTask() {
	if imodules.CheckModule("iwork"){
		if metas, err := iwork.QueryAllCronMeta(); err == nil {
			c := cron.New()
			for _, meta := range metas {
				c.AddJob(meta.CronStr, &iworkJob{meta: &meta})
			}
			c.Start()
		}
	}
}

type iworkJob struct {
	meta *iwork.CronMeta
}

func (this *iworkJob) Run() {
	fmt.Print(this.meta.CronStr)
}
