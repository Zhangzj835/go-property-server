package admin

import (
	"fmt"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	fmt.Println("123")
	c.Data["json"] = map[string]string{"name": "李四"}
	c.ServeJSON()
}
