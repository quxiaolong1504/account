package controllers

import (
	"github.com/quxiaolong/account/pkg/models"
	"github.com/quxiaolong/account/pkg/utils/storage"
)

func CreateUser(Phone string) error {
	user := &models.User{
		Phone: Phone,
	}
	storage.Mysql.Master.Save(user)
	return nil
}