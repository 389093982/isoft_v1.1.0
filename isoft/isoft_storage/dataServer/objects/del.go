package objects

import (
	"../locate"
	"isoft/isoft_storage/cfg"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func del(w http.ResponseWriter, r *http.Request) {
	hash := strings.Split(r.URL.EscapedPath(), "/")[2]
	files, _ := filepath.Glob(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/objects/" + hash + ".*")
	if len(files) != 1 {
		return
	}
	locate.Del(hash)
	os.Rename(files[0], cfg.GetConfigValue(cfg.STORAGE_ROOT)+"/garbage/"+filepath.Base(files[0]))
}
