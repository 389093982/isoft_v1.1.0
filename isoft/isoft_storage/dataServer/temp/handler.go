package temp

import (
	"net/http"
)

// STORAGE_ROOT/temp/uuid 用于记录临时对象的 uuid、名字和大小
// STORAGE_ROOT/temp/uuid.dat 用于保存该临时对象的内容
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodPatch {
		patch(w, r)
		return
	}
	if m == http.MethodPost {
		post(w, r)
		return
	}
	if m == http.MethodDelete {
		del(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
