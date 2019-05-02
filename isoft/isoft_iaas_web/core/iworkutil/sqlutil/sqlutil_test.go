package sqlutil

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_SqlParse(t *testing.T) {
	sql := `select {{id,app_address,created_by,created_time,last_updated_by,last_updated_time}} 
from {{app_register}}`
	reg := regexp.MustCompile("^.+{{.+}}.+{{.+}}.*$")
	fmt.Println(reg.Match([]byte(sql)))
}
