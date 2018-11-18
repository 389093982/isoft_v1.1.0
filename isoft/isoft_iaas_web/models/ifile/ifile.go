package ifile

import (
	"github.com/astaxie/beego/orm"
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
