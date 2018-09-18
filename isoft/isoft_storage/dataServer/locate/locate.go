package locate

import (
	"isoft/isoft_storage/cfg"
	"isoft/isoft_storage/lib/models"
	"isoft/isoft_storage/lib/rabbitmq"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

var objects = make(map[string]int)
var mutex sync.Mutex

// 传入 hash 值返回存储的分片 id
func Locate(hash string) int {
	mutex.Lock()
	id, ok := objects[hash]
	mutex.Unlock()
	if !ok {
		return -1
	}
	return id
}

func Add(hash string, id int) {
	mutex.Lock()
	objects[hash] = id
	mutex.Unlock()
}

func Del(hash string) {
	mutex.Lock()
	delete(objects, hash)
	mutex.Unlock()
}

func StartLocate() {
	// 直接将 RabbitMQ 消息队列里收到的对象散列值作为 Locate 参数
	q := rabbitmq.New(cfg.GetConfigValue(cfg.RABBITMQ_SERVER))
	defer q.Close()
	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		// 接收 hash 值
		hash, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		// 定位 hash 值是否存在
		id := Locate(hash)
		if id != -1 {
			// 不存在则不返回消息,存在则返回消息
			q.Send(msg.ReplyTo, models.LocateMessage{Addr: cfg.GetConfigValue(cfg.LISTEN_ADDRESS), Id: id})
		}
	}
}

// 应用启动时对节点本地磁盘上的对象进行定位的,缓存对象定位信息,防止过于频繁的磁盘访问
func CollectObjects() {
	// 格式如下 66WuRH0s0albWDZ9nTmjFo9JIqTTXmB6EiRkhTh1zeA=.0.xPZ9Cf8mShrJsL32FnbSVcayc9W5Y05clRo3GOkLyG0=
	files, _ := filepath.Glob(cfg.GetConfigValue(cfg.STORAGE_ROOT) + "/objects/*")
	for i := range files {
		file := strings.Split(filepath.Base(files[i]), ".")
		if len(file) != 3 {
			panic(files[i])
		}
		// 第一个元素是存储对象的 hash 值
		hash := file[0]
		// 第二个参数是分片 id,第三个参数是分片对应的 hash 值
		id, e := strconv.Atoi(file[1])
		if e != nil {
			panic(e)
		}
		objects[hash] = id
	}
}
