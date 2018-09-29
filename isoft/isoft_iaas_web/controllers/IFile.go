package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"isoft/isoft/common/hashutil"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
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
			this.Data["json"] = &map[string]string{"status": "ERROR", "errorMsg": fmt.Sprint(err)}
			this.ServeJSON()
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
	url := fmt.Sprintf("http://%s/objects/%s", isoft_istorage_web, url.PathEscape(h.Filename))
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
	hash := strings.Replace(strings.TrimSpace(this.GetString("hash"))," ","+",-1)
	url := fmt.Sprintf("http://%s/locate/%s", isoft_istorage_web, hash)
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
	url := fmt.Sprintf("http://%s/objects/%s?version=%s", isoft_istorage_web, name, version)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		panic(err)
	}
	ctype := mime.TypeByExtension(filepath.Ext(name))
	if ctype != ""{
		ctype = "application/octet-stream"
	}
	this.Ctx.ResponseWriter.Header().Set("Content-Type", ctype)
	// Content-Disposition 响应头,设置文件在浏览器打开还是下载
	// Content-Disposition 属性有两种类型:inline和attachment inline:将文件内容直接显示在页面 attachment:弹出对话框让用户下载具体例子
	this.Ctx.ResponseWriter.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", name))
	io.Copy(this.Ctx.ResponseWriter, res.Body)
}

func (this *IFileController) GetImgOrVedioMedia() {
	name := strings.TrimSpace(this.GetString("name", ""))
	version := strings.TrimSpace(this.GetString("version", ""))
	url := fmt.Sprintf("http://%s/objects/%s?version=%s", isoft_istorage_web, name, version)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		panic(err)
	}
	ctype := mime.TypeByExtension(filepath.Ext(name))
	if ctype != ""{
		ctype = "application/octet-stream"
	}
	this.Ctx.ResponseWriter.Header().Set("Content-Type", ctype)
	io.Copy(this.Ctx.ResponseWriter, res.Body)
}

