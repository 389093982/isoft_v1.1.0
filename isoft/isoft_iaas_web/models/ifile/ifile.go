package ifile

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type IFile struct {
	Id              int       `json:"id"`
	Fid             string    `json:"fid"`
	FileName        string    `json:"file_name"`
	FileSize        int64     `json:"file_size"`
	Url             string    `json:"url"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func InsertOrUpdateIFile(ifile *IFile) (id int64, err error) {
	o := orm.NewOrm()
	if ifile.Id > 0 {
		id, err = o.Update(ifile)
	} else {
		id, err = o.Insert(ifile)
	}
	return
}

func FilterIFileList(condArr map[string]string, page int, offset int) (ifiles []IFile, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("i_file")
	if search_name, ok := condArr["search_name"]; ok && strings.TrimSpace(search_name) != ""{
		qs = qs.Filter("file_name", search_name)
	}
	qs = qs.OrderBy("-last_updated_time")
	counts, _ = qs.Count()
	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&ifiles)
	return
}