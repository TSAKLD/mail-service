package bootstrap

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	GmailLogin    string
	GmailPassword string
	GmailHost     string
	GmailPort     string

	HTTPPort string

	KafkaAddr    string
	KafkaTopic   string
	KafkaGroupID string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		GmailLogin:    os.Getenv("GMAIL_LOGIN"),
		GmailPassword: os.Getenv("GMAIL_PASSWORD"),
		GmailHost:     os.Getenv("GMAIL_HOST"),
		GmailPort:     os.Getenv("GMAIL_PORT"),
		HTTPPort:      os.Getenv("HTTP_PORT"),
		KafkaAddr:     os.Getenv("KAFKA_ADDR"),
		KafkaTopic:    os.Getenv("KAFKA_TOPIC"),
		KafkaGroupID:  os.Getenv("KAFKA_GROUP_ID"),
	}, nil
}

func (c *Config) Validate() []error {
	var errorList []error

	if c.GmailLogin == "" {
		err := errors.New("invalid GMAIL Login field \n")
		errorList = append(errorList, err)
	}

	if c.GmailPassword == "" {
		err := errors.New("invalid GMAIL Password field \n")
		errorList = append(errorList, err)
	}

	if c.HTTPPort == "" {
		err := errors.New("invalid HTTP port field \n")
		errorList = append(errorList, err)
	}

	if c.GmailHost == "" {
		err := errors.New("invalid Gmail host field \n")
		errorList = append(errorList, err)
	}

	if c.GmailPort == "" {
		err := errors.New("invalid Gmail port field \n")
		errorList = append(errorList, err)
	}

	if len(errorList) != 0 {
		return errorList
	}

	return nil
}
