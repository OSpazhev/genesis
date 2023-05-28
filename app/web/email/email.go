package email

import (
	"fmt"
	"github.com/ospazhev/genesis/app/db"
	"log"
	"net/smtp"
	"os"
	"strings"
)

func SendRate(rate float64) error {
	from := os.Getenv("SENDER_EMAIL")
	password := os.Getenv("SENDER_PASSWORD")

	to := db.Db.Subscribers()

	smtpHost := os.Getenv("SMTP_HOST")
	if len(smtpHost) == 0 {
		smtpHost = "sandbox.smtp.mailtrap.io"
	}
	smtpPort := os.Getenv("SMTP_PORT")
	if len(smtpPort) == 0 {
		smtpPort = "2525"
	}

	body := fmt.Sprintf("%f", rate)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := "From: " + from + "\n" +
		"To: " + strings.Join(to, ";") + "\n" +
		"Subject: BTC to UAH rate\n\n" +
		body

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		log.Println("Email sending failed! Please check email settings and credentials")
		return err
	}
	log.Println("Email Sent Successfully!")
	return nil
}
