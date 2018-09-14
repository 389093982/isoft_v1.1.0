package hashutil

import (
	"fmt"
	"io"
	"crypto/sha256"
	"encoding/base64"
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