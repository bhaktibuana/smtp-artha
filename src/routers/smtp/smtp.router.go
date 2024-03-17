package smtpRouter

import (
	smtpController "smtp-artha/src/controllers/smtp"

	"github.com/gin-gonic/gin"
)

func Routes(basePath string, router *gin.RouterGroup) {
	authGroup := router.Group(basePath)
	{
		authGroup.POST("/mail-verification", smtpController.MailVerification)
	}
}
