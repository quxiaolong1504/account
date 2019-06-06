package utils

import (
	"github.com/quxiaolong/account/pkg/config"
	"github.com/silenceper/wechat"
	_ "github.com/silenceper/wechat"
)

var WeChatCli  *wechat.Wechat


func InitWeChat(conf *config.WeChatConfig) {
	WeChatCli = wechat.NewWechat(&wechat.Config{
		AppID: conf.AppID,
		AppSecret: conf.AppSecret,
	})
}