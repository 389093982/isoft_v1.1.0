package locate

import (
	"encoding/json"
	"net/http"
	"strings"
)

// 接口服务对外提供定位功能,向数据服务节点群发定位消息并接收反馈消息
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// strings.Split(r.URL.EscapedPath(), "/")[2] 从请求地址中获取定位对象名称
	// 并向数据服务节点群发对象名字的定位消息,并接收反馈消息
	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, _ := json.Marshal(info)
	w.Write(b)
}
