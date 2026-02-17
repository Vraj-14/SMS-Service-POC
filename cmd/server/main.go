package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/vraj/notification-service/internal/batch"
	"github.com/vraj/notification-service/internal/handler"
	"github.com/vraj/notification-service/internal/provider"
	"github.com/vraj/notification-service/internal/service"
	"github.com/vraj/notification-service/internal/validators"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	port := os.Getenv("PORT")

	twilioProvider := provider.NewTwilioProvider(accountSid, authToken, fromNumber)

	smsValidator := validators.NewSMSValidator()

	smsService := service.NewSMSService(twilioProvider, smsValidator)

	excelReader := batch.NewExcelReader()

	bulkService := batch.NewBulkService(excelReader, smsService)


	smsHandler := handler.NewSMSHandler(smsService)
	bulkHandler := handler.NewBulkHandler(bulkService)

	router := mux.NewRouter()

	smsHandler.RegisterRoutes(router)
	bulkHandler.RegisterRoutes(router)

	// Optional health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Notification service is running"))
	}).Methods("GET")

	log.Println("Server running on port", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
