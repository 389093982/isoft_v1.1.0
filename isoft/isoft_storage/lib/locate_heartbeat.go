package lib

import (
	"encoding/json"
	"fmt"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/models"
	"isoft/isoft_storage/lib/rabbitmq"
	"isoft/isoft_storage/lib/rs"
	"strconv"
	"sync"
	"time"
)

type LocateAndHeartbeatProxy struct {
	
}

func (this *LocateAndHeartbeatProxy) SendHeartbeat()  {
	defer func() {
		if err := recover(); err != nil{
			// 出现异常时需要重新执行
			this.SendHeartbeat()
		}
	}()
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

func (this *LocateAndHeartbeatProxy) ReceiveDealAndSendLocateInfo(locateFunc func(hash string) int)  {
	defer func() {
		if err := recover(); err != nil{
			// 出现异常时需要重新执行
			this.ReceiveDealAndSendLocateInfo(locateFunc)
		}
	}()
	// 直接将 RabbitMQ 消息队列里收到的对象散列值作为 Locate 参数
	q := rabbitmq.New(cfg.GetConfigValue(cfg.RABBITMQ_SERVER))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		// 接收 hash 值
		hash, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		// 定位 hash 值是否存在
		id := locateFunc(hash)
		if id != -1 {
			// 不存在则不返回消息,存在则返回消息
			q.Send(msg.ReplyTo, models.LocateMessage{Addr: cfg.GetConfigValue(cfg.LISTEN_ADDRESS), ShardId: id})
		}
	}
}

func (this *LocateAndHeartbeatProxy) ReceiveAndModifyHeartbeat(dataServers map[string]time.Time, mutex *sync.Mutex)  {
	defer func() {
		if err := recover(); err != nil{
			// 异常退出后需要重新执行
			this.ReceiveAndModifyHeartbeat(dataServers, mutex)
		}
	}()
	q := rabbitmq.New(cfg.GetConfigValue(cfg.RABBITMQ_SERVER))
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()
	// 循环遍历数据服务监听地址
	for msg := range c {
		// 获取监听地址
		dataServer, e := strconv.Unquote(string(msg.Body))
		fmt.Println("receive dataServer :", dataServer)
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		// 更新监听地址的时间
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

// 发送和接收定位失败需要进行重试,最多重试 retry 次
func (this *LocateAndHeartbeatProxy) RetrySendAndReceiveLocateInfo(hash string, retry int) (locateInfo map[int]string) {
	if retry <= 0{
		return
	}
	defer func() {
		if err := recover(); err != nil{
			retry--
			locateInfo = this.RetrySendAndReceiveLocateInfo(hash, retry)
			return
		}
	}()
	q := rabbitmq.New(cfg.GetConfigValue(cfg.RABBITMQ_SERVER))
	q.Publish("dataServers", hash)
	c := q.Consume()
	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()
	locateInfo = make(map[int]string)
	// 循环获取 6 条定位消息
	for i := 0; i < rs.ALL_SHARDS; i++ {
		msg := <-c
		if len(msg.Body) == 0 {
			return
		}
		var info models.LocateMessage
		json.Unmarshal(msg.Body, &info)
		locateInfo[info.ShardId] = info.Addr
	}
	return
}
