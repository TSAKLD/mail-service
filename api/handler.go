package api

import (
	"encoding/json"
	"mail.project/entity"
	"mail.project/service"
	"net/http"
)

type Handler struct {
	us *service.MailService
}

func (h *Handler) SendMail(w http.ResponseWriter, r *http.Request) {
	var mail entity.Mail

	err := json.NewDecoder(r.Body).Decode(&mail)
	if err != nil {
		sendError(w, err)
		return
	}

	err = h.us.SendMail(mail)
	if err != nil {
		sendError(w, err)
		return
	}
}

func NewHandler(userService *service.MailService) *Handler {
	return &Handler{us: userService}
}
