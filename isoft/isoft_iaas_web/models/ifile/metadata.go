package ifile

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type MetaData struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`    // 对象名称
	Version         int       `json:"version"` // 对象版本
	Size            int64     `json:"size"`    // 对象大小
	Hash            string    `json:"hash"`    // 对象 hash 值
	AppName			string	  `json:"app_name"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func SearchLatestVersion(name string) (metaData MetaData, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("meta_data")
	err = qs.Filter("name", name).OrderBy("-version").One(&metaData)
	return
}

func GetMetadata(name string, version int, app_name string) (metaData MetaData, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("meta_data")
	err = qs.Filter("name", name).Filter("version", version).One(&metaData)
	return
}

func PutMetadata(metadata *MetaData) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(metadata)
	return
}

func SearchAllVersions(name string, from, size int64) (metadatas []MetaData, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("meta_data")
	_, err = qs.Filter("name",name).Limit(size, from).All(&metadatas)
	return
}

func DelMetadata(name string, version int) error {
	o := orm.NewOrm()
	qs := o.QueryTable("meta_data")
	_, err := qs.Filter("name", name).Filter("version", version).Delete()
	return err
}

func HasHash(hash string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("meta_data")
	b := qs.Filter("hash", hash).Exist()
	return b
}


func FilterPageMetadatas(condArr map[string]interface{}, page int, offset int) (metaDatas []MetaData, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("meta_data")

	if name, ok := condArr["name"]; ok {
		qs.Filter("name__contains", name)
	}
	counts, _ = qs.Count()
	_, err = qs.OrderBy("-last_updated_time","-version").Limit(offset, (page-1)*offset).All(&metaDatas)
	return
}

func SearchHashSize()  {

}

