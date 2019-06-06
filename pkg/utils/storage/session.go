package storage

import (
	"github.com/go-redis/redis"
	"github.com/icza/session"
	"github.com/quxiaolong/account/pkg/utils"
	"net/http"
	"time"
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

func NewSession(uid string, w http.ResponseWriter) session.Session {
	sess := session.NewSessionOptions(&session.SessOptions{
		Timeout: time.Hour * 24 * 30,
		IDLength: 144,
	})
	sess.SetAttr("uid", uid)
	SessMgr.Add(sess, w)
	return sess
}
