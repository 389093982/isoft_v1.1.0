package heartbeat

import (
	"isoft/isoft/common/logutil"
	"isoft/isoft_storage/lib/utils"
	"math/rand"
	"time"
)

// n 表示选取的随机数据服务节点数,exclude 表示返回的随机数据服务节点不能包含哪些节点
func ChooseRandomDataServers(n int, exclude map[int]string) (ds []string) {
	defer utils.RecordTimeCostForMethod("apiServer heartbeat ChooseRandomDataServers", time.Now())

	candidates := make([]string, 0)
	reverseExcludeMap := make(map[string]int)
	for id, addr := range exclude {
		// 排除节点主要用于数据修复时,根据目前已有的分片将丢失的分片复原出来并在此上传到数据服务
		// 目前已有的分片所在数据服务节点需要被排除
		reverseExcludeMap[addr] = id
	}
	servers := GetDataServers()
	for i := range servers {
		s := servers[i]
		_, excluded := reverseExcludeMap[s]
		if !excluded {
			candidates = append(candidates, s)
		}
	}
	length := len(candidates)
	// 判断节点数量是否满足要求
	if length < n {
		logutil.Errorln("ChooseRandomDataServers err, only find:", length)
		return
	}
	// length >= n 时
	// 将 0 ~ length-1 的所有整数乱序排列返回一个数组
	p := rand.Perm(length)
	// 取前 n 个作为 candidates 数组的下标取数据节点地址返回
	for i := 0; i < n; i++ {
		ds = append(ds, candidates[p[i]])
	}
	return
}
