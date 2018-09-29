package temp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/cfg"
	"net/http"
	"os"
	"strings"
)

func patch(w http.ResponseWriter, r *http.Request) {
	uuid := strings.Split(r.URL.EscapedPath(), "/")[2]
	// 读取 STORAGE_ROOT/temp/uuid 文件获取临时对象信息
	tempinfo, err := readFromFile(uuid)
	if err != nil {
		logutil.Errorln(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	infoFile := cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/temp/" + uuid
	datFile := infoFile + ".dat"
	f, err := os.OpenFile(datFile, os.O_WRONLY|os.O_APPEND, 0)
	defer f.Close()
	if err != nil {
		logutil.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 将上传文件内容写入 STORAGE_ROOT/temp/uuid.dat 文件中去
	_, err = io.Copy(f, r.Body)
	if err != nil {
		logutil.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 判断 STORAGE_ROOT/temp/uuid.dat 文件是否存在
	info, err := f.Stat()
	if err != nil {
		logutil.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 判断 STORAGE_ROOT/temp/uuid.dat 文件大小是否正确
	actual := info.Size()
	if actual > tempinfo.Size {
		os.Remove(datFile)
		os.Remove(infoFile)
		logutil.Errorln("actual size", actual, "exceeds", tempinfo.Size)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func readFromFile(uuid string) (*tempInfo, error) {
	// 读取 STORAGE_ROOT/temp/uuid 文件获取临时对象信息
	f, err := os.Open(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/temp/" + uuid)
	if err != nil {
		logutil.Errorln(err)
		return nil, err
	}
	defer f.Close()
	b, _ := ioutil.ReadAll(f)
	var info tempInfo
	json.Unmarshal(b, &info)
	return &info, nil
}
