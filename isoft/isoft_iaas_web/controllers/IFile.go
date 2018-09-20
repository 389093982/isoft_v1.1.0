package controllers

import (
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
