package admin

import (
	"go-property-server/enums"
	"go-property-server/models"
	"go-property-server/service"
	"go-property-server/utils"
	"strings"
)

type LoginController struct {
	CommonController
}

// 登录
func (c *LoginController) Dologin() {
	var params struct {
		UserName string `json:"user_name"`
		UserPwd  string `json:"user_pwd"`
	}
	c.paseRequestBody(&params)
	userName := strings.TrimSpace(params.UserName)
	userPwd := strings.TrimSpace(params.UserPwd)
	if len(userName) == 0 || len(userPwd) == 0 {
		c.jsonResult("", enums.JRCodeFailed, "账号或密码不能为空")
	}
	userPwd = utils.String2md5(userPwd)
	// 获取数据
	userInfo, _ := models.SysUserOneByUserName(userName)
	// 校验账号密码
	if userInfo == nil {
		c.jsonResult("", enums.JRCodeFailed, "账号不存在")
	}
	if userInfo.UserPwd != userPwd {
		c.jsonResult("", enums.JRCodeFailed, "账号或密码有误")
	}
	if userInfo.Status == 0 {
		c.jsonResult("", enums.JRCodeFailed, "账号被禁用")
	}
	// 账号密码正确，生成token
	token, _ := service.BuildToken(userInfo.Id)
	// 数据输出
	result := map[string]string{"token": token}
	c.jsonResult(result, enums.JRCodeSucc, "")
}

// 退出登录
func (c *LoginController) Logout() {
	var token string
	if c.Ctx.Request.Header["Token"] != nil {
		token = c.Ctx.Request.Header["Token"][0]
	}
	if token == "" {
		c.jsonResult("", enums.JRCodeFailed, "token为空")
	}
	service.DelToken(token)
	result := map[string]string{"status": "ok"}
	c.jsonResult(result, enums.JRCodeSucc, "")
}
