package smtpResult

import smtpRequest "smtp-artha/src/requests/smtp"

type S_MailVerificationResult struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func MailVerification(userInfo *smtpRequest.S_MailVerificationRequest) S_MailVerificationResult {
	return S_MailVerificationResult{
		Name:     userInfo.Name,
		Username: userInfo.Username,
		Email:    userInfo.Email,
	}
}
