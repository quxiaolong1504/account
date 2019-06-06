package utils

import (
	"github.com/quxiaolong/account/pkg/config"
	"github.com/silenceper/wechat"
	_ "github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
)

var WeChatCli  *wechat.Wechat


func InitWeChat(conf *config.WeChatConfig, redisConf *config.RedisConf) {
	WeChatCli = wechat.NewWechat(&wechat.Config{
		AppID: conf.AppID,
		AppSecret: conf.AppSecret,
		Cache: cache.NewRedis(&cache.RedisOpts{
			Host: redisConf.Host,
			Password: redisConf.Password,
			Database: int(redisConf.DB),
		}),
	})
}