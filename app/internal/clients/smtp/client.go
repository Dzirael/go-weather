package smtp

import (
	"go-weather/app/internal/config"
	"go-weather/app/internal/models"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	dialer *gomail.Dialer
	config *config.Config
}

func New(cfg *config.Config) *Mailer {
	d := gomail.NewDialer(cfg.SMTP.Host, cfg.SMTP.Port, cfg.SMTP.Username, cfg.SMTP.Password)

	return &Mailer{
		dialer: d,
		config: cfg,
	}
}

func (m *Mailer) Send(to string, msg *models.MailMessage) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", m.config.SMTP.Username)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", msg.Subject)
	mail.SetBody("text/html", msg.Message)

	err := m.dialer.DialAndSend(mail)
	return err
}
