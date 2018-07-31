package temp

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodHead {
		// head 方法用于获取当前上传字节数
		head(w, r)
		return
	}
	if m == http.MethodPut {
		// 实际的上传方法,支持断点上传
		put(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
