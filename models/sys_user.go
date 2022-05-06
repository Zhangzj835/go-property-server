package models

import "github.com/astaxie/beego/orm"

type SysUser struct {
	Id           int
	UserName     string
	UserPwd      string
	Nickname     string
	RealName     string
	Phone        string
	Email        string
	Status       int
	Avatar       string
	RegisterTime string
}

func init() {
	orm.RegisterModel(new(SysUser))
}

func (c *SysUser) TableName() string {
	return SysUserTBName()
}

//获取 SysUser 对应的表名称
func SysUserTBName() string {
	return "sys_user"
}

// 根据用户名获取单条
func SysUserOneByUserName(userName string) (*SysUser, error) {
	m := SysUser{}
	err := orm.NewOrm().QueryTable(SysUserTBName()).Filter("user_name", userName).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// 根据用户ID获取用户信息
func SysUserOneByUserId(userId int) (*SysUser, error) {
	m := SysUser{}
	err := orm.NewOrm().QueryTable(SysUserTBName()).Filter("id", userId).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
