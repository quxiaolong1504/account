package storage

import (
	"github.com/go-redis/redis"
	"github.com/icza/session"
	"github.com/quxiaolong/account/pkg/utils"
)

var SessMgr session.Manager

func InitSessionManager(client *redis.Client) {
	SessMgr = session.NewCookieManager(utils.NewRedisStore(client))
}
