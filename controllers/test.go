package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Index() {
	c.Data["json"] = map[string]string{"say": "hello worl"}
	c.ServeJSON()
}
