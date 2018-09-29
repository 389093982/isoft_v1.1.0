package objects

import (
	"fmt"
	"io"
	"isoft/isoft/common/hashutil"
	"isoft/isoft_storage/apiServer/locate"
	"net/http"
	"net/url"
	"time"
)

func storeObject(r io.Reader, hash string, size int64) (int, error) {
	startTime := time.Now()
	// 判断对象的 hash 值是否存在
	if locate.Exist(url.PathEscape(hash)) {
		return http.StatusOK, nil
	}
	endTime := time.Now()
	fmt.Println("storeObject locate.Exist cost time :", endTime.Sub(startTime))

	startTime = time.Now()
	// 创建存储对象 hash 对应的 putStream,生成临时文件对象
	stream, err := putStream(url.PathEscape(hash), size)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	endTime = time.Now()
	fmt.Println("storeObject putStream cost time :", endTime.Sub(startTime))

	startTime = time.Now()
	reader := io.TeeReader(r, stream)
	endTime = time.Now()
	fmt.Println("storeObject io.TeeReader cost time :", endTime.Sub(startTime))

	startTime = time.Now()
	// 重新计算 hash 值
	d := hashutil.CalculateHash(reader)
	endTime = time.Now()
	fmt.Println("storeObject hashutil.CalculateHash cost time :", endTime.Sub(startTime))

	if d != hash {
		// commit = false,底层调用数据服务的 temp 接口的 delete 方法删除临时文件
		// 删除临时文件 uuid 和 uuid.dat 文件
		stream.Commit(false)
		return http.StatusBadRequest, fmt.Errorf("object hash mismatch, calculated=%s, requested=%s", d, hash)
	}
	// hash 值正确才可以提交保存
	// stream 的 Commit 方法先调用数据服务 temp 接口的 patch 方法将请求的正文写入该临时对象,并将临时对象转正
	startTime = time.Now()
	stream.Commit(true)
	endTime = time.Now()
	fmt.Println("storeObject stream.Commit cost time :", endTime.Sub(startTime))
	return http.StatusOK, nil
}
