package temp

import (
	"fmt"
	"isoft/isoft_storage/lib/rs"
	"log"
	"net/http"
	"strings"
)

func head(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 根据 token 恢复出 stream
	stream, e := rs.NewRSResumablePutStreamFromToken(token)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 获取对象当前大小
	current := stream.CurrentSize()
	if current == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("content-length", fmt.Sprintf("%d", current))
}
