package hashutil

import (
	"fmt"
	"isoft/isoft/common/osutil"
	"testing"
)

func Test_GetLocalIp(t *testing.T) {
	fmt.Println(osutil.GetLocalIp())
}
