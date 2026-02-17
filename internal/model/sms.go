package model

type SMSRequest struct{
	To string `json:"to"`
	Message string `json:"message"`
}

type SMSResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

