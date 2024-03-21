package notification

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/reneval/rcom/domain"
)

type EmailSender struct {
	Account string
	Secret  string
	URL     string
}

func NewEmailSender() *EmailSender {
	// here we would initialize the email sender with the necessary credentials
	// for the purpose of this test, we will just return an empty struct
	return &EmailSender{}
}

// Send will send an email to the target using X email service
func (e *EmailSender) Send(m domain.Message) (string, error) {
	fmt.Printf("Sending email to %s with message %s\n", m.Target, m.Body)
	id := e.getID()
	return id, nil
}

func (e *EmailSender) getID() string {
	// here we would get the tracking id from the email service
	// for the purpose of this test, we woudd generate a uuid
	return uuid.New().String()
}


// We could add a validat method to validate the message before sending it. For example
// we could check if the target is a valid email address and if the body is not empty.