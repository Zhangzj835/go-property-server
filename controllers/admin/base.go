package admin

import (
	"go-property-server/enums"
	"go-property-server/service"
)

type BaseController struct {
	CommonController
	controllerName string
	actionName     string
	tokenUserInfo  service.TokenUserInfo
}

func (c *BaseController) Prepare() {
	c.controllerName, c.actionName = c.GetControllerAndAction()
	// 校验登录权限
	c.checkToken()
}

// 校验token
func (c *BaseController) checkToken() {
	var token string
	if c.Ctx.Request.Header["Token"] != nil {
		token = c.Ctx.Request.Header["Token"][0]
	}
	if token == "" {
		c.jsonResult("", enums.JRCodeFailed, "token不能为空")
	}
	tokenUserInfo := service.CheckToken(token)
	if tokenUserInfo.UserId == 0 {
		c.jsonResult("", enums.JRCode10001, "token已过期")
	} else {
		c.tokenUserInfo = tokenUserInfo
	}

}
