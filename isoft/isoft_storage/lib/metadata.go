package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/models"
	"isoft/isoft_storage/lib/utils"
	"net/http"
	"strings"
	"time"
)

type MetaDataProxy struct {
	AppName string
}

func convertToMetadata(metadataMap map[string]interface{}) (meta models.Metadata) {
	meta.Name = metadataMap["name"].(string)
	meta.Version = int(metadataMap["version"].(float64))
	meta.Size = int64(metadataMap["size"].(float64))
	meta.Hash = metadataMap["hash"].(string)
	return
}

func (this *MetaDataProxy) SearchLatestVersion(name string) (meta models.Metadata, e error) {
	defer utils.RecordTimeCostForMethod("lib metadata SearchLatestVersion", time.Now())

	url := fmt.Sprintf("http://%s/api/metadata/searchLatestVersion", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader("name="+name))
	if err != nil {
		panic(err)
	}
	responseBody, _ := ioutil.ReadAll(resp.Body)
	responseMap := make(map[string]interface{})
	json.Unmarshal(responseBody, &responseMap)
	if responseMap["status"] != "SUCCESS" {
		logutil.Errorln(errors.New(responseMap["errorMsg"].(string)))
		return meta, errors.New(responseMap["errorMsg"].(string))
	}
	return convertToMetadata(responseMap["metadata"].(map[string]interface{})), nil
}

func (this *MetaDataProxy) GetMetadata(name string, version int) (meta models.Metadata, err error) {
	defer utils.RecordTimeCostForMethod("lib metadata GetMetadata", time.Now())

	url := fmt.Sprintf("http://%s/api/metadata/getMetadata", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("name=%s&version=%d", name, version)))
	if err != nil {
		panic(err)
	}
	responseBody, _ := ioutil.ReadAll(resp.Body)
	responseMap := make(map[string]interface{})
	json.Unmarshal(responseBody, &responseMap)
	if responseMap["status"] != "SUCCESS" {
		logutil.Errorln(errors.New(responseMap["errorMsg"].(string)))
		return meta, errors.New(responseMap["errorMsg"].(string))
	}
	return convertToMetadata(responseMap["metadata"].(map[string]interface{})), nil
}

func (this *MetaDataProxy) PutMetadata(name string, version int, size int64, hash string) error {
	defer utils.RecordTimeCostForMethod("lib metadata PutMetadata", time.Now())

	url := fmt.Sprintf("http://%s/api/metadata/putMetadata", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("name=%s&version=%d&size=%d&hash=%s", name, version, size, hash)))
	if err != nil {
		panic(err)
	}
	responseBody, _ := ioutil.ReadAll(resp.Body)
	responseMap := make(map[string]interface{})
	json.Unmarshal(responseBody, &responseMap)
	if responseMap["status"] != "SUCCESS" {
		logutil.Errorln(errors.New(responseMap["errorMsg"].(string)))
		return errors.New(responseMap["errorMsg"].(string))
	}
	return nil
}

func (this *MetaDataProxy) AddVersion(name, hash string, size int64) error {
	defer utils.RecordTimeCostForMethod("lib metadata AddVersion", time.Now())
	if this.AppName == "" {
		this.AppName = "dataServer"
	}

	url := fmt.Sprintf("http://%s/api/metadata/addVersion", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("name=%s&size=%d&hash=%s&appName=%s", name, size, hash, this.AppName)))
	if err != nil {
		panic(err)
	}
	responseBody, _ := ioutil.ReadAll(resp.Body)
	responseMap := make(map[string]interface{})
	json.Unmarshal(responseBody, &responseMap)
	if responseMap["status"] != "SUCCESS" {
		logutil.Errorln(errors.New(responseMap["errorMsg"].(string)))
		return errors.New(responseMap["errorMsg"].(string))
	}
	return nil
}

func (this *MetaDataProxy) SearchAllVersions(name string, from, size int) (metadatas []models.Metadata, err error) {
	defer utils.RecordTimeCostForMethod("lib metadata SearchAllVersions", time.Now())

	url := fmt.Sprintf("http://%s/api/metadata/addVersion", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("name=%s&from=%d&size=%d", name, from, size)))
	if err != nil {
		panic(err)
	}
	responseBody, _ := ioutil.ReadAll(resp.Body)
	responseMap := make(map[string]interface{})
	json.Unmarshal(responseBody, &responseMap)
	if responseMap["status"] != "SUCCESS" {
		logutil.Errorln(errors.New(responseMap["errorMsg"].(string)))
		return nil, errors.New(responseMap["errorMsg"].(string))
	} else {
		metadataMaps := responseMap["metadatas"].([]interface{})
		for _, metadataMap := range metadataMaps {
			meta := convertToMetadata(metadataMap.(map[string]interface{}))
			metadatas = append(metadatas, meta)
		}
	}
	return
}

func (this *MetaDataProxy) DelMetadata(name string, version int) error {
	defer utils.RecordTimeCostForMethod("lib metadata DelMetadata", time.Now())

	url := fmt.Sprintf("http://%s/api/metadata/delMetadata", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("name=%s&version=%s", name, version)))
	if err != nil {
		panic(err)
	}
	responseBody, _ := ioutil.ReadAll(resp.Body)
	responseMap := make(map[string]interface{})
	json.Unmarshal(responseBody, &responseMap)
	if responseMap["status"] != "SUCCESS" {
		logutil.Errorln(errors.New(responseMap["errorMsg"].(string)))
		return errors.New(responseMap["errorMsg"].(string))
	}
	return nil
}

func (this *MetaDataProxy) HasHash(hash string) (bool, error) {
	defer utils.RecordTimeCostForMethod("lib metadata HasHash", time.Now())

	url := fmt.Sprintf("http://%s/api/metadata/hasHash", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("hash=%s", hash)))
	if err != nil {
		panic(err)
	}
	responseBody, _ := ioutil.ReadAll(resp.Body)
	responseMap := make(map[string]interface{})
	json.Unmarshal(responseBody, &responseMap)
	if responseMap["status"] != "SUCCESS" {
		logutil.Errorln(errors.New(responseMap["errorMsg"].(string)))
		return false, errors.New(responseMap["errorMsg"].(string))
	}
	return true, nil
}

func (this *MetaDataProxy) SearchHashSize(hash string) (size int64, e error) {
	defer utils.RecordTimeCostForMethod("lib metadata SearchHashSize", time.Now())

	url := fmt.Sprintf("http://%s/api/metadata/searchHashSize", cfg.GetConfigValue(cfg.ISOFT_IAAS_WEB))
	resp, err := http.Post(url, "application/x-www-form-urlencoded",
		strings.NewReader(fmt.Sprintf("hash=%s", hash)))
	if err != nil {
		panic(err)
	}
	responseBody, _ := ioutil.ReadAll(resp.Body)
	responseMap := make(map[string]interface{})
	json.Unmarshal(responseBody, &responseMap)
	if responseMap["status"] != "SUCCESS" {
		logutil.Errorln(errors.New(responseMap["errorMsg"].(string)))
		return size, errors.New(responseMap["errorMsg"].(string))
	}
	return responseMap["size"].(int64), nil
}

// 查询所有版本数量大于等于 minVersionCount 的对象
// 返回值 key 为对象名, value 为对象现有版本数量、最小版本信息
func (this *MetaDataProxy) SearchVersionStatus(minVersionCount int) (m map[string][]int, err error) {
	//return es.SearchVersionStatus(minVersionCount)
	return
}
