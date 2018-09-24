package heartbeat

import (
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/lib"
	"sync"
	"time"
)

// 缓存所有数据节点信息
var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

// 心跳检测
func ListenHeartbeat() {
	// 主要用于接收数据服务节点发送过来的心跳消息
	go ListenDataServerHeartbeat()
	// 清除超过指定时间没收到心跳消息的数据服务节点,默认使用 10 s
	go removeExpiredDataServer()
}

// 主要用于接收数据服务节点发送过来的心跳消息
func ListenDataServerHeartbeat() {
	defer func() {
		if err := recover(); err != nil {
			logutil.Errorln("ListenDataServerHeartbeat error:", err)
			// 异常控制, 5s 后重新心跳检测
			time.Sleep(5 * time.Second)
			ListenDataServerHeartbeat()
		}
	}()
	proxy := &lib.LocateAndHeartbeatProxy{}
	proxy.ReceiveAndModifyHeartbeat(dataServers, &mutex)
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
