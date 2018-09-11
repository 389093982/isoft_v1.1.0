package objects

import (
	"isoft/isoft_storage/cfg"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"isoft/isoft_storage/dataServer/locate"
)

func del(w http.ResponseWriter, r *http.Request) {
	hash := strings.Split(r.URL.EscapedPath(), "/")[2]
	files, _ := filepath.Glob(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/objects/" + hash + ".*")
	if len(files) != 1 {
		return
	}
	// 将该散列值移出对象定位缓存
	locate.Del(hash)
	// 将对象文件移动到垃圾文件目录中
	os.Rename(files[0], cfg.GetConfigValue(cfg.STORAGE_ROOT)+"/garbage/"+filepath.Base(files[0]))
}
