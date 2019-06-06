package storage

import (
	"github.com/go-redis/redis"
	"github.com/icza/session"
	"github.com/quxiaolong/account/pkg/utils"
	"github.com/quxiaolong/account/pkg/utils/logger"
	"net/http"
	"time"
)

var SessMgr session.Manager

func InitSessionManager(client *redis.Client) {
	logger.Logger.Infof("Init Session Manager")
	store := utils.NewRedisStore(client)
	SessMgr = session.NewCookieManagerOptions(store,
		&session.CookieMngrOptions{
			SessIDCookieName:"q_x0",
		},
	)
}

func NewSession(uid string, w http.ResponseWriter) session.Session {
	sess := session.NewSessionOptions(&session.SessOptions{
		Timeout: time.Hour * 24 * 30,
		IDLength: 144,
	})
	sess.SetAttr("uid", uid)
	SessMgr.Add(sess, w)
	return sess
}
