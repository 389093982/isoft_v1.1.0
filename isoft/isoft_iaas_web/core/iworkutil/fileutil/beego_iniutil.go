package fileutil

import "github.com/astaxie/beego/config"

func ReadBeegoIniFile(filename, section, key string) (string, error) {
	config, err := config.NewConfig("ini", filename)
	if err != nil {
		return "", err
	}
	return config.String(getAdapterKey(section, key)), nil
}

func WriteBeegoIniFile(filename, section, key, value string) error {
	config, err := config.NewConfig("ini", filename)
	if err != nil {
		return err
	}
	if err := config.Set(getAdapterKey(section, key), value); err != nil {
		return err
	}
	return config.SaveConfigFile(filename)
}

// ini 配置文件支持 section 操作, key 通过 section::key 方式获取
func getAdapterKey(section, key string) string {
	if section == "" {
		return key
	}
	return section + "::" + key
}
