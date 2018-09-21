package es

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func Test_New(t *testing.T) {
	client := http.Client{}
	url := fmt.Sprintf("http://%s/metadata/_search", "193.112.162.61:9200")
	body := fmt.Sprintf(`
        {
			"from": %d, "size": %d,
			"sort":[
				{"name":{"order":"asc"}},
				{"version":{"order":"desc"}}
			],
			"query":{
				"wildcard":{"name":"*%s*"}
			}
		}`, 0, 10, "ng")
	fmt.Println(body)
	request, _ := http.NewRequest("POST", url, strings.NewReader(body))
	r, e := client.Do(request)
	if e != nil {
		fmt.Println(e)
	}
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(b))

}
