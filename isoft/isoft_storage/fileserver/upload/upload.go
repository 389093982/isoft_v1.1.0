package upload

import (
	"io"
	"isoft/isoft_storage/fileserver/cfg"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		hash := strings.Split(r.URL.EscapedPath(), "/")[2]
		STORAGE_ROOT := cfg.GetConfigValue(cfg.STORAGE_ROOT)
		file, err := os.Create(filepath.Join(STORAGE_ROOT, hash))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		_, err = io.Copy(file, r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
