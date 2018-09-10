package cfg

import (
	"fmt"
	"github.com/larspensjo/config"
	"isoft/isoft/common/fileutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetConfigValue(key string) string {
	if value, ok := configmap[key]; ok == true {
		return value
	}
	log.Println("get cfg value error!")
	return ""
}

func InitConfigWithOsArgs(args []string) {
	if len(args) != 3 {
		log.Fatal("Input parameter length is not valid...")
	}
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	// 初始化配置,两个参数分别是环境名称和 section 名称
	initCfg(args[1], args[2])
}

func initCfg(env_name, section_name string) {
	initCfgFilePath(env_name)
	// 读取默认配置
	initConfigData("default")
	// 读取 sectionSearch 相关配置
	initConfigData(section_name)
}

// 配置文件路径
var configFilePath string

func initCfgFilePath(env_name string) {
	// 获取环境变量
	env_var := os.Getenv("ISOFT_STORAGE_CFG")
	if strings.TrimSpace(env_var) == "" {
		panic("Failed to get the environment variable of ISOFT_STORAGE_CFG")
	}
	configFilePath = fileutil.ChangeToLinuxSeparator(filepath.Join(env_var, env_name+"_cfg.ini"))
}

// 存储所有配置信息
var configmap = make(map[string]string)

// 根据 section_name 加载所有配置到 configmap 中去
func initConfigData(section_name string) {
	cfg, err := config.ReadDefault(configFilePath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read config file, %s", err.Error()))
	}
	if !cfg.HasSection(section_name) {
		panic("Failed to read section info with config file...")
	}
	if section, err := cfg.SectionOptions(section_name); err == nil {
		for _, optionKey := range section {
			if optionVal, err := cfg.String(section_name, optionKey); err == nil {
				configmap[optionKey] = optionVal
			}
		}
	}
}
