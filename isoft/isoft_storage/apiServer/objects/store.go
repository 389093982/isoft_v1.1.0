package objects

import (
	"fmt"
	"io"
	"isoft/isoft_storage/apiServer/locate"
	"isoft/isoft_storage/lib/utils"
	"net/http"
	"net/url"
)

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	// 判断对象的 hash 值是否存在
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}

	// 创建存储对象 hash 对应的 putStream,生成临时文件对象
	stream, e := putStream(url.PathEscape(hash), size)
	if e != nil {
		return http.StatusInternalServerError, e
	}

	reader := io.TeeReader(r, stream)

	// 重新计算 hash 值
	d := utils.CalculateHash(reader)
	if d != hash {
		// commit = false,底层调用数据服务的 temp 接口的 delete 方法删除临时文件
		// 删除临时文件 uuid 和 uuid.dat 文件
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	// hash 值正确才可以提交保存
	// stream 的 Commit 方法先调用数据服务 temp 接口的 patch 方法将请求的正文写入该临时对象,并将临时对象转正
	stream.Commit(true)
	return http.StatusOK, nil
}
