package models

import (
	"context"
	"git.in.zhihu.com/bit/zerzura/borm"
	"github.com/quxiaolong/account/pkg/consts"
	"time"
)

type User struct {
	ID int64
	Phone	string
	UID 	string
	Password string
	Status	consts.UserStatus
	CreatedAt	time.Time
	UpdatedAt	time.Time
	LastLoginAt	time.Time
}

func (u *User) TableName() string {
	return "account_user"
}

func (u *User) Profile(ctx context.Context) *UserProfile{
	var profile *UserProfile
	err := borm.New().Filter("uid", u.UID).One(ctx, profile)
	if err != nil {
		return nil
	}
	return profile
}


type UserProfile struct {
	ID int64
	UID 	string
	Avatar 	string
	Name	string
	Gender  consts.UserGender
}

func (u *UserProfile) TableName() string {
	return "account_user_profile"
}