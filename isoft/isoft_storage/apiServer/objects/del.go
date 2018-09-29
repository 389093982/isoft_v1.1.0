package objects

import (
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/lib"
	"net/http"
	"strings"
)

// 删除一个对象
func del(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	proxy := &lib.MetaDataProxy{}
	version, err := proxy.SearchLatestVersion(name)
	if err != nil {
		logutil.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 在元数据中给对象添加一个表示删除的特殊版本,而在数据节点上保留其数据 (size为0,hash为空字符串,表示这是一个删除标记)
	err = proxy.PutMetadata(name, version.Version+1, 0, "")
	if err != nil {
		logutil.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
