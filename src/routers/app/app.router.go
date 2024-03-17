package appRouter

import (
	smtpRouter "smtp-artha/src/routers/smtp"

	"github.com/gin-gonic/gin"
)

func AppRouters(path string, router *gin.Engine) {
	apiGroup := router.Group(path)
	{
		smtpRouter.Routes("/smtp", apiGroup)
	}
}
