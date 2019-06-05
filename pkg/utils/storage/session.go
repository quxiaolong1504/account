package storage

import (
	"github.com/go-redis/redis"
	"github.com/icza/session"
	"github.com/quxiaolong/account/pkg/utils"
)

var SessMgr session.Manager

func InitSessionManager(client *redis.Client) {
	store := utils.NewRedisStore(client)
	SessMgr = session.NewCookieManagerOptions(store,
		&session.CookieMngrOptions{
			SessIDCookieName:"q_x0",
		},
	)
}
