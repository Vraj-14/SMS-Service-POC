package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vraj/notification-service/internal/batch"
)

type BulkHandler struct {
	bulkService *batch.BulkService
}

func NewBulkHandler(b *batch.BulkService) *BulkHandler {
	return &BulkHandler{bulkService: b}
}

func (h *BulkHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/send-bulk-sms", h.SendBulkSMS).Methods("POST")
}

type BulkRequest struct {
	Message string `json:"message"`
}

func (h *BulkHandler) SendBulkSMS(w http.ResponseWriter, r *http.Request) {

	var req BulkRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		http.Error(w, "message is required", http.StatusBadRequest)
		return
	}

	filePath := "data/numbers.xlsx"

	response, err := h.bulkService.SendBulkSMS(filePath, req.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
