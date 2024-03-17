package configs

import (
	"os"
)

type S_SMTPConfig struct {
	Port     string
	Host     string
	Username string
	Password string
}

func SMTPConfig() S_SMTPConfig {
	return S_SMTPConfig{
		Port:     os.Getenv("SMTP_PORT"),
		Host:     os.Getenv("SMTP_HOST"),
		Username: os.Getenv("SMTP_USERNAME"),
		Password: os.Getenv("SMTP_PASSWORD"),
	}
}
