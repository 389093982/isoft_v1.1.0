package api

import (
	"encoding/json"
	"io/ioutil"
	"isoft/isoft_storage/lib"
	"net/http"
)

func FilterPageMetadatas(w http.ResponseWriter, r *http.Request) {
	var (
		result []byte
		err    error
	)
	body, _ := ioutil.ReadAll(r.Body)
	type Param struct {
		Name string `json:"name"`
		From int    `json:"from"`
		Size int    `json:"size"`
	}
	var p Param
	json.Unmarshal(body, &p)
	// 非法参数则使用默认值
	if p.From < 0 {
		p.From = 0
	}
	if p.Size <= 0 {
		p.Size = 10
	}
	proxy := lib.MetaDataProxy{}
	metadatas, paginator, err := proxy.FilterPageMetadatas(p.Name, p.From, p.Size)
	if err != nil {
		result, _ = json.Marshal(map[string]interface{}{"status": "ERROR", "msg": "SearchAllVersions error"})
	} else {
		result, err = json.Marshal(map[string]interface{}{"status": "SUCCESS", "metadatas": metadatas, "paginator": paginator})
		if err != nil {
			result, err = json.Marshal(map[string]interface{}{"status": "ERROR", "msg": "SearchAllVersions error"})
		}
	}
	w.Header().Set("Content-Type", "application/json;")
	w.Write(result)
}
