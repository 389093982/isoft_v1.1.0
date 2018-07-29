package temp

import (
	"isoft/isoft_storage/cfg"
	"net/http"
	"os"
	"strings"
)

// 删除临时文件 uuid 和 uuid.dat 文件
func del(w http.ResponseWriter, r *http.Request) {
	uuid := strings.Split(r.URL.EscapedPath(), "/")[2]
	infoFile := cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/temp/" + uuid
	datFile := infoFile + ".dat"
	os.Remove(infoFile)
	os.Remove(datFile)
}
