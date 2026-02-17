package validator

import (
	"errors"
	"regexp"
	"strings"

	"github.com/vraj/notification-service/internal/model"
)

type SMSValidator struct{}

func NewSMSValidator() *SMSValidator {
	return &SMSValidator{}
}

func (vx`` *SMSValidator) Validate(req model.SMSRequest) error {

	// Trim whitespace
	req.To = strings.TrimSpace(req.To)
	req.Message = strings.TrimSpace(req.Message)

	// 1️⃣ Mandatory fields
	if req.To == "" {
		return errors.New("phone number is required")
	}

	if req.Message == "" {
		return errors.New("message cannot be empty")
	}

	// 2️⃣ E.164 Phone Validation
	e164Regex := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	if !e164Regex.MatchString(req.To) {
		return errors.New("invalid phone number format (must be E.164 format)")
	}

	// 3️⃣ Message length validation
	if len(req.Message) > 160 {
		return errors.New("message exceeds 160 characters limit")
	}

	return nil
}
