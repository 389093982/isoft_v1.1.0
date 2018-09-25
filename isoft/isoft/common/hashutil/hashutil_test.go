package hashutil

import (
	"fmt"
	"testing"
)

func Test_CalculateHashWithString(t *testing.T) {
	fmt.Println(CalculateHashWithString("admin"))
}

func Test_CalculateHashWithFile(t *testing.T) {
	hash1, _ := CalculateHashWithFile(`D:\zhourui\soft\zipAndExe\go\goland-2018.2.2.exe`)
	fmt.Println(hash1)

	hash3, _ := CalculateHashWithFileS(`D:\zhourui\soft\zipAndExe\go\goland-2018.2.2.exe`)
	fmt.Println(hash3)
}
