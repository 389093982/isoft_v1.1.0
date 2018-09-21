package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"isoft/isoft/common/hashutil"
	"net/http"
	"net/url"
	"strings"
)

var isoft_istorage_web string

func init()  {
	isoft_istorage_web = beego.AppConfig.String("isoft_istorage_web")
}


type IFileController struct {
	beego.Controller
}

func (this *IFileController) FileUpload() {
	defer func() {
		if err := recover(); err != nil{
			fmt.Println(err)
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
			this.ServeJSON()
			return
		}
	}()
	// 判断是否是文件上传
	_, h, err := this.GetFile("file")
	if err != nil {
		panic(err)
	}
	file, _, err := this.Ctx.Request.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	filecontent := string(bytes)
	// 调用 isoft_istorage_web 发送 put 请求调用分布式对象存储接口,对象名称使用文件名称的 MD5 值
	url := fmt.Sprintf("%s/objects/%s", isoft_istorage_web, url.PathEscape(h.Filename))
	req, err := http.NewRequest("PUT", url, strings.NewReader(filecontent))
	if err != nil {
		panic(err)
	}
	// 在请求头中添加 hash 值
	req.Header.Add("digest", "SHA-256=" + hashutil.CalculateHash(strings.NewReader(filecontent)))
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200{
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "Status":res.Status,"body":body, "filename":h.Filename}
	this.ServeJSON()
}

// 分页查询元数据信息
func (this *IFileController) FilterPageMetadatas()  {
	defer func() {
		if err := recover(); err != nil{
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": err}
			this.ServeJSON()
		}
	}()
	var (
		name string
		current_page int
		offset int
		err error
	)
	name = strings.TrimSpace(this.GetString("name", ""))
	if current_page, err = this.GetInt("current_page",1); err != nil{
		panic(err)
	}
	if offset, err = this.GetInt("offset",1); err != nil{
		panic(err)
	}
	paramMap := make(map[string]interface{})
	paramMap["name"] = name
	paramMap["from"] = (current_page-1) * offset
	paramMap["size"] = offset
	paramByte, err := json.Marshal(paramMap)
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("%s/api/filterPageMetadatas", isoft_istorage_web)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(paramByte))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200{
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}
	var metadatasMap map[string]interface{}
	err = json.Unmarshal(body, &metadatasMap)
	if err != nil{
		panic(err)
	}
	this.Data["json"] = &metadatasMap
	this.ServeJSON()
}