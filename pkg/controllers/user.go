package controllers

import (
	"github.com/quxiaolong/account/pkg/models"
	"github.com/quxiaolong/account/pkg/utils/storage"
)

func GetOrCreateUser(phone string) (*models.User, bool) {
	isNewUser := false
	user := &models.User{}
	storage.Mysql.GetSlave().Where("phone = ?", phone).First(user)

	if user.ID == 0 {
		user, _ = CreateUser(phone)
		isNewUser = true
	}
	return user, isNewUser
}

func CreateUser(phone string) (*models.User, error) {
	user := &models.User{
		Phone: phone,
	}
	// 这里考虑一下 db insert failed 的情况
	storage.Mysql.Master.Save(user)
	return user, nil
}

