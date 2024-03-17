package smtpService

import (
	"net/http"
	"smtp-artha/src/helpers"
	smtpRequest "smtp-artha/src/requests/smtp"
	templates "smtp-artha/src/templates/mailVerification"

	"github.com/gin-gonic/gin"
)

func MailVerification(context *gin.Context, request *smtpRequest.S_MailVerificationRequest) *smtpRequest.S_MailVerificationRequest {
	template := templates.MailVerification(templates.S_MailVerificationProps{
		Name:       request.Name,
		Username:   request.Username,
		Email:      request.Email,
		AppLogoUrl: request.AppLogoUrl,
		LoginUrl:   request.LoginUrl,
		ActionUrl:  request.ActionUrl,
	})

	recipient := helpers.S_Recipient{To: []string{request.Email}}

	if err := helpers.SendMail(recipient, "Welcome to Artha", template); err != nil {
		helpers.Response(err.Error(), http.StatusInternalServerError, context, nil)
		return nil
	}

	userInfo := smtpRequest.S_MailVerificationRequest{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
	}

	return &userInfo
}
