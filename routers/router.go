package routers

import (
	"auditaudit/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/test", &controllers.MainController{}, "get:Test")
    beego.Router("/add", &controllers.ManageController{}, "get,post:Add")
    beego.Router("/view", &controllers.ManageController{}, "get:View")
	beego.Router("/datatables/data/json", &controllers.OperationRecord{},"*:AjaxData")
}
