package httputil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_Http(t *testing.T) {
	go func() {
		url := "http://127.0.0.1:9002/locate/pS5m73EPd5zPuUAuoFRHbo1seh868Uy8G6bLHf7xGOY="
		resp, err := http.Post(url,"application/x-www-form-urlencoded", nil)
		if err == nil || resp.StatusCode == 200{
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil{
				fmt.Println(string(body))
			}
		}else{
			fmt.Println(err)
		}
	}()
	time.Sleep(time.Second * 5)
}


