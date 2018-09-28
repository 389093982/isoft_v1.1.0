package monitor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 每隔 5s 发送一次心跳检测信息
func RecordMonitorHeartBeatLog(isoft_deploy_web_server string, server_address string) {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
			// 间隔 5s 发送一次心跳消息
			time.Sleep(5 * time.Second)
			RecordMonitorHeartBeatLog(isoft_deploy_web_server, server_address)
		}
	}()

	for{
		url := fmt.Sprintf("http://%s/api/monitor/sendMonitorHeartBeat", isoft_deploy_web_server)
		resp, err := http.Post(url,"application/x-www-form-urlencoded", strings.NewReader("addr=" + server_address))
		if err != nil{
			panic(err)
		}
		responseBody, _ := ioutil.ReadAll(resp.Body)
		responseMap := make(map[string]interface{})
		json.Unmarshal(responseBody, &responseMap)
		if responseMap["status"] != "SUCCESS"{
			fmt.Println("error:",responseMap["errorMsg"])
		}
		// 间隔 5s 发送一次心跳消息
		time.Sleep(5 * time.Second)
	}
}
