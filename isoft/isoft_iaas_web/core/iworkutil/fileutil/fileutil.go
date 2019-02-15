package fileutil

import (
	"io"
	"io/ioutil"
	"os"
)

// 写文件,文件不存在时会自动创建
func WriteFile(filename string, data []byte, append bool) error {
	if append {
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		n, err := f.Write(data)
		if err == nil && n < len(data) {
			err = io.ErrShortWrite
		}
		if err1 := f.Close(); err == nil {
			err = err1
		}
		return err
	}else{
		return ioutil.WriteFile(filename, data, 0666)
	}
}