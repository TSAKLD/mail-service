package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"mail.project/bootstrap"
	"mail.project/entity"
	"net/smtp"
)

type MailService struct {
	cfg   *bootstrap.Config
	kafka *kafka.Reader
}

func New(cfg *bootstrap.Config, kr *kafka.Reader) *MailService {
	return &MailService{
		cfg:   cfg,
		kafka: kr,
	}
}

func (ms *MailService) SendMail(mail entity.Mail) error {
	auth := smtp.PlainAuth("", ms.cfg.GmailLogin, ms.cfg.GmailPassword, ms.cfg.GmailHost)

	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", mail.Receiver, mail.Subject, mail.Message))

	err := smtp.SendMail(fmt.Sprintf("%s:%s", ms.cfg.GmailHost, ms.cfg.GmailPort), auth, ms.cfg.GmailLogin, []string{mail.Receiver}, msg)
	if err != nil {
		return err
	}

	return nil
}

func (ms *MailService) OnCreateUserEvent() {
	for {
		err := func() error {
			msg, err := ms.kafka.ReadMessage(context.Background())
			if err != nil {
				return err
			}

			var message entity.Mail
			err = json.Unmarshal(msg.Value, &message)
			if err != nil {
				return err
			}
			log.Println("got kafka message:", message)

			err = ms.SendMail(message)
			if err != nil {
				return err
			}

			return nil
		}()
		if err != nil {
			log.Println("OnCreateUserEvent:", err)
		}
	}
}
