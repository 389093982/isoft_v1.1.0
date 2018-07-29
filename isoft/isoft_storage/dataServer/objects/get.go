package objects

import (
	"net/http"
	"strings"
)

func get(w http.ResponseWriter, r *http.Request) {
	// 从请求 URL 中获取对象的散列值,根据散列值获取对象的文件名 file
	file := getFile(strings.Split(r.URL.EscapedPath(), "/")[2])
	if file == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 将文件内容写入 writer
	sendFile(w, file)
}
