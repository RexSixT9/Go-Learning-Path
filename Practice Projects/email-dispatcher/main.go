package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChannel := make(chan Recipient)

	go func() {
		loadRecipients("./emails.csv", recipientChannel)
	}()

	var wg sync.WaitGroup
	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}
	wg.Wait()
}

func executeEmailTemplate(recipient Recipient) (string, error) {
	template, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = template.Execute(&buf, recipient)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
