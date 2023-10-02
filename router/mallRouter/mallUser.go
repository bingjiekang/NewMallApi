package mallrouter

import (
	"NewMallApi/utils/verification"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	userRoute := router.Group("v1")
	{
		userRoute.Any("/user/register", MallUser.Register) // 用户注册
		userRoute.Any("/user/login", MallUser.Login)
		userRoute.Any("/user/index", MallUser.Index)
		userRoute.GET("/captcha", verification.Captcha)
	}

}
