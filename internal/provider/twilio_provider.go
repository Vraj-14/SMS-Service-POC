package provider

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioProvider struct{
	client *twilio.RestClient
	fromNumber string
}

func NewTwilioProvider(accountSid, authToken, fromNumber string) *TwilioProvider {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return &TwilioProvider{
		client:     client,
		fromNumber: fromNumber,
	}
}

func (t *TwilioProvider) SendSMS(to string, message string) error{

	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(t.fromNumber)
	params.SetBody(message)


	_, err := t.client.Api.CreateMessage(params)

	if err != nil {
		return fmt.Errorf("failed to send SMS: %w", err)
	}

	return nil

}