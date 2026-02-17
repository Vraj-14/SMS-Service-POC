package batch

import (
	"github.com/vraj/notification-service/internal/model"
	"github.com/vraj/notification-service/internal/service"
)

type BulkService struct {
	reader     *ExcelReader
	smsService *service.SMSService
}

func NewBulkService(reader *ExcelReader, smsService *service.SMSService) *BulkService {
	return &BulkService{
		reader:     reader,
		smsService: smsService,
	}
}

type BulkResponse struct {
	Total   int `json:"total"`
	Success int `json:"success"`
	Failed  int `json:"failed"`
}

func (b *BulkService) SendBulkSMS(filePath string, message string) (*BulkResponse, error) {

	numbers, err := b.reader.ReadPhoneNumbers(filePath)
	if err != nil {
		return nil, err
	}

	total := len(numbers)
	success := 0
	failed := 0

	for _, number := range numbers {

		req := model.SMSRequest{
			To:      number,
			Message: message,
		}

		_, err := b.smsService.SendSMS(req)
		if err != nil {
			failed++
			continue
		}

		success++
	}

	return &BulkResponse{
		Total:   total,
		Success: success,
		Failed:  failed,
	}, nil
}
