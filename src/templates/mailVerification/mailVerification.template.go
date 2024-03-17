package templates

import (
	"fmt"
	"os"
	"smtp-artha/src/configs"
	"strings"
	"text/template"
)

type S_MailVerificationProps struct {
	Name       string
	LoginUrl   string
	Username   string
	Email      string
	ActionUrl  string
	AppLogoUrl string
}

func MailVerification(props S_MailVerificationProps) *strings.Builder {
	var bodyContent strings.Builder
	var templateFile string

	if configs.AppConfig().GinMode == "debug" {
		templateFile = "./src/templates/mailVerification/index.html"
	} else {
		templateFile = "./assets/templates/mailVerification/index.html"
	}

	tmplContent, err := os.ReadFile(templateFile)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		return nil
	}

	tmpl, err := template.New("email").Parse(string(tmplContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return nil
	}

	if err := tmpl.Execute(&bodyContent, props); err != nil {
		return nil
	}

	return &bodyContent
}
