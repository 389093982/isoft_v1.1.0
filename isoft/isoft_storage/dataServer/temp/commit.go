package temp

import (
	"compress/gzip"
	"io"
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/utils"
	"net/url"
	"os"
	"strconv"
	"strings"
	"isoft/isoft_storage/dataServer/locate"
)

// 获取分片 hash 值
func (t *tempInfo) hash() string {
	s := strings.Split(t.Name, ".")
	return s[0]
}

// 获取分片 id
func (t *tempInfo) id() int {
	s := strings.Split(t.Name, ".")
	id, _ := strconv.Atoi(s[1])
	return id
}

// 将临时对象转正
func commitTempObject(datFile string, tempinfo *tempInfo) {
	f, _ := os.Open(datFile)
	defer f.Close()
	// d 表示当前分片计算出来的 hash 值
	d := url.PathEscape(utils.CalculateHash(f))
	f.Seek(0, io.SeekStart)
	// 正式文件名称
	w, _ := os.Create(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/objects/" + tempinfo.Name + "." + d)
	// 转正默认使用 gzip 进行压缩
	w2 := gzip.NewWriter(w)
	io.Copy(w2, f)
	w2.Close()
	os.Remove(datFile)

	// 将正式文件信息添加到缓存的定位信息中去
	locate.Add(tempinfo.hash(), tempinfo.id())
}
