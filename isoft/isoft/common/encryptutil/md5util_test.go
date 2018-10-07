package encryptutil

import (
	"fmt"
	"testing"
)

func Test_EncryptWithMD5(t *testing.T) {
	fmt.Println(EncryptWithMD5("admin"))
}
