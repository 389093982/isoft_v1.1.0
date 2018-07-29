package objects

import (
	"isoft/isoft_storage/lib/es"
	"log"
	"net/http"
	"strings"
)

// 删除一个对象
func del(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	version, e := es.SearchLatestVersion(name)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 在元数据中给对象添加一个表示删除的特殊版本,而在数据节点上保留其数据 (hash置空)
	e = es.PutMetadata(name, version.Version+1, 0, "")
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
