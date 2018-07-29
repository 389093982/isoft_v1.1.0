package heartbeat

import (
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/rabbitmq"
	"strconv"
	"sync"
	"time"
)

// 缓存所有数据节点信息
var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

// 主要用于接收数据服务节点发送过来的心跳消息
func ListenHeartbeat() {
	q := rabbitmq.New(cfg.GetConfigValue(cfg.RABBITMQ_SERVER))
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()

	// 清除超过指定时间没收到心跳消息的数据服务节点,默认使用 10 s
	go removeExpiredDataServer()

	// 循环遍历数据服务监听地址
	for msg := range c {
		// 获取监听地址
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		// 更新监听地址的时间
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

// 清除超过指定时间没收到心跳消息的数据服务节点,默认使用 10 s
func removeExpiredDataServer() {
	for {
		// 每隔 5s 清除一次
		time.Sleep(5 * time.Second)
		mutex.Lock()
		for s, t := range dataServers {
			// 判断是否超过 10 s
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

// 遍历 dataServers 并返回当前所有的数据服务节点
func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s := range dataServers {
		ds = append(ds, s)
	}
	return ds
}
