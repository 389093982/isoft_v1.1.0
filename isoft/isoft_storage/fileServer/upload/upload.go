package upload

import (
	"io"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib"
	"isoft/isoft_storage/lib/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		STORAGE_ROOT := cfg.GetConfigValue(cfg.STORAGE_ROOT)

		// 获取对象名称
		name := strings.Split(r.URL.EscapedPath(), "/")[2]
		// 特殊字符进行转义
		hash := strings.Split(r.URL.EscapedPath(), "/")[3]
		size := utils.GetSizeFromHeader(r.Header)
		hash = strings.Replace(strings.TrimSpace(hash), " ", "+", -1)

		dst, err := os.Create(filepath.Join(STORAGE_ROOT, hash))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		if _, err := io.Copy(dst, r.Body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 添加对象版本信息
		proxy := &lib.MetaDataProxy{AppName: "fileServer"}
		err = proxy.AddVersion(name, hash, size)
		if err != nil {
			logutil.Errorln(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
