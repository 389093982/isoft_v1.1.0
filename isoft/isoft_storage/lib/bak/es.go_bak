package es

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/models"
	"net/http"
	"net/url"
	"strings"
)

type hit struct {
	Source models.Metadata `json:"_source"`
}

type searchResult struct {
	Hits struct {
		Total int
		Hits  []hit
	}
}

func getMetadata(name string, versionId int) (meta models.Metadata, e error) {
	url := fmt.Sprintf("http://%s/metadata/objects/%s_%d/_source",
		cfg.GetConfigValue(cfg.ES_SERVER), name, versionId)
	r, e := http.Get(url)
	if e != nil {
		return
	}
	if r.StatusCode != http.StatusOK {
		e = fmt.Errorf("fail to get %s_%d: %d", name, versionId, r.StatusCode)
		return
	}
	result, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(result, &meta)
	return
}

func SearchLatestVersion(name string) (meta models.Metadata, e error) {
	url := fmt.Sprintf("http://%s/metadata/_search?q=name:%s&size=1&sort=version:desc",
		cfg.GetConfigValue(cfg.ES_SERVER), url.PathEscape(name))
	r, e := http.Get(url)
	if e != nil {
		logutil.Errorln(e)
		return
	}
	if r.StatusCode != http.StatusOK {
		logutil.Errorln(r)
		return
	}
	result, _ := ioutil.ReadAll(r.Body)
	var sr searchResult
	json.Unmarshal(result, &sr)
	if len(sr.Hits.Hits) != 0 {
		meta = sr.Hits.Hits[0].Source
	}
	return
}

func GetMetadata(name string, version int) (models.Metadata, error) {
	if version == 0 {
		return SearchLatestVersion(name)
	}
	return getMetadata(name, version)
}

// 四个参数分别表示元数据名称、版本、大小、hash值
func PutMetadata(name string, version int, size int64, hash string) error {
	doc := fmt.Sprintf(`{"name":"%s","version":%d,"size":%d,"hash":"%s"}`,
		name, version, size, hash)
	client := http.Client{}
	url := fmt.Sprintf("http://%s/metadata/objects/%s_%d?op_type=create",
		cfg.GetConfigValue(cfg.ES_SERVER), name, version)
	request, _ := http.NewRequest("PUT", url, strings.NewReader(doc))
	r, e := client.Do(request)
	if e != nil {
		return e
	}
	if r.StatusCode == http.StatusConflict {
		return PutMetadata(name, version+1, size, hash)
	}
	if r.StatusCode != http.StatusCreated {
		result, _ := ioutil.ReadAll(r.Body)
		return fmt.Errorf("fail to put metadata: %d %s", r.StatusCode, string(result))
	}
	return nil
}

func AddVersion(name, hash string, size int64) error {
	version, e := SearchLatestVersion(name)
	if e != nil {
		return e
	}
	return PutMetadata(name, version.Version+1, size, hash)
}

// 根据对象名称 name 模糊匹配统计总数量
func MetadataListCount(name string) (int64, error) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/metadata/_count", cfg.GetConfigValue(cfg.ES_SERVER))
	body := fmt.Sprintf(`
        {
			"query":{
				"wildcard":{"name":"*%s*"}
			}
		}`, name)
	request, _ := http.NewRequest("POST", url, strings.NewReader(body))
	r, e := client.Do(request)
	if e != nil {
		fmt.Println(e)
	}
	result, _ := ioutil.ReadAll(r.Body)
	var countResult map[string]interface{}
	json.Unmarshal(result, &countResult)
	return int64(countResult["count"].(float64)), nil
}

