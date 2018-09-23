package hashutil

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func CalculateHash(r io.Reader) string {
	h := sha256.New()
	_, err := io.Copy(h, r)
	if err != nil{
		fmt.Println("CalculateHash err")
		return ""
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func CalculateHashWithString(msg string) string {
	h := sha256.New()
	reader := strings.NewReader(msg)
	_, err := io.Copy(h, reader)
	if err != nil{
		fmt.Println("CalculateHashWithString err")
		return ""
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func CalculateHashWithFileS(filepath string) (hash string, err error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil{
		return "", err
	}
	return CalculateHashWithString(string(bytes)), nil
}

func CalculateHashWithFile(filepath string) (hash string, err error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return "",err
	}
	h := sha256.New()
	_, err = io.Copy(h,file)
	if err != nil{
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

func CalculateHashWithBigFile(filepath string) (hash string, err error){
	const filechunk = 1024 * 1024*10 // 文件块大小设置为 10M 大小可以调整
	file, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	// calculate the file size
	info, err := file.Stat()
	if err != nil{
		return "", err
	}
	filesize := info.Size()
	blocks := uint64(math.Ceil(float64(filesize) / float64(filechunk)))
	h := sha256.New()
	for i := uint64(0); i < blocks; i++ {
		blocksize := int(math.Min(filechunk, float64(filesize-int64(i*filechunk))))
		buf := make([]byte, blocksize)
		file.Read(buf)
		io.WriteString(h, string(buf)) // append into the hash
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}

