package admin

import "go-property-server/enums"

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	result := map[string]string{"userName": c.tokenUserInfo.UserName}
	c.jsonResult(result, enums.JRCodeSucc, "")
}
