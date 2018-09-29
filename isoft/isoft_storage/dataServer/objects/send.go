package objects

import (
	"compress/gzip"
	"io"
	"isoft/isoft/common/logutil"
	"os"
)

func sendFile(w io.Writer, file string) {
	f, err := os.Open(file)
	if err != nil {
		logutil.Errorln(err)
		return
	}
	defer f.Close()
	// gzip 解压,再读取数据
	gzipStream, err := gzip.NewReader(f)
	if err != nil {
		logutil.Errorln(err)
		return
	}
	io.Copy(w, gzipStream)
	gzipStream.Close()
}
