package admin

import (
	"encoding/json"
	"go-property-server/enums"
	"go-property-server/models"

	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (c *CommonController) Prepare() {
	c.controllerName, c.actionName = c.GetControllerAndAction()

}

// 统一结果输出
func (c *CommonController) jsonResult(obj interface{}, code enums.JsonResultCode, msg string) {
	res := &models.JsonResult{Code: code, Msg: msg, Obj: obj}
	c.Data["json"] = res
	c.ServeJSON()
	c.StopRun()
}

// 解析请求参数
func (c *CommonController) paseRequestBody(params interface{}) {
	err := json.Unmarshal(c.Ctx.Input.RequestBody, params)
	if err != nil {
		c.jsonResult(map[string]string{}, enums.JRCodeFailed, "参数格式错误")
		return
	}
}
