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

	if (strings.TrimSpace(m["create_account"]) == "" || strings.TrimSpace(m["create_passwd"]) == "") &&
		strings.TrimSpace(m["create_database"]) == "" {
		return errors.New("参数不合法!")
	}

	if strings.TrimSpace(m["create_account"]) != "" && strings.TrimSpace(m["create_passwd"]) != "" {
		// 创建用户
		// 将单个的%转换为%%,而%%又会被当做字面%打印,避免问题出现
		sql := fmt.Sprintf("CREATE USER '%s'@'%%' IDENTIFIED BY '%s'",
			strings.TrimSpace(m["create_account"]), strings.TrimSpace(m["create_passwd"]))
		this.TrackingLogResolver.WriteSuccessLog(sql)
		if _, err := db.Exec(sql); err != nil {
			this.TrackingLogResolver.WriteErrorLog(err.Error())
		} else {
			this.TrackingLogResolver.WriteSuccessLog("创建、初始化用户成功!")
		}
	}

	if strings.TrimSpace(m["create_database"]) != "" {
		// 创建数据库
		sql := fmt.Sprintf("create database %s", strings.TrimSpace(m["create_database"]))
		this.TrackingLogResolver.WriteSuccessLog(sql)
		if _, err := db.Exec(sql); err != nil {
			this.TrackingLogResolver.WriteErrorLog(err.Error())
		} else {
			this.TrackingLogResolver.WriteSuccessLog("创建数据库成功!")
		}
	}

	if strings.TrimSpace(m["create_database"]) != "" && strings.TrimSpace(m["create_account"]) != "" &&
		strings.TrimSpace(m["create_passwd"]) != "" {
		// 添加授权信息
		sql := fmt.Sprintf("GRANT ALL PRIVILEGES ON %s.* TO '%s'@'%%'",
			strings.TrimSpace(m["create_database"]), strings.TrimSpace(m["create_account"]))
		this.TrackingLogResolver.WriteSuccessLog(sql)
		if _, err = db.Exec(sql); err != nil {
			this.TrackingLogResolver.WriteErrorLog(err.Error())
		} else {
			this.TrackingLogResolver.WriteSuccessLog("为用户添加数据库访问授权成功!")
		}
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
