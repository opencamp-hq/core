package notify

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/opencamp-hq/core/models"
)

type SMTPConfig struct {
	Host     string
	Port     string
	Email    string
	Password string
}

type SMTPSender struct {
	cfg      SMTPConfig
	template *template.Template
}

func NewSMTPSender(cfg SMTPConfig) (*SMTPSender, error) {
	t, err := template.New("email").Parse(EmailTemplate)
	if err != nil {
		return nil, errors.New("Unable to parse email template")
	}

	return &SMTPSender{
		cfg:      cfg,
		template: t,
	}, nil
}

func (s SMTPSender) Send(cg *models.Campground, startDate, endDate string, sites models.Campsites) error {
	headers := make(map[string]string)
	headers["From"] = s.cfg.Email
	headers["To"] = s.cfg.Email
	headers["Subject"] = fmt.Sprintf("Good news, %s is available!", cg.Name)
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""

	buf := new(bytes.Buffer)
	for k, v := range headers {
		fmt.Fprintf(buf, "%s: %s\r\n", k, v)
	}
	buf.WriteString("\r\n")

	tmplData := &EmailData{
		Campground:     cg,
		StartDate:      startDate,
		EndDate:        endDate,
		AvailableSites: sites,
	}
	err := s.template.Execute(buf, tmplData)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", s.cfg.Email, s.cfg.Password, s.cfg.Host)
	addr := fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port)
	err = smtp.SendMail(addr, auth, s.cfg.Email, []string{s.cfg.Email}, buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}
