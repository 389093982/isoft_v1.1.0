package locate

import (
	"encoding/json"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/models"
	"net/http"
	"strings"
)

func post(w http.ResponseWriter, r *http.Request) {
	hash := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 定位 hash 值是否存在
	id := Locate(hash)
	if id == -1 {
		// 定位信息不存在则直接报错
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		locateInfo := models.LocateMessage{Addr: cfg.GetConfigValue(cfg.LISTEN_ADDRESS), ShardId: id}
		js, err := json.Marshal(locateInfo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(js)
		}
	}
	return
}
