package versions

import (
	"encoding/json"
	"isoft/isoft_storage/lib/es"
	"log"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// from 和 size 用于分页显示, ES 默认的分页是从 0 开始显示 10 条
	from := 0
	size := 1000
	// 对象名称,如果请求地址是 /versions/ 则对象名称 name 就是空字符串
	name := strings.Split(r.URL.EscapedPath(), "/")[2]
	for {
		// 查询元数据
		metas, e := es.SearchAllVersions(name, from, size)
		if e != nil {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// 遍历元数据数组,将元数据一一写入 HTTP 响应的正文
		for i := range metas {
			b, _ := json.Marshal(metas[i])
			w.Write(b)
			w.Write([]byte("\n"))
		}
		// 本次迭代查询的元数据数据不等于 size,表示元数据服务中没有更多的数据,直接返回结束
		if len(metas) != size {
			return
		}
		// from = 1000,size = 1000,进行下一次迭代,查询下一批元数据
		from += size
	}
}
