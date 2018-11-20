package weedfsutil

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

// 封装 WeedFs 信息实体类
type WeedFsInfo struct {
	Fid       string `json:"fid"`
	Url       string `json:"url"`
	PublicUrl string `json:"publicUrl"`
	Count     int8   `json:"count"`
}

// 调用 /dir/assign 接口获取 fid
func getWeedFsInfo(masterAddress string) (weedFsInfo WeedFsInfo, err error) {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", fmt.Sprintf("http://%s/dir/assign", masterAddress), nil)
	if err != nil {
		return
	}
	response, err := client.Do(reqest)
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != 200{
		return weedFsInfo, errors.New(fmt.Sprintf("getWeedFsInfo error with %d, %s", response.StatusCode, string(body)))
	}
	// 格式 {"fid":"1,0a42e9808d","url":"127.0.0.1:9080","publicUrl":"127.0.0.1:9080","count":1}
	err = json.Unmarshal(body, &weedFsInfo)
	if err == nil {
		// 将 127.0.0.1 的 ip 地址换成 ipv4 地址
		weedFsInfo.Url = strings.Replace(weedFsInfo.Url, "127.0.0.1", strings.Split(masterAddress, ":")[0], -1)
		weedFsInfo.PublicUrl = strings.Replace(weedFsInfo.PublicUrl, "127.0.0.1", strings.Split(masterAddress, ":")[0], -1)
	}
	return
}

// 实际存储文件的方法
func SaveFile(masterAddress string, file multipart.File) (weedFsInfo WeedFsInfo, err error) {
	weedFsInfo, err = getWeedFsInfo(masterAddress)
	if err != nil {
		return
	}
	bReader := bufio.NewReader(file)
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/%s", weedFsInfo.PublicUrl, weedFsInfo.Fid), bReader)
	if err != nil {
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 201 { // 201 Created
		return weedFsInfo, errors.New(fmt.Sprintf("save file failed, error status code %d", res.StatusCode))
	}
	return
}


