package rabbitmq

import (
	"fmt"
	"strconv"
	"testing"
)

const host = "amqp://test:test@193.112.162.61:5673"

func Test_New(t *testing.T) {
	q := New(host)
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consume()
	// 循环遍历数据服务监听地址
	for msg := range c {
		// 获取监听地址
		dataServer, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(dataServer)
		}
	}
}
