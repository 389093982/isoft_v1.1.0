package heartbeat

import (
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/rabbitmq"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(cfg.GetConfigValue(cfg.RABBITMQ_SERVER))
	defer q.Close()
	for {
		q.Publish("apiServers", cfg.GetConfigValue(cfg.LISTEN_ADDRESS))
		time.Sleep(5 * time.Second)
	}
}
