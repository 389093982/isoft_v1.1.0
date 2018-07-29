package temp

import (
	"isoft/isoft_storage/cfg"
	"log"
	"net/http"
	"os"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	uuid := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 读取 STORAGE_ROOT/temp/uuid 文件获取临时对象信息
	tempinfo, e := readFromFile(uuid)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// STORAGE_ROOT/temp/uuid.dat 文件,含上传对象内容的临时文件
	infoFile := cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/temp/" + uuid
	datFile := infoFile + ".dat"
	f, e := os.Open(datFile)
	defer f.Close()
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	info, e := f.Stat()
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	actual := info.Size()
	// 删除 STORAGE_ROOT/temp/uuid 文件,不在需要临时对象描述文件
	os.Remove(infoFile)
	// 判断文件大小是否正确
	if actual != tempinfo.Size {
		os.Remove(datFile)
		log.Println("actual size mismatch, expect", tempinfo.Size, "actual", actual)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 临时文件转正
	commitTempObject(datFile, tempinfo)
}
