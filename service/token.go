package service

import (
	"go-property-server/models"
	"strconv"

	"go-property-server/utils"
)

func BuildToken(userId int) (string, error) {
	// 获取用户信息
	userInfo, err := models.SysUserOneByUserId(userId)
	if userInfo == nil {
		return "", err
	}
	// 这里获取用户的权限...
	token := utils.String2md5(strconv.Itoa(userInfo.Id) + userInfo.UserName)
	// 将权限信息保存到缓存中

	return token, nil
}

func GetInfoByToken() {

}

func CheckToken() {

}
