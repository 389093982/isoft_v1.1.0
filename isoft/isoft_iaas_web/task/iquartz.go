package task

import (
	"fmt"
	"github.com/robfig/cron"
	"isoft/isoft_iaas_web/models/iquartz"
)

func StartIQuartzInitialTask()  {
	if metas, err := iquartz.GetAllCronMeta();err == nil{
		c := cron.New()
		for _,meta := range metas{
			c.AddJob(meta.CronStr, &IQuartzJob{meta:&meta})
		}
		c.Start()
	}
}

type IQuartzJob struct{
	meta *iquartz.CronMeta
}

func (this *IQuartzJob) Run()  {
	fmt.Print(this.meta.CronStr)
}