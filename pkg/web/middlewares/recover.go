package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/errs"
	"net/http"
)

func Recover() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		defer func() {
			if rvr := recover(); rvr != nil {
				switch rvr.(type) {
				case *errs.APIError:
					err := rvr.(*errs.APIError)
					c.JSON(err.HttpStatusCode, err.DetailError)
				default:
					c.JSON(http.StatusInternalServerError, gin.H{"message": "Server internal error!"})
				}
			}
		}()

		c.Next()
	}

	return gin.HandlerFunc(fn)
}
