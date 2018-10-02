package download

import (
	"fmt"
	"isoft/isoft_storage/cfg"
	"net/http"
	"path/filepath"
	"strings"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		STORAGE_ROOT := cfg.GetConfigValue(cfg.STORAGE_ROOT)
		hash := strings.Split(r.URL.EscapedPath(), "/")[2]

		fmt.Println(hash)

		w.WriteHeader(http.StatusOK)
		http.ServeFile(w, r, filepath.Join(STORAGE_ROOT, hash))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
