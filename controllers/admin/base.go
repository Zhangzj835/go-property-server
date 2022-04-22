package admin

import (
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (c *BaseController) Prepare() {
	c.controllerName, c.actionName = c.GetControllerAndAction()
	// 校验登录权限
	c.checkToken()
}

// 校验token
func (c *BaseController) checkToken() {
	fmt.Println("这里校验token")
}
