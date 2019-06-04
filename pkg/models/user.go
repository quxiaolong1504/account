package models

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/quxiaolong/account/pkg/config"
	"github.com/quxiaolong/account/pkg/consts"
	"github.com/quxiaolong/account/pkg/rpcs/seqd"
	"github.com/quxiaolong/account/pkg/utils"
	"github.com/quxiaolong/account/pkg/utils/storage"
	"time"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	Phone    string `grom:"type:varchar(50);unique_index"`
	UID      string `gorm:"type:varchar(100);unique_index"`
	Password string
	Status   consts.UserStatus `gorm:"default:1"`
	DateTimeModel
	//LastLoginAt time.Time
}

func (u *User) TableName() string {
	return "user"
}

// gen uid for user
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UID", fmt.Sprintf("%d", seqd.Generate()))
	scope.SetColumn("LastLoginAt", time.Now())
	return nil
}

func (u *User) SetPassword(password string) {
	slat := fmt.Sprintf("%s:%s:%s", u.UID, password, config.Conf.Auth.SecretKey)
	sha := sha512.New()
	sha.Write([]byte(slat))
	u.Password = base64.URLEncoding.EncodeToString(sha.Sum(nil))
	u.Password = utils.MakePassword(password, u.UID)
	storage.Mysql.Master.Save(u)
}

func (u *User) VerifyPassword(password string) bool {
	return u.Password == utils.MakePassword(password, u.UID)
}

type UserProfile struct {
	ID     uint   `gorm:"primary_key"`
	UID    string `grom:"type:varchar(50);unique_index"`
	Avatar string
	Name   string
	Gender consts.UserGender `gorm:"default:1"`
}

func (u *UserProfile) TableName() string {
	return "user_profile"
}
