package validators

import(
	"errors"
	"regexp"
	"strings"

	"github.com/vraj/notification-service/internal/model"
)

type SMSValidator struct{}

func NewSMSValidator() *SMSValidator{
	return &SMSValidator{}
}

func (v *SMSValidator) Validate(req model.SMSRequest) error {

	// trim whitespace
	req.To = strings.TrimSpace(req.To)
	req.Message = strings.TrimSpace(req.Message)

	// Mandatory fields
	if req.To == "" {
		return errors.New("Phone number cannot be empty.")
	}

	if req.Message == "" {
		return errors.New("Message cannot be empty.")
	}

	// Phone Number validation
	e164Regex := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)

	if !e164Regex.MatchString(req.To) {
		return errors.New("Invlid phone number format (must be E.164 format)")
	}

	//Messge lenght validation
	if len(req.Message) > 160 {
		return errors.New("Message length exceeds 160 characters.")
	}

	return nil
}