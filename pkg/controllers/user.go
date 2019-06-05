package controllers

import (
	"github.com/quxiaolong/account/pkg/models"
	"github.com/quxiaolong/account/pkg/utils/storage"
)

func CreateUser(Phone string) (*models.User, error) {
	user := &models.User{
		Phone: Phone,
	}
	// 这里考虑一下 db insert failed 的情况
	storage.Mysql.Master.Save(user)
	return user, nil
}