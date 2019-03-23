package fileutil

import (
	"gopkg.in/ini.v1"
)

type IniParser struct {
	conf_reader *ini.File
}

func (this *IniParser) load(filename string) (err error) {
	this.conf_reader, err = ini.Load(filename)
	return
}

func (this *IniParser) getValue(section string, key string) string {
	if this.conf_reader == nil {
		return ""
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return ""
	}
	return s.Key(key).String()
}

// 根据 section 和 key 设置 value, 不存在则添加,存在则覆盖
func (this *IniParser) setValue(section string, key, value string) {
	s := this.conf_reader.Section(section)
	s.Key(key).SetValue(value)
}

// 重新写回文件
func (this *IniParser) RewriteToFile(filename string) error {
	return this.conf_reader.SaveTo(filename)
}

func WriteIniFile(filename, section, key, value string) error {
	parser := &IniParser{}
	if err := parser.load(filename); err != nil {
		return err
	}
	parser.setValue(section, key, value)
	return parser.RewriteToFile(filename)
}

func ReadIniFile(filename, section, key string) (string, error) {
	parser := &IniParser{}
	if err := parser.load(filename); err != nil {
		return "", err
	}
	return parser.getValue(section, key), nil
}
