package objects

import (
	"../locate"
	"crypto/sha256"
	"encoding/base64"
	"isoft/isoft_storage/cfg"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func getFile(name string) string {
	files, _ := filepath.Glob(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/objects/" + name + ".*")
	if len(files) != 1 {
		return ""
	}
	file := files[0]
	h := sha256.New()
	sendFile(h, file)
	d := url.PathEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	hash := strings.Split(file, ".")[2]
	// 验证 hash 值
	if d != hash {
		log.Println("object hash mismatch, remove", file)
		locate.Del(hash)
		os.Remove(file)
		return ""
	}
	return file
}
