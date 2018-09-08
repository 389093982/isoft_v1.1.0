package ziputil

import (
	"testing"
)

func Test_Compress(t *testing.T) {
	CompressZip("D:\\isoft_storage\\temp", "D:\\a.zip")
	DeCompressZip("D:\\a.zip", "D:\\abc")
}
