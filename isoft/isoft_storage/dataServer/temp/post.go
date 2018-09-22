package temp

import (
	"encoding/json"
	"isoft/isoft/common/stringutil"
	"isoft/isoft_storage/cfg"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type tempInfo struct {
	Uuid string
	Name string
	Size int64
}

// 接口服务以 post 方法访问数据服务 temp 接口的 post 方法,用于在数据服务节点上创建一个临时对象
// 本方法不会存储文件内容,只创建临时对象,并返回临时对象关联的 uuid
func post(w http.ResponseWriter, r *http.Request) {
	// 生成 uuid
	uuid := stringutil.RandomUUID()
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	size, e := strconv.ParseInt(r.Header.Get("size"), 0, 64)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 创建临时对象描述信息
	t := tempInfo{uuid, name, size}
	// 将临时对象描述信息写到临时文件中去,保存文件路径为 STORAGE_ROOT/temp/uuid
	e = t.writeToFile()
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// 将生产的 uuid 存放在 STORAGE_ROOT/temp/uuid.dat 临时文件中去
	f, err := os.Create(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/temp/" + t.Uuid + ".dat")
	if err != nil{
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	w.Write([]byte(uuid))
}

// 将临时对象描述信息写到临时文件中去,保存文件路径为 STORAGE_ROOT/temp/uuid
func (t *tempInfo) writeToFile() error {
	f, e := os.Create(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/temp/" + t.Uuid)
	if e != nil {
		return e
	}
	defer f.Close()
	b, _ := json.Marshal(t)
	f.Write(b)
	return nil
}
