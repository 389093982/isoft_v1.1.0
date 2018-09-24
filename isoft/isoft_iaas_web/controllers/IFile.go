package controllers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"isoft/isoft/common/hashutil"
	"net/http"
	"net/url"
	"strings"
)

var isoft_istorage_web string

func init() {
	isoft_istorage_web = beego.AppConfig.String("isoft_istorage_web")
}

type IFileController struct {
	beego.Controller
}

func (this *IFileController) FileUpload() {
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
			this.ServeJSON()
			return
		}
	}()
	// 判断是否是文件上传
	f, h, err := this.GetFile("file")
	if err != nil {
		panic(err)
	}

	bReader := bytes.Buffer{}
	// io.TeeReader、io.MultiReader
	reader := io.TeeReader(f, &bReader)
	// 定位对象用对象名,存储对象用 hash 值
	hash := hashutil.CalculateHash(reader)

	// 调用 isoft_istorage_web 发送 put 请求调用分布式对象存储接口
	url := fmt.Sprintf("%s/objects/%s", isoft_istorage_web, url.PathEscape(h.Filename))
	req, err := http.NewRequest("PUT", url, &bReader)
	if err != nil {
		panic(err)
	}
	// 在请求头中添加 hash 值
	req.Header.Add("digest", "SHA-256="+hash)
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "Status": res.Status, "body": body, "filename": h.Filename}
	this.ServeJSON()
}

// 分片定位
func (this *IFileController) LocateShards() {
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": err}
			this.ServeJSON()
		}
	}()
	hash := strings.TrimSpace(this.GetString("hash", ""))
	url := fmt.Sprintf("%s/locate/%s", isoft_istorage_web, hash)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var locateShardsMap map[string]interface{}
	err = json.Unmarshal(body, &locateShardsMap)
	if err != nil {
		panic(err)
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "shards": &locateShardsMap}
	this.ServeJSON()
}

func (this *IFileController) FileDownload() {
	name := strings.TrimSpace(this.GetString("name", ""))
	version := strings.TrimSpace(this.GetString("version", ""))
	url := fmt.Sprintf("%s/objects/%s?version=%s", isoft_istorage_web, name, version)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		panic(err)
	}
	raw := res.Body
	defer raw.Close()
	reader := bufio.NewReaderSize(raw, 1024*1024*10)
	this.Ctx.ResponseWriter.Header().Set("Content-Type", "application/octet-stream")
	this.Ctx.ResponseWriter.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", name))
	io.Copy(this.Ctx.ResponseWriter, reader)
}
