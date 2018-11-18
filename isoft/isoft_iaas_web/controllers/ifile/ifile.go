package ifile

import (
	"fmt"
	"github.com/astaxie/beego"
	"isoft/isoft/common/weedfsutil"
	"isoft/isoft_iaas_web/models/ifile"
	"time"
)

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
	weedFsInfo, err := weedfsutil.SaveFile("193.112.162.61:9333", f)
	if err != nil {
		panic(err)
	}
	ff := &ifile.IFile{
		Fid:             weedFsInfo.Fid,
		FileName:        h.Filename,
		FileSize:        h.Size,
		Url:             fmt.Sprintf("http://%s/%s", weedFsInfo.PublicUrl, weedFsInfo.Fid),
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}
	_, err = ifile.InsertOrUpdateIFile(ff)
	if err != nil {
		panic(err)
	}
	this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "Status": 200, "filename": h.Filename}
	this.ServeJSON()
}
