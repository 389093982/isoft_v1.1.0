package lib

import (
	"isoft/isoft_storage/lib/es"
	"isoft/isoft_storage/lib/models"
)


type MetaDataProxy struct {

}

func (this *MetaDataProxy) SearchLatestVersion(name string) (meta models.Metadata, e error) {
	return es.SearchLatestVersion(name)
}

func (this *MetaDataProxy) GetMetadata(name string, version int) (models.Metadata, error) {
	return es.GetMetadata(name, version)
}

func (this *MetaDataProxy) PutMetadata(name string, version int, size int64, hash string) error {
	return es.PutMetadata(name, version, size, hash)
}

func (this *MetaDataProxy) AddVersion(name, hash string, size int64) error {
	return es.AddVersion(name, hash, size)
}

func (this *MetaDataProxy) SearchAllVersions(name string, from, size int) ([]models.Metadata, error) {
	return es.SearchAllVersions(name, from, size)
}

func (this *MetaDataProxy) DelMetadata(name string, version int) {
	es.DelMetadata(name, version)
}

func (this *MetaDataProxy) HasHash(hash string) (bool, error) {
	return es.HasHash(hash)
}

func (this *MetaDataProxy) SearchHashSize(hash string) (size int64, e error) {
	return es.SearchHashSize(hash)
}

// 查询所有版本数量大于等于 minVersionCount 的对象
// 返回值 key 为对象名, value 为对象现有版本数量、最小版本信息
func (this *MetaDataProxy) SearchVersionStatus(minVersionCount int) (map[string][]int, error) {
	return es.SearchVersionStatus(minVersionCount)
}