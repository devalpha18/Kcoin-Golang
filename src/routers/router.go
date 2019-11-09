package routers

import (
	"github.com/astaxie/beego"
	"Kcoin-Goloang/src/controllers"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/homepage", &controllers.HomePageController{})
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/join", &controllers.JoinController{})
}
