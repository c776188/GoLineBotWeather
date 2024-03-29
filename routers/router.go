package routers

import (
	"BeeLinebot/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/callback", &controllers.LineCallback{})
}
