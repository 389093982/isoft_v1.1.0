package lib

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/models"
	"isoft/isoft_storage/lib/rabbitmq"
	"isoft/isoft_storage/lib/rs"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)


type LocateAndHeartbeatProxy struct {

}

func (this *LocateAndHeartbeatProxy) SendHeartbeat() {
	defer func() {
		if err := recover(); err != nil {
			time.Sleep(5 * time.Second)
			// 出现异常时需要重新执行,间隔 5s
			this.SendHeartbeat()
		}
	}()
	// 无线循环发送心跳信息
	for {
		url := fmt.Sprintf("http://%s/api/heartbeat/sendHeartBeat", cfg.GetConfigValue(cfg.ISOFT_STORAGE_API))
		resp, err := http.Post(url,"application/x-www-form-urlencoded", strings.NewReader("addr=" + cfg.GetConfigValue(cfg.LISTEN_ADDRESS)))
		if err != nil{
			panic(err)
		}
		responseBody, _ := ioutil.ReadAll(resp.Body)
		responseMap := make(map[string]interface{})
		json.Unmarshal(responseBody, &responseMap)
		if responseMap["status"] != "SUCCESS"{
			panic(errors.New(fmt.Sprintf("error:",responseMap["errorMsg"])))
		}
		// 间隔 5s 发送一次心跳消息
		time.Sleep(5 * time.Second)
	}
}

func (this *LocateAndHeartbeatProxy) ReceiveAndModifyHeartbeat(dataServers map[string]time.Time, mutex *sync.Mutex) {
	defer func() {
		if err := recover(); err != nil {
			// 异常退出后需要重新执行
			time.Sleep(5 * time.Second)
			this.ReceiveAndModifyHeartbeat(dataServers, mutex)
		}
	}()

	// 每隔 5 s 循环查询一次心跳信息
	for{
		url := fmt.Sprintf("http://%s/api/heartbeat/queryAllAliveHeartBeat", cfg.GetConfigValue(cfg.ISOFT_STORAGE_API))
		resp, err := http.Get(url)
		if err != nil{
			panic(err)
		}
		responseBody, _ := ioutil.ReadAll(resp.Body)
		responseMap := make(map[string]interface{})
		json.Unmarshal(responseBody, &responseMap)
		if responseMap["status"] != "SUCCESS"{
			panic(errors.New(fmt.Sprintf("error:",responseMap["errorMsg"])))
		}else{
			heartbeats := responseMap["heartbeats"].([]interface {})
			for _, heartbeat := range heartbeats{
				// 获取监听地址
				dataServer := heartbeat.(map[string]interface{})["addr"].(string)
				fmt.Println("receive dataServer :", dataServer)
				mutex.Lock()
				// 更新监听地址的时间
				dataServers[dataServer] = time.Now()
				mutex.Unlock()
			}
		}
		time.Sleep(5 * time.Second)
	}
}

// 发送和接收定位失败需要进行重试,最多重试 retry 次
func (this *LocateAndHeartbeatProxy) RetrySendAndReceiveLocateInfo(hash string, retry int) (locateInfo map[int]string) {
	if retry <= 0 {
		return
	}
	defer func() {
		if err := recover(); err != nil {
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

func (this *LocateAndHeartbeatProxy) ReceiveDealAndSendLocateInfo(locateFunc func(hash string) int) {
	defer func() {
		if err := recover(); err != nil {
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
