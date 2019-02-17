package imodules

import (
	"github.com/astaxie/beego"
	"strings"
)

func CheckModule(moduleName string) bool {
	return strings.Contains(beego.AppConfig.String("open.modules"), moduleName)
}
