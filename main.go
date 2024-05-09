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

	ms := service.New(cfg)
	h := api.NewHandler(ms)

	http.HandleFunc("POST /mail", h.SendMail)
	err = http.ListenAndServe(":"+cfg.HTTPPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
