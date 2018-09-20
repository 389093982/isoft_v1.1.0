package api

import (
	"encoding/json"
	"isoft/isoft_storage/lib"
	"net/http"
	"strconv"
)

func SearchAllVersions(w http.ResponseWriter, r *http.Request)  {
	var (
		result []byte
		err error
	)
	name := r.PostFormValue("name")
	from, err := strconv.Atoi(r.PostFormValue("from"))
	// 非法参数则使用默认值
	if err != nil{
		from = 0
	}
	size, err := strconv.Atoi(r.PostFormValue("size"))
	if err != nil{
		size = 10
	}
	proxy := lib.MetaDataProxy{}
	metadatas, err := proxy.SearchAllVersions(name, from, size)
	if err != nil{
		result, _ = json.Marshal(map[string]interface{}{"status":"ERROR","msg":"SearchAllVersions error"})
	}else{
		result, err = json.Marshal(map[string]interface{}{"status":"SUCCESS","data":metadatas})
		if err != nil{
			result, err = json.Marshal(map[string]interface{}{"status":"ERROR","msg":"SearchAllVersions error"})
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