// 根据对象名称 name 模糊匹配
func MetadataList(name string, from, size int) ([]models.Metadata, error) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/metadata/_search", cfg.GetConfigValue(cfg.ES_SERVER))
	body := fmt.Sprintf(`
        {
			"from": %d, "size": %d,
			"sort":[
				{"name":{"order":"asc"}},
				{"version":{"order":"desc"}}
			],
			"query":{
				"wildcard":{"name":"*%s*"}
			}
		}`, from, size, name)
	request, _ := http.NewRequest("POST", url, strings.NewReader(body))
	r, e := client.Do(request)
	if e != nil {
		fmt.Println(e)
	}
	metas := make([]models.Metadata, 0)
	result, _ := ioutil.ReadAll(r.Body)
	var sr searchResult
	json.Unmarshal(result, &sr)
	for i := range sr.Hits.Hits {
		metas = append(metas, sr.Hits.Hits[i].Source)
	}
	return metas, nil
}

// 根据对象名称 name 精确匹配
func SearchAllVersions(name string, from, size int) ([]models.Metadata, error) {
	url := fmt.Sprintf("http://%s/metadata/_search?sort=name,version&from=%d&size=%d",
		cfg.GetConfigValue(cfg.ES_SERVER), from, size)
	if name != "" {
		url += "&q=name:" + name
	}
	r, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	metas := make([]models.Metadata, 0)
	result, _ := ioutil.ReadAll(r.Body)
	var sr searchResult
	json.Unmarshal(result, &sr)
	for i := range sr.Hits.Hits {
		metas = append(metas, sr.Hits.Hits[i].Source)
	}
	return metas, nil
}

func DelMetadata(name string, version int) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/metadata/objects/%s_%d",
		cfg.GetConfigValue(cfg.ES_SERVER), name, version)
	request, _ := http.NewRequest("DELETE", url, nil)
	client.Do(request)
}

type Bucket struct {
	Key         string // 对象的名字
	Doc_count   int    // 对象目前有多少版本
	Min_version struct {
		Value float32 // 对象当前最小的版本号
	}
}

type aggregateResult struct {
	Aggregations struct {
		Group_by_name struct {
			Buckets []Bucket
		}
	}
}

// 返回值 key 为对象名, value 为对象现有版本数量、最小版本信息
func SearchVersionStatus(min_doc_count int) (versionMap map[string][]int, err error) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/metadata/_search", cfg.GetConfigValue(cfg.ES_SERVER))
	body := fmt.Sprintf(`
        {
          "size": 0,
          "aggs": {
            "group_by_name": {
              "terms": {
                "field": "name",
                "min_doc_count": %d
              },
              "aggs": {
                "min_version": {
                  "min": {
                    "field": "version"
                  }
                }
              }
            }
          }
        }`, min_doc_count)
	request, _ := http.NewRequest("GET", url, strings.NewReader(body))
	r, e := client.Do(request)
	if e != nil {
		return nil, e
	}
	b, _ := ioutil.ReadAll(r.Body)
	var ar aggregateResult
	json.Unmarshal(b, &ar)

	buckets := ar.Aggregations.Group_by_name.Buckets
	for i := range buckets {
		bucket := buckets[i]
		versionMap[bucket.Key] = []int{bucket.Doc_count, int(bucket.Min_version.Value)}
	}
	return versionMap, nil
}

func HasHash(hash string) (bool, error) {
	url := fmt.Sprintf("http://%s/metadata/_search?q=hash:%s&size=0", cfg.GetConfigValue(cfg.ES_SERVER), hash)
	r, e := http.Get(url)
	if e != nil {
		return false, e
	}
	b, _ := ioutil.ReadAll(r.Body)
	var sr searchResult
	json.Unmarshal(b, &sr)
	return sr.Hits.Total != 0, nil
}

func SearchHashSize(hash string) (size int64, e error) {
	url := fmt.Sprintf("http://%s/metadata/_search?q=hash:%s&size=1",
		cfg.GetConfigValue(cfg.ES_SERVER), hash)
	r, e := http.Get(url)
	if e != nil {
		return
	}
	if r.StatusCode != http.StatusOK {
		e = fmt.Errorf("fail to search hash size: %d", r.StatusCode)
		return
	}
	result, _ := ioutil.ReadAll(r.Body)
	var sr searchResult
	json.Unmarshal(result, &sr)
	if len(sr.Hits.Hits) != 0 {
		size = sr.Hits.Hits[0].Source.Size
	}
	return
}
