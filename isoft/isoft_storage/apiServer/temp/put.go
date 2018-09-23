package temp

import (
	"io"
	"isoft/isoft/common/hashutil"
	"isoft/isoft_storage/apiServer/locate"
	"isoft/isoft_storage/lib"
	"isoft/isoft_storage/lib/rs"
	"isoft/isoft_storage/lib/utils"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func put(w http.ResponseWriter, r *http.Request) {
	// 从请求 url 地址中获取 token
	token := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := rs.NewRSResumablePutStreamFromToken(token)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// 计算已经上传的数据量
	current := stream.CurrentSize()
	if current == -1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 从请求头中获取偏移量
	offset := utils.GetOffsetFromHeader(r.Header)
	// 已经上传数据量和偏移量不一致则报错
	if current != offset {
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
		return
	}
	// 一次 putStream 写入的数据量, 32000 字节
	bytes := make([]byte, rs.BLOCK_SIZE)
	for {
		// 多次写入,每次写入 32000 字节,直到写完为止
		n, e := io.ReadFull(r.Body, bytes) // n 表示一次读取的实际读取量
		if e != nil && e != io.EOF && e != io.ErrUnexpectedEOF {
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		current += int64(n)
		if current > stream.Size { // 读取错误
			stream.Commit(false)
			log.Println("resumable put exceed size")
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if n != rs.BLOCK_SIZE && current != stream.Size {
			return
		}
		// 写入本次循环实际读取量 n
		stream.Write(bytes[:n])
		// 读到的总长度等于对象的大小,说明客户端上传了对象的全部数据
		if current == stream.Size {
			// 立即写入临时对象
			stream.Flush()
			// 获取临时对象 getStram
			getStream, e := rs.NewRSResumableGetStream(stream.Servers, stream.Uuids, stream.Size)
			// 计算 hash
			hash := url.PathEscape(hashutil.CalculateHash(getStream))
			// hash 不一致,上传有误
			if hash != stream.Hash {
				stream.Commit(false)
				log.Println("resumable put done but hash mismatch")
				w.WriteHeader(http.StatusForbidden)
				return
			}

			if locate.Exist(url.PathEscape(hash)) {
				// hash 值已经存在,则删除临时对象
				stream.Commit(false)
			} else {
				// 将临时对象转正
				stream.Commit(true)
			}
			// 记录版本
			proxy := &lib.MetaDataProxy{}
			e = proxy.AddVersion(stream.Name, stream.Hash, stream.Size)
			if e != nil {
				log.Println(e)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}
}
