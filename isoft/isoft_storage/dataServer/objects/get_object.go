package objects

import (
	"crypto/sha256"
	"encoding/base64"
	"isoft/isoft_storage/cfg"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"isoft/isoft_storage/dataServer/locate"
)

func getFile(name string) string {
	// 找正式文件 (对象hash.分片id.分片hash)
	files, _ := filepath.Glob(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/objects/" + name + ".*")
	// 找的个数不是 1 个
	if len(files) != 1 {
		return ""
	}
	file := files[0]
	h := sha256.New()
	sendFile(h, file)
	// 从文件中读取并计算的 hash
	d := url.PathEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	// 从文件名中截取的 hash
	hash := strings.Split(file, ".")[2]
	if d != hash {
		log.Println("object hash mismatch, remove", file)
		// hash 值不匹配,表示是一个无效文件,需要删除文件和定位信息
		locate.Del(hash)
		os.Remove(file)
		return ""
	}
	// 返回当前文件名
	return file
}
