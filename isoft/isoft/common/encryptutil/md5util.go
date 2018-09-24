package encryptutil

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func EncryptWithMD5(decrypt string) string {
	h := md5.New()
	_, err := h.Write([]byte(decrypt)) // 	需要加密的字符串为 decrypt
	if err != nil {
		fmt.Println("EncryptWithMD5 error")
		return ""
	}
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 返回加密结果
}
