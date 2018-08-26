package executors

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"isoft/isoft/db"
	"strings"
)

func (this *ExecutorRouter) RunExecuteCommonTask(operate_type, extra_params string) {
	var err error
	switch operate_type {
	case "mysql_connection_test":
		err = this.MysqlConnectionTest(operate_type)
	case "mysql_init":
		err = this.MysqlInit(operate_type, extra_params)
	}

	if err == nil {
		this.TrackingLogResolver.WriteSuccessLog(operate_type + "__SUCCESS")
	} else {
		this.TrackingLogResolver.WriteErrorLog(operate_type + "__FAILED")
		this.TrackingLogResolver.WriteErrorLog(err.Error())
	}
	// 结束任务
	this.TrackingLogResolver.EndRecordTask()
}

func (this *ExecutorRouter) MysqlInit(operate_type, extra_params string) error {
	db, err := db.GetConnection("root", this.ServiceInfo.MysqlRootPwd,
		this.ServiceInfo.EnvInfo.EnvIp, this.ServiceInfo.ServicePort, "mysql")
	defer db.Close()
	if err != nil {
		return err
	}
	// json 格式参数 extra_params 转换
	m := make(map[string]string)
	err = json.Unmarshal([]byte(extra_params), &m)
	if err != nil {
		return err
	}
	// 创建数据库
	if value, ok := m["create_database"]; ok && strings.TrimSpace(value) != "" {
		if _, err := db.Exec(fmt.Sprintf("create database %s", strings.TrimSpace(value))); err != nil {
			return err
		} else {
			this.TrackingLogResolver.WriteSuccessLog("创建数据库成功!")
		}
	} else {
		return errors.New("参数 create_database 不合法!")
	}
	// 创建用户
	create_account, ok1 := m["create_account"]
	create_passwd, ok2 := m["create_passwd"]
	if ok1 && ok2 && strings.TrimSpace(create_account) != "" && strings.TrimSpace(create_passwd) != "" {
		if _, err := db.Exec(fmt.Sprintf("CREATE USER '%s'@'%' IDENTIFIED BY '%s'", create_account, create_passwd)); err != nil {
			return err
		} else {
			this.TrackingLogResolver.WriteSuccessLog("创建、初始化用户成功!")
		}
	} else {
		return errors.New("参数 create_account 或 create_passwd 不合法!")
	}
	return err
}

func (this *ExecutorRouter) MysqlConnectionTest(operate_type string) (err error) {
	db, err := db.GetConnection("root", this.ServiceInfo.MysqlRootPwd,
		this.ServiceInfo.EnvInfo.EnvIp, this.ServiceInfo.ServicePort, "mysql")
	defer db.Close()
	if err != nil {
		return
	}
	_, err = db.Exec("select * from user")
	if err != nil {
		return
	}
	this.TrackingLogResolver.WriteSuccessLog("连接成功！")
	return
}
