package cfg

import (
	"flag"
	"fmt"
	"github.com/larspensjo/config"
	"log"
)

const RABBITMQ_SERVER string = "RABBITMQ_SERVER"
const LISTEN_ADDRESS string = "LISTEN_ADDRESS"
const ES_SERVER string = "ES_SERVER"
const STORAGE_ROOT string = "STORAGE_ROOT"

var (
	configFile = flag.String("configfile", "D:/zhourui/program/go/goland_workspace/src/isoft/isoft_storage/cfg/config.ini", "configuration file")
)

var configmap = make(map[string]string)

func GetConfigValue(key string) string {
	if value, ok := configmap[key]; ok == true {
		return value
	}
	log.Println("get cfg value error!")
	return ""
}

func PutConfigValue(key, value string) {
	configmap[key] = value
}

func InitConfig(sectionSearch string) {
	// 读取默认配置
	initConfigData("default")
	// 读取 sectionSearch 相关配置
	initConfigData(sectionSearch)
}

func initConfigData(sectionSearch string) {
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if cfg.HasSection(sectionSearch) {
		section, err := cfg.SectionOptions(sectionSearch)
		if err == nil {
			for _, optionKey := range section {
				optionVal, err := cfg.String(sectionSearch, optionKey)
				if err == nil {
					configmap[optionKey] = optionVal
				}
			}
		}
	}
}
