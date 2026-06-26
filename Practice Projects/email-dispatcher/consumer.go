package main

import (
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\nHello %s,\nThis is a test email.\n", recipient.Email, recipient.Name)
		// msg := []byte(formattedMsg)

		msg, err := executeEmailTemplate(recipient)
		if err != nil {
			log.Fatal("Worker", id, "failed to execute email template for:", recipient.Name, "<"+recipient.Email+">", "Error:", err)
			continue
		}

		log.Printf("Worker %d: Sending email to %s", id, recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "test@example.com", []string{recipient.Email}, []byte(msg))
		if err != nil {
			log.Fatal("Worker", id, "failed to send email to:", recipient.Name, "<"+recipient.Email+">", "Error:", err)
		}
		time.Sleep(50 * time.Millisecond)
		log.Printf("Worker %d: Sent email to %s", id, recipient.Email)
	}
}
