package controllers

import (
	"github.com/quxiaolong/account/pkg/models"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"github.com/silenceper/wechat/oauth"
)

func GetOrCrateUserByUnionID(unionid string, userInfo *oauth.UserInfo) (*models.User, bool){
	user := &models.User{}
	isNew := false
	weChatUserShip := &models.UserWeChatShip{}
	storage.Mysql.GetSlave().Where("unionid = ?", unionid).First(weChatUserShip)

	if weChatUserShip.ID == 0 {
		u, _ := CreateUser("")
		user = u
		isNew = true
	}

	return user, isNew
}

func BindWechat() {

}

func UnBindWechat(user *models.User) {

}