package task

import "github.com/astaxie/beego/toolbox"

// 运行定时任务
func RunCronTask() {
	// 创建服务监控任务,每分钟触发一次
	tk := toolbox.NewTask("serviceMonitorTask", "0 * * * * *", func() error {
		return RunServiceMonitorTask()
	})
	toolbox.AddTask("serviceMonitorTask", tk)
	toolbox.StartTask()
}
