package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vraj/notification-service/internal/model"
	"github.com/vraj/notification-service/internal/service"

)


type SMSHandler struct {
	service *service.SMSService
}

func NewSMSHandler(s *service.SMSService) *SMSHandler {
	return &SMSHandler{service: s}
}

func (h *SMSHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/send-sms", h.SendSMS).Methods("POST")
}

func (h *SMSHandler) SendSMS(w http.ResponseWriter, r *http.Request) {

	var req model.SMSRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.service.SendSMS(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
