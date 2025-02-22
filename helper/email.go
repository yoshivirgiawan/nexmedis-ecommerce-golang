package helper

import (
	"fmt"
	"net/smtp"
	"os"
)

// SendEmail mengirim email dengan format HTML menggunakan SMTP
func SendEmail(to, subject, htmlBody string) error {
	// Ambil konfigurasi dari file .env
	from := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT")
	fromName := os.Getenv("MAIL_FROM_NAME")

	// Autentikasi untuk koneksi SMTP
	auth := smtp.PlainAuth("", from, password, host)

	// Membuat pesan email dengan format HTML
	toEmails := []string{to}
	msg := []byte("From: " + fromName + "<" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		htmlBody + "\r\n")

	// Kirim email
	err := smtp.SendMail(host+":"+port, auth, from, toEmails, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}
	return nil
}
