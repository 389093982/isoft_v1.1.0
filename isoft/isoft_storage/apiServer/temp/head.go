package temp

import (
	"fmt"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/lib/rs"
	"net/http"
	"strings"
)

func head(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 根据 token 恢复出 stream
	stream, err := rs.NewRSResumablePutStreamFromToken(token)
	if err != nil {
		logutil.Errorln(err)
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
