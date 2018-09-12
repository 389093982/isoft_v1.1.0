package logutil

import (
	"testing"
)

func Test_Log(t *testing.T) {
	SetLogger(`D:/log`,"log.txt")
	Errorln("error....")
}
