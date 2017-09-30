package email

import (
	"bytes"
	"fmt"
	"github.com/danman113/gobeet/config"
	"github.com/danman113/gobeet/site"
	"html/template"
	"net/smtp"
)

var (
	// Set from main
	Config        config.EmailConfig
	auth          smtp.Auth
	EmailPassword *string

	// Temp variables
	temp *template.Template
)

type TemplateData struct {
	Config config.EmailConfig
	Site   *site.Page
	Error  error
}

func parseTemplate() {
	if temp == nil {
		t, tErr := template.ParseFiles(Config.Template)
		if tErr != nil {
			fmt.Println("Template Parse Error: ")
			fmt.Println(tErr)
			return
		}
		temp = t
	}
}

func buildEmailBody(err error, pg *site.Page) ([]byte, error) {
	if temp == nil {
		parseTemplate()
	}

	buf := new(bytes.Buffer)

	eErr := temp.Execute(buf, TemplateData{Config, pg, err})

	if eErr != nil {
		fmt.Println("Template Execute Error: ")
		fmt.Println(eErr)
		return nil, eErr
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	return []byte(fmt.Sprintf("Subject: %s\r\n%s %s", "Gobeet Alert", mime, buf.String())), nil
}

func SendAlert(err error, pg *site.Page) {
	if auth == nil {
		auth = smtp.PlainAuth(
			"",
			Config.Address,
			*EmailPassword,
			Config.Server,
		)
	}

	body, bodyErr := buildEmailBody(err, pg)

	if bodyErr != nil {
		return
	}

	sendErr := smtp.SendMail(
		Config.Server+":"+Config.Port,
		auth,
		Config.Address,
		Config.Recipients,
		body,
	)
	if sendErr != nil {
		fmt.Println("Email Error!")
		fmt.Println(sendErr)
	}
}
