package utils

import (
	"os"
	"strconv"
	"suranect_api/utils/email"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func SendVerifyEmail(to_email string, token_verify int) error {
	godotenv.Load(".env")

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv("EMAIL_NAME"))
	mailer.SetHeader("To", to_email)
	mailer.SetHeader("Subject", "Verify Email")
	mailer.SetBody("text/html", email.GetTemplateVerifyEmail(token_verify))

	port, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))

	dialer := gomail.NewDialer(
		os.Getenv("EMAIL_HOSTNAME"),
		port,
		os.Getenv("EMAIL_NAME"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	err := dialer.DialAndSend(mailer)

	if err != nil {
		return err
	} else {
		return nil
	}
}
