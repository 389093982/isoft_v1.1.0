package task

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_HttpGet(t *testing.T) {
	response, err := http.Get("http://193.112.162.61:8080/user/login")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(body))
}
