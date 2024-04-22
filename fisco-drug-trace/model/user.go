package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//### 用户基本信息
//- id
//- 账号
//- 呢称
//- 密码
//- 角色 （学生 老师 管理员）
//- 状态 （待审核 启用 禁用）
//- 注册时间

// User 用户模型
type User struct {
	gorm.Model
	Username       string
	PasswordDigest string
	Nickname       string
	Status         string
	Role           string
	RegisterTime   string
}

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
	// 状态
	// UserStatusEnabled 启用
	UserStatusEnabled = "enabled"
	// UserStatusDisabled 禁用
	UserStatusDisabled = "disabled"
	// 角色
	// UserRoleAdmin 管理员
	UserRoleAdmin = "admin"
	// 生成商
	UserRoleMerchant = "merchant"
	// 经销商
	UserRoleDealer = "dealer"
	// 用户
	UserRoleUser = "user"
)

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
