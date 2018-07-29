package heartbeat

import (
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/rabbitmq"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(cfg.GetConfigValue(cfg.RABBITMQ_SERVER))
	defer q.Close()
	// 无线循环发送心跳信息
	for {
		// 数据服务节点向所有接口服务节点通知自身的存在,每一台数据服务节点都会持续向这个 exchange 发送心跳消息,把本服务的监听地址发送出去
		q.Publish("apiServers", cfg.GetConfigValue(cfg.LISTEN_ADDRESS))
		// 间隔 5s 发送一次心跳消息
		time.Sleep(5 * time.Second)
	}
}
