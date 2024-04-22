package service

import (
	"fisco-drug-trace/model"
	"fisco-drug-trace/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname string `form:"nickname" json:"nickname" `
	Username string `form:"username" json:"username" `
	Password string `form:"password" json:"password" `
	Role     string `form:"role" json:"role" binding:"required"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	count := int64(0)
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("username = ?", service.Username).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	// 检测用户名是否已经注册
	tx := model.DB.First(&model.User{}, "username = ?", service.Username)
	if tx.RowsAffected > 0 {
		return serializer.ParamErr("用户名已经注册", nil)
	}
	user := model.User{
		Nickname: service.Nickname,
		Username: service.Username,
		Role:     service.Role,
		Status:   model.UserStatusDisabled,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}
