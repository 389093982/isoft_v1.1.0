package objects

import (
	"isoft/isoft_storage/lib/es"
	"isoft/isoft_storage/lib/utils"
	"log"
	"net/http"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	// 定位对象用对象名,存储对象用 hash 值
	// 从请求头中获取 hash 值
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	size := utils.GetSizeFromHeader(r.Header)

	// 存储对象,底层调用数据服务节点的存储功能
	c, e := storeObject(r.Body, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(c)
		return
	}
	if c != http.StatusOK {
		w.WriteHeader(c)
		return
	}

	// 获取对象名称
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 添加对象版本信息
	e = es.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
