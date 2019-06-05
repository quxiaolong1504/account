package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/utils/storage"
)

func Authentication() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		sess := storage.SessMgr.Get(c.Request)
		uid := "0"
		if sess != nil {
			uid = sess.Attr("uid").(string)
		}
		c.Set("uid", uid)
		c.Next()
	}

	return gin.HandlerFunc(fn)
}
