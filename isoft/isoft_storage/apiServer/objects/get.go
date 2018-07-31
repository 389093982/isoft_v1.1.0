package objects

import (
	"compress/gzip"
	"fmt"
	"io"
	"isoft/isoft_storage/lib/es"
	"isoft/isoft_storage/lib/utils"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// 默认使用最新的版本,可以不传 ?version=xx,版本号从 1 开始递增
func get(w http.ResponseWriter, r *http.Request) {
	// 从请求路径中截取对象名称
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 从请求参数中获取对象版本
	versionId := r.URL.Query()["version"]
	version := 0 // 0 代表最新版本
	var e error
	if len(versionId) != 0 {
		version, e = strconv.Atoi(versionId[0])
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	// 获取对象元数据信息
	meta, e := es.GetMetadata(name, version)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// hash 值为空字符串说明该对象该版本是一个删除标记
	if meta.Hash == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	hash := url.PathEscape(meta.Hash)

	// 根据对象 hash 值获取对象的 getStream
	stream, e := GetStream(hash, meta.Size)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 从 HTTP 请求的 Range 头部获得客户端要求的偏移量
	offset := utils.GetOffsetFromHeader(r.Header)
	if offset != 0 {
		// 跳到 offset 位置进行下载,两个参数分别表示需要跳过的字节数和起跳点
		stream.Seek(offset, io.SeekCurrent)
		w.Header().Set("content-range", fmt.Sprintf("bytes %d-%d/%d", offset, meta.Size-1, meta.Size))
		w.WriteHeader(http.StatusPartialContent)
	}
	acceptGzip := false
	encoding := r.Header["Accept-Encoding"]
	for i := range encoding {
		if encoding[i] == "gzip" {
			acceptGzip = true
			break
		}
	}
	if acceptGzip {
		// 下载可支持 gzip 格式下载
		w.Header().Set("content-encoding", "gzip")
		w2 := gzip.NewWriter(w)
		io.Copy(w2, stream)
		w2.Close()
	} else {
		io.Copy(w, stream)
	}
	stream.Close()
}
