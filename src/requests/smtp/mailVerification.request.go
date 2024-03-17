package smtpRequest

import (
	"net/http"
	"smtp-artha/src/helpers"

	"github.com/gin-gonic/gin"
)

type S_MailVerificationRequest struct {
	Name       string `json:"name" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required"`
	AppLogoUrl string `json:"app_logo_url" binding:"required"`
	LoginUrl   string `json:"login_url" binding:"required"`
	ActionUrl  string `json:"action_url" binding:"required"`
}

func MailVerification(context *gin.Context) *S_MailVerificationRequest {
	var request S_MailVerificationRequest

	if err := context.ShouldBindJSON(&request); err != nil {
		helpers.Response(err.Error(), http.StatusBadRequest, context, nil)
		return nil
	}

	return &request
}
