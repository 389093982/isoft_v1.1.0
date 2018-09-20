package api

import (
	"net/http"
	"strings"
)

// 管理分布式对象存储系统统一接口
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// 获取操作名称
	methodName := strings.Split(r.URL.EscapedPath(), "/")[2]
	switch strings.TrimSpace(methodName) {
	case "searchAllVersions":
		SearchAllVersions(w, r)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}