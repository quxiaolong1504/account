package wechat

import (
	"github.com/gin-gonic/gin"
	"github.com/quxiaolong/account/pkg/controllers"
	"github.com/quxiaolong/account/pkg/errs"
	"github.com/quxiaolong/account/pkg/utils"
)

func GetToken(c *gin.Context) {

	// weChat miniprogram code to session key
	code := c.Param("code")
	ret, err := utils.WeChatCli.GetMiniProgram().Code2Session(code)
	if err != nil {
		panic(errs.WeChatLoginFailed)
	}
	accessToken, err := utils.WeChatCli.GetAccessToken()
	wxUserInfo, err := utils.WeChatCli.GetOauth().GetUserInfo(accessToken, ret.OpenID)
	user, isNew := controllers.GetOrCrateUserByUnionID(ret.UnionID, &wxUserInfo)
	controllers.WriteSession2Resp(user, isNew, c)
}