package executors

import (
	"fmt"
	"isoft/isoft/db"
)

func (this *ExecutorRouter) RunExecuteCommonTask(operate_type string) {
	if operate_type == "mysql_connection_test" {
		_, err := db.GetConnection("root", this.ServiceInfo.MysqlRootPwd,
			this.ServiceInfo.EnvInfo.EnvIp, this.ServiceInfo.ServicePort, "mysql")
		if err != nil {
			this.TrackingLogResolver.WriteErrorLog("mysql_connection_test__FAILED")
			this.TrackingLogResolver.WriteErrorLog(fmt.Sprintf("连接失败,%s", err.Error()))
		} else {
			this.TrackingLogResolver.WriteSuccessLog("mysql_connection_test__SUCCESS")
			this.TrackingLogResolver.WriteSuccessLog("连接成功!")
		}
	}
	// 结束任务
	this.TrackingLogResolver.EndRecordTask()
}
