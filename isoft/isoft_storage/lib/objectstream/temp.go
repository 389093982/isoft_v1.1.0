package objectstream

import (
	"fmt"
	"io/ioutil"
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/lib/utils"
	"net/http"
	"strings"
	"time"
)

type TempPutStream struct {
	Server string
	Uuid   string
}

func NewTempPutStream(server, object string, size int64) (*TempPutStream, error) {
	// 先调用数据服务 temp 接口的 post 方法生产临时文件,接收返回的 uuid 信息
	request, err := http.NewRequest("POST", "http://"+server+"/temp/"+object, nil)
	if err != nil {
		logutil.Errorln(err)
		return nil, err
	}
	request.Header.Set("size", fmt.Sprintf("%d", size))
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logutil.Errorln(err)
		return nil, err
	}
	// 接收返回的 uuid 信息
	uuid, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logutil.Errorln(err)
		return nil, err
	}
	return &TempPutStream{server, string(uuid)}, nil
}

func (w *TempPutStream) Write(p []byte) (n int, err error) {
	request, err := http.NewRequest("PATCH", "http://"+w.Server+"/temp/"+w.Uuid, strings.NewReader(string(p)))
	if err != nil {
		logutil.Errorln(err)
		return 0, err
	}
	client := http.Client{}
	r, err := client.Do(request)
	if err != nil {
		logutil.Errorln(err)
		return 0, err
	}
	if r.StatusCode != http.StatusOK {
		logutil.Errorln("dataServer return http code %d", r.StatusCode)
		return 0, fmt.Errorf("dataServer return http code %d", r.StatusCode)
	}
	return len(p), nil
}

func (w *TempPutStream) Commit(good bool) {
	method := "DELETE"
	if good {
		method = "PUT"
	}

	defer utils.RecordTimeCostForMethod("lib objectstream Commit "+method, time.Now())

	// delete 情况下删除临时文件, put 情况下将临时文件转正
	request, _ := http.NewRequest(method, "http://"+w.Server+"/temp/"+w.Uuid, nil)
	client := http.Client{}
	client.Do(request)
}

func NewTempGetStream(server, uuid string) (*GetStream, error) {
	return newGetStream("http://" + server + "/temp/" + uuid)
}
