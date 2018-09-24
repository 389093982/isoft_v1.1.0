package objects

import "net/http"

// 不直接访问本地磁盘上的对象,而是将 http 请求转发给数据服务节点
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	// post 方法暂不支持,不支持断点续传
	//if m == http.MethodPost {
	//	// post 方法主要用于创建 token, token 可以还原出 stream
	//	// stream, e := rs.NewRSResumablePutStreamFromToken(token)
	//	//post(w, r)
	//	//return
	//}
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
	if m == http.MethodDelete {
		del(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
