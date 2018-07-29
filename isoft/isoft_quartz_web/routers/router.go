package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_quartz_web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
