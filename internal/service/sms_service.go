package service

import (
	

	"github.com/vraj/notification-service/internal/model"
	"github.com/vraj/notification-service/internal/provider"
	"github.com/vraj/notification-service/internal/validators"


)


type SMSService struct {
	provider *provider.TwilioProvider
	validator *validators.SMSValidator
}

func NewSMSService(p *provider.TwilioProvider, v *validators.SMSValidator) *SMSService {
	return &SMSService{
		provider: p,
		validator: v,
	}
}

func (s *SMSService)  SendSMS(req model.SMSRequest) (*model.SMSResponse, error) {

	// if req.To == "" || req.Message == "" {
	// 	return nil, errors.New("to and message are required")
	// }

	if err := s.validator.Validate(req); err != nil{
		return nil, err
	}

	err := s.provider.SendSMS(req.To, req.Message)
	if err != nil {
		return nil, err
	}

	return &model.SMSResponse{
		Status:  "SUCCESS",
		Message: "SMS sent successfully",
	}, nil
}
