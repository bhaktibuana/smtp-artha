package smtpController

import (
	"net/http"
	"smtp-artha/src/helpers"
	smtpRequest "smtp-artha/src/requests/smtp"
	smtpResult "smtp-artha/src/results/smtp"
	smtpService "smtp-artha/src/services/smtp"

	"github.com/gin-gonic/gin"
)

func MailVerification(context *gin.Context) {
	request := smtpRequest.MailVerification(context)
	if request == nil {
		return
	}

	userInfo := smtpService.MailVerification(context, request)
	if userInfo == nil {
		return
	}

	helpers.Response("Request successful", http.StatusOK, context, smtpResult.MailVerification(userInfo))
}
