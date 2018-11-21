package ifile

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft/common/weedfsutil"
	"isoft/isoft_iaas_web/models/ifile"
	"time"
)

var (
	WEEDFS_URL string
)

func init()  {
	WEEDFS_URL = beego.AppConfig.String("weedfs_url")
}

type IFileController struct {
	beego.Controller
}

func (this *IFileController) FilterPageIFiles()  {
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	search_name := this.GetString("search_name")
	ifiles, count, err := ifile.FilterIFileList(map[string]string{"search_name": search_name}, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)
	//初始化
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "ifiles": &ifiles, "paginator": &paginatorMap}
	}
	this.ServeJSON()
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
	weedFsInfo, err := weedfsutil.SaveFile(WEEDFS_URL, f)
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
