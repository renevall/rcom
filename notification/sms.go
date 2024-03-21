package notification

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/reneval/rcom/domain"
)

// SmeagolSMS is a theorical external service that would be used to send SMS
type SmeagolSMS struct {
}

func (s *SmeagolSMS) SendSMS(m domain.Message) {

}

type SMSSender struct {
	client SmeagolSMS
}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

// Send will send an email to the target using X email service
func (e *SMSSender) Send(m domain.Message) (string, error) {
	e.client.SendSMS(m)

	fmt.Printf("Sending sms to %s with message %s\n", m.Target, m.Body)
	id := e.getID()
	return id, nil
}

func (e *SMSSender) getID() string {
	// here we would get the tracking id from the email service
	// for the purpose of this test, we woudd generate a uuid
	return uuid.New().String()
}
