package main

import (
	"log"
	"mail.project/api"
	"mail.project/bootstrap"
	"mail.project/service"
	"net/http"
)

func main() {
	cfg, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatal("Problem with config load: ", err)
	}

	errorList := cfg.Validate()
	if errorList != nil {
		log.Fatal("Problem with config validation: ", errorList)
	}

	kafkaReader, err := bootstrap.KafkaConnect(cfg.KafkaAddr, cfg.KafkaTopic, cfg.KafkaGroupID)
	if err != nil {
		log.Fatal("kafka connection:", err)
	}
	log.Println("kafka connection: OK")

	ms := service.New(cfg, kafkaReader)
	h := api.NewHandler(ms)

	go ms.OnCreateUserEvent()

	http.HandleFunc("POST /mail", h.SendMail)
	err = http.ListenAndServe(":"+cfg.HTTPPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
