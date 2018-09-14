package objects

import (
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/apiServer/heartbeat"
	"isoft/isoft_storage/apiServer/locate"
	"isoft/isoft_storage/lib/es"
	"isoft/isoft_storage/lib/rs"
	"isoft/isoft_storage/lib/utils"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// post 函数和 put 函数处理流程在前半段是一样的,都是从请求 URL 中获取对象名称,从请求的响应头部获取对象的大小和散列值,
// 然后对散列值进行定位
func post(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	size, e := strconv.ParseInt(r.Header.Get("size"), 0, 64)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	hash := utils.GetHashFromHeader(r.Header)
	if hash == "" {
		log.Println("missing object hash in digest header")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if locate.Exist(url.PathEscape(hash)) {
		// 散列值已经存在,则不做任何保存操作,直接添加版本
		e = es.AddVersion(name, hash, size)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		return
	}
	ds := heartbeat.ChooseRandomDataServers(rs.ALL_SHARDS, nil)
	if len(ds) != rs.ALL_SHARDS {
		logutil.Errorln("cannot find enough dataServer:", len(ds))
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	// 生成对象的 putStream
	stream, e := rs.NewRSResumablePutStream(ds, name, url.PathEscape(hash), size)
	if e != nil {
		logutil.Errorln(e)
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 上传文件之后返回 toToken 信息
	w.Header().Set("location", "/temp/"+url.PathEscape(stream.ToToken()))
	w.WriteHeader(http.StatusCreated)
}
