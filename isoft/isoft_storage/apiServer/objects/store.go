package objects

import (
	"../locate"
	"fmt"
	"io"
	"isoft/isoft_storage/lib/utils"
	"net/http"
	"net/url"
)

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	// 判断对象的 hash 值是否存在
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}

	// 创建存储对象 hash 值得 putStream
	stream, e := putStream(url.PathEscape(hash), size)
	if e != nil {
		return http.StatusInternalServerError, e
	}

	reader := io.TeeReader(r, stream)

	// 重新计算 hash 值
	d := utils.CalculateHash(reader)
	if d != hash {
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	// hash 值正确才可以提交保存
	stream.Commit(true)
	return http.StatusOK, nil
}
