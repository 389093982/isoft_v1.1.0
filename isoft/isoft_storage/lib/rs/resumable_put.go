package rs

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/lib/objectstream"
	"isoft/isoft_storage/lib/utils"
	"net/http"
)

type resumableToken struct {
	Name    string
	Size    int64
	Hash    string
	Servers []string
	Uuids   []string
}

type RSResumablePutStream struct {
	*RSPutStream
	*resumableToken
}

func NewRSResumablePutStream(dataServers []string, name, hash string, size int64) (*RSResumablePutStream, error) {
	putStream, e := NewRSPutStream(dataServers, hash, size)
	if e != nil {
		return nil, e
	}
	uuids := make([]string, ALL_SHARDS)
	for i := range uuids {
		// 获取各个分片的 uuid
		uuids[i] = putStream.writers[i].(*objectstream.TempPutStream).Uuid
	}
	token := &resumableToken{name, size, hash, dataServers, uuids}
	return &RSResumablePutStream{putStream, token}, nil
}

// 根据 token 中的信息创建 RSResumablePutStream putStream
func NewRSResumablePutStreamFromToken(token string) (*RSResumablePutStream, error) {
	b, e := base64.StdEncoding.DecodeString(token)
	if e != nil {
		return nil, e
	}

	var t resumableToken
	e = json.Unmarshal(b, &t)
	if e != nil {
		return nil, e
	}

	writers := make([]io.Writer, ALL_SHARDS)
	for i := range writers {
		writers[i] = &objectstream.TempPutStream{t.Servers[i], t.Uuids[i]}
	}
	enc := NewEncoder(writers)
	return &RSResumablePutStream{&RSPutStream{enc}, &t}, nil
}

// ToToken 方法将自身数据以 JSON 格式编入，然后返回经过 Base64 编码后的字符串
func (s *RSResumablePutStream) ToToken() string {
	b, _ := json.Marshal(s)
	return base64.StdEncoding.EncodeToString(b)
}

func (s *RSResumablePutStream) CurrentSize() int64 {
	r, err := http.Head(fmt.Sprintf("http://%s/temp/%s", s.Servers[0], s.Uuids[0]))
	if err != nil {
		logutil.Errorln(err)
		return -1
	}
	if r.StatusCode != http.StatusOK {
		logutil.Errorln(r.StatusCode)
		return -1
	}
	size := utils.GetSizeFromHeader(r.Header) * DATA_SHARDS
	if size > s.Size {
		size = s.Size
	}
	return size
}
