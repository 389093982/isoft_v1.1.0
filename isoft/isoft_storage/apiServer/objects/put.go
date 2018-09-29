package objects

import (
	"fmt"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/lib"
	"isoft/isoft_storage/lib/utils"
	"net/http"
	"strings"
	"time"
)

func put(w http.ResponseWriter, r *http.Request) {
	// 定位对象用对象名,存储对象用 hash 值
	// 从请求头中获取 hash 值
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		logutil.Errorln("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	size := utils.GetSizeFromHeader(r.Header)
	// 存储对象,底层调用数据服务节点的存储功能
	// 不一定能存储成功,比如定位 hash 值已经存储,则无需存储文件,但是元数据版本必须要升版本
	startTime := time.Now()
	c, err := storeObject(r.Body, hash, size)
	fmt.Println("storeObject cost time :", time.Now().Sub(startTime))
	if err != nil {
		logutil.Errorln(err)
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
	proxy := &lib.MetaDataProxy{}
	err = proxy.AddVersion(name, hash, size)
	if err != nil {
		logutil.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
