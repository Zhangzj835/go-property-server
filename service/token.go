package service

import (
	"encoding/json"
	"fmt"
	"go-property-server/models"
	"strconv"
	"time"

	"go-property-server/utils"
)

type TokenUserInfo struct {
	UserId   int
	UserName string
}

const cacheTokenPre = "token_"

func BuildToken(userId int) (string, error) {
	// 获取用户信息
	userInfo, err := models.SysUserOneByUserId(userId)
	if userInfo == nil {
		return "", err
	}
	// 这里获取用户的权限...
	token := utils.String2md5(strconv.Itoa(userInfo.Id) + userInfo.UserName + strconv.Itoa(int(time.Now().Unix())))
	// 将权限信息保存到缓存中
	var cacheData TokenUserInfo
	cacheData.UserId = userInfo.Id
	cacheData.UserName = userInfo.UserName
	cacheStr, _ := json.Marshal(cacheData)
	cRes := SetCache(cacheTokenPre+token, cacheStr, 3600)
	if cRes != nil {
		fmt.Printf("缓存服务异常:%v\n", cRes)
	}
	return token, nil
}

func DelToken(token string) {
	DelCache(cacheTokenPre + token)
}

func CheckToken(token string) TokenUserInfo {
	var data []uint8
	var tokenUserInfo TokenUserInfo
	_ = GetCache(cacheTokenPre+token, &data)
	_ = json.Unmarshal(data, &tokenUserInfo)
	return tokenUserInfo
}
