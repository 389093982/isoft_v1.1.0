package objects

import (
	"bytes"
	"io"
	"isoft/isoft/common/hashutil"
	"isoft/isoft_storage/lib"
	"isoft/isoft_storage/lib/utils"
	"log"
	"net/http"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	bReader := bytes.Buffer{}
	// io.TeeReader、io.MultiReader
	reader := io.TeeReader(r.Body, &bReader)
	// 定位对象用对象名,存储对象用 hash 值
	hash := hashutil.CalculateHash(reader)

	// 定位对象用对象名,存储对象用 hash 值
	// 从请求头中获取 hash 值
	//hash := utils.GetHashFromHeader(r.Header)
	//if hash == "" {
	//	log.Println("missing object hash in digest header")
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}

	size := utils.GetSizeFromHeader(r.Header)

	// 存储对象,底层调用数据服务节点的存储功能
	// 不一定能存储成功,比如定位 hash 值已经存储,则无需存储文件,但是元数据版本必须要升版本
	c, e := storeObject(&bReader, hash, size)
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
	proxy := &lib.MetaDataProxy{}
	e = proxy.AddVersion(name, hash, size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
