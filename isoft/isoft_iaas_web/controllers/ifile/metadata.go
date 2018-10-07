package ifile

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common/pageutil"
	"isoft/isoft_iaas_web/models/ifile"
	"strings"
	"time"
)

type MetadataController struct {
	beego.Controller
}

func (this *MetadataController) SearchLatestVersion() {
	name := this.GetString("name")
	metadata, err := ifile.SearchLatestVersion(name)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "metadata": metadata}
	}
	this.ServeJSON()
}

func (this *MetadataController) GetMetadata() {
	name := this.GetString("name")
	version, _ := this.GetInt("version", -1)
	app_name := this.GetString("app_name")
	metadata, err := ifile.GetMetadata(name, version, app_name)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "metadata": metadata}
	}
	this.ServeJSON()
}

func (this *MetadataController) PutMetadata() {
	name := this.GetString("name")
	version, _ := this.GetInt("version", -1)
	size, _ := this.GetInt64("size", -1)
	hash := strings.Replace(strings.TrimSpace(this.GetString("hash")), " ", "+", -1)
	metadata := &ifile.MetaData{
		Name:            name,
		Version:         version,
		Size:            size,
		Hash:            hash,
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}
	err := ifile.PutMetadata(metadata)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *MetadataController) AddVersion() {
	name := this.GetString("name")
	size, _ := this.GetInt64("size", -1)
	hash := strings.Replace(strings.TrimSpace(this.GetString("hash")), " ", "+", -1)
	appName := this.GetString("appName")
	metadata := &ifile.MetaData{
		Name:            name,
		Size:            size,
		Hash:            hash,
		AppName:         appName,
		CreatedBy:       "AutoInsert",
		CreatedTime:     time.Now(),
		LastUpdatedBy:   "AutoInsert",
		LastUpdatedTime: time.Now(),
	}
	oldmetadata, err := ifile.SearchLatestVersion(name)
	if err == nil {
		metadata.Version = oldmetadata.Version + 1
	} else {
		metadata.Version = 1
	}
	err = ifile.PutMetadata(metadata)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *MetadataController) SearchAllVersions() {
	name := this.GetString("name")
	from, _ := this.GetInt64("from", 0)
	size, _ := this.GetInt64("size", 10)
	metadatas, err := ifile.SearchAllVersions(name, from, size)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "metadatas": metadatas}
	}
	this.ServeJSON()
}

func (this *MetadataController) DelMetadata() {
	name := this.GetString("name")
	version, _ := this.GetInt("version", -1)
	err := ifile.DelMetadata(name, version)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": err.Error()}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

func (this *MetadataController) HasHash() {
	hash := strings.Replace(strings.TrimSpace(this.GetString("hash")), " ", "+", -1)
	b := ifile.HasHash(hash)
	if !b {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "hash was not found!"}
	} else {
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
	}
	this.ServeJSON()
}

// 分页查询元数据信息
func (this *MetadataController) FilterPageMetadatas() {
	defer func() {
		if err := recover(); err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "msg": err}
			this.ServeJSON()
		}
	}()
	var (
		name         string
		current_page int
		offset       int
		err          error
	)
	name = strings.TrimSpace(this.GetString("name", ""))
	if current_page, err = this.GetInt("current_page", 1); err != nil {
		panic(err)
	}
	if offset, err = this.GetInt("offset", 1); err != nil {
		panic(err)
	}
	metadatas, count, err := ifile.FilterPageMetadatas(map[string]interface{}{"name": name}, current_page, offset)
	if err != nil {
		this.Data["json"] = &map[string]interface{}{"status": "ERROR"}
	} else {
		paginator := pagination.SetPaginator(this.Ctx, offset, count)
		paginatorMap := pageutil.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
		this.Data["json"] = &map[string]interface{}{"status": "SUCCESS", "metadatas": metadatas, "paginator": paginatorMap}
	}
	this.ServeJSON()
}

func (this *MetadataController) SearchHashSize() {

}

func (this *MetadataController) SearchVersionStatus() {

}
