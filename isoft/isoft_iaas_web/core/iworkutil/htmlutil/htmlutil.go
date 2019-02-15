package htmlutil

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func getUrlBytes(url string) (bytes []byte) {
	resp,err := http.Get(url)
	if err!=nil {
		return
	}
	defer resp.Body.Close()
	bytes, _ = ioutil.ReadAll(resp.Body)
	return
}

func GetAllHref(url string) (results []string) {
	// 发送 HTTP 请求获取 URL 地址 bytes
	body := getUrlBytes(url)
	href_pattern :=   "href=\"(.+?)\""
	href_reg := regexp.MustCompile(href_pattern)
	// 根据正则表达式获取所有的地址
	hrefs := href_reg.FindAllString(string(body),-1)
	for _,v := range hrefs{
		str := strings.Split(v,"\"")[1]
		if len(str)<1{
			continue
		}
		switch str[0]{
		case 'h':
			results = append(results,str)
		case '/':
			if len(str)!=1 && str[1]=='/'{
				results = append(results,"http:"+str)
			}

			if len(str)!=1 && str[1]!='/'{
				results = append(results,url+str[1:])
			}
		default:
			results = append(results,url+str)
		}
	}
	return results
}
