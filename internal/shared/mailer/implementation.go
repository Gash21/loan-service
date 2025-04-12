package mailer

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/Gash21/amartha-test/internal/shared/logger"
	"go.uber.org/zap"
	gomail "gopkg.in/mail.v2"
)

func (s *MailerAdapter) SendEmail(opts EmailOpts) error {
	l := logger.NewLogger("MailerAdapter", "SendEmail")
	message := gomail.NewMessage()
	// Set email headers
	message.SetHeader("From", s.Sender)
	message.SetHeader("To", opts.To...)
	message.SetHeader("Subject", opts.Subject)
	for _, body := range opts.Body {
		for contentType, content := range body {
			message.SetBody(contentType, content.(string))
		}
	}

	// Set up the SMTP dialer
	dialer := gomail.NewDialer(s.Host, s.Port, s.Username, s.Password)
	if s.TLS {
		dialer.StartTLSPolicy = gomail.MandatoryStartTLS // Ensure TLS is mandatory
		dialer.SSL = true
	} else {
		dialer.SSL = false
	}

	err := dialer.DialAndSend(message)
	if err != nil {
		l.Error("Failed to send email", zap.Error(err))
		return err
	}

	return nil
}

func (s *MailerAdapter) PrepareHtmlBody(templateStr string, params map[string]interface{}) (string, error) {
	tmpl, err := template.New("email").Funcs(template.FuncMap{
		"toUpper": strings.ToUpper,
	}).Parse(templateStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, params)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
