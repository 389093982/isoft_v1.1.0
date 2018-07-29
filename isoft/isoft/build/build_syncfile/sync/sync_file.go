package sync

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"isoft/isoft/common/fileutil"
	"log"
	"path/filepath"
	"sync"
)

type SyncFile struct {
	XMLName xml.Name `xml:"build_syncfile"` // 指定最外层的标签为 build_syncfile
	Source  string   `xml:"source"`         // 读取source配置项,并将结果保存到Source变量中
	Targets []Target `xml:"target"`         // 读取target标签下的内容,以结构方式
}

type Target struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

func ReadSyncFile(filepath string) (syncFile SyncFile) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(content, &syncFile)
	if err != nil {
		log.Fatal(err)
	}
	return syncFile
}

// 同步所有目录
func StartAllSyncFile(dirPath string, syncFile SyncFile, filterTargetName string) {
	var wg sync.WaitGroup

	source := syncFile.Source
	targets := syncFile.Targets
	for _, target := range targets {
		if filterTargetName == "" || (filterTargetName != "" && filterTargetName == target.Name) {

			wg.Add(1)

			go func(src, target string) {
				StartOneSyncFile(src, target)

				defer wg.Done()

			}(filepath.Join(dirPath, source), filepath.Join(dirPath, target.Value))
		}
	}

	wg.Wait()
}

// 开始同步一个目录
func StartOneSyncFile(source, target string) {
	// 判断源文件夹是否存在
	if exist, _ := fileutil.PathExists(source); exist == false {
		fmt.Printf("file %s not exists!\n", source)
	}

	// 删除模板文件
	fileutil.RemoveFileOrDirectory(target)

	// 拷贝文件
	err := fileutil.CopyDir(source, target)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("copy dir %s to %s\n", source, target)
	}
}
