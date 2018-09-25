package locate

import (
	"net/http"
)

// 数据服务读取对象 get 方法
// 数据服务写入对象 put 方法由数据服务 temp 接口的 post、patch、put 替代
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPost {
		post(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
