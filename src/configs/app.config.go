package configs

import (
	"os"
)

type S_AppConfig struct {
	Port      string
	GinMode   string
	BaseUrl   string
	JwtSecret string
}

func AppConfig() S_AppConfig {
	return S_AppConfig{
		Port:      ":" + os.Getenv("PORT"),
		GinMode:   os.Getenv("GIN_MODE"),
		BaseUrl:   os.Getenv("BASE_URL"),
		JwtSecret: os.Getenv("JWT_SECRET_KEY"),
	}
}
