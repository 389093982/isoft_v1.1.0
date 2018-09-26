package hashutil

import (
	"fmt"
	"testing"
)

func Test_CalculateHashWithString(t *testing.T) {
	fmt.Println(CalculateHashWithString("admin"))
}

func Test_CalculateHashWithFile(t *testing.T) {
	hash1, _ := CalculateHashWithFile(`D:\demo.mp4`)
	fmt.Println(hash1)
}
