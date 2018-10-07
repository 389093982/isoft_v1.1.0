package lib

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/models"
	"isoft/isoft_storage/lib/utils"
	"net/http"
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
		url := fmt.Sprintf("http://%s/api/heartbeat/sendHeartBeat", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
		resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("addr="+cfg.GetConfigValue(cfg.LISTEN_ADDRESS)))
		if err != nil {
			panic(err)
		}
		responseBody, _ := ioutil.ReadAll(resp.Body)
		responseMap := make(map[string]interface{})
		json.Unmarshal(responseBody, &responseMap)
		if responseMap["status"] != "SUCCESS" {
			panic(errors.New(fmt.Sprintf("error:", responseMap["errorMsg"])))
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
	for {
		url := fmt.Sprintf("http://%s/api/heartbeat/queryAllAliveHeartBeat", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		responseBody, _ := ioutil.ReadAll(resp.Body)
		responseMap := make(map[string]interface{})
		json.Unmarshal(responseBody, &responseMap)
		if responseMap["status"] != "SUCCESS" {
			panic(errors.New(fmt.Sprintf("error:", responseMap["errorMsg"])))
		} else {
			heartbeats := responseMap["heartbeats"].([]interface{})
			for _, heartbeat := range heartbeats {
				// 获取监听地址
				dataServer := heartbeat.(map[string]interface{})["addr"].(string)
				mutex.Lock()
				// 更新监听地址的时间
				dataServers[dataServer] = time.Now()
				mutex.Unlock()
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func (this *LocateAndHeartbeatProxy) SendAndReceiveLocateInfo(dataServers []string, hash string, retry int) (locateInfo map[int]string) {
	defer utils.RecordTimeCostForMethod("lib locate_heartbeat SendAndReceiveLocateInfo", time.Now())

	locateInfoStrChan := make(chan string, len(dataServers))
	wg := &sync.WaitGroup{}
	for _, server := range dataServers {
		wg.Add(1)
		go sendAndReceiveLocateInfo(locateInfoStrChan, server, hash, wg)
	}
	wg.Wait()
	close(locateInfoStrChan)
	if len(locateInfoStrChan) != 0 {
		locateInfo = make(map[int]string, len(dataServers))
		// 获取所有定位信息
		for locateInfoStr := range locateInfoStrChan {
			var info models.LocateMessage
			err := json.Unmarshal([]byte(locateInfoStr), &info)
			if err != nil {
				break
			}
			locateInfo[info.ShardId] = info.Addr
		}
	}
	return locateInfo
}

func sendAndReceiveLocateInfo(locateInfoStrChan chan<- string, server string, hash string, wg *sync.WaitGroup) {
	defer wg.Done()
	// 先调用数据服务 temp 接口的 post 方法生产临时文件,接收返回的 uuid 信息
	url := fmt.Sprintf("http://%s/locate/%s", server, hash)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err == nil && resp != nil && resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			locateInfoStrChan <- string(body)
		}
	}
}
