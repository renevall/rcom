package notification

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/reneval/rcom/domain"
	"github.com/slack-go/slack"
)

type SlackSender struct {
	client slack.Client
}

func NewSlackSender() *SlackSender {
	return &SlackSender{}
}

// Send will send an email to the target using X email service
func (e *SlackSender) Send(m domain.Message) (string, error) {
	fmt.Printf("Sending slack message to %s with message %s\n", m.Target, m.Body)

	// e.client.SendMessage(m.Target, m.Body)
	id := e.getID()
	return id, nil
}

func (e *SlackSender) getID() string {
	// here we would get slack tracking id from the slack service
	// for the purpose of this test, we woudd generate a uuid
	return uuid.New().String()
}

// validate will validate the message. We expect the target to be in the form of organization/channel
func (e *SlackSender) validate(m domain.Message) error {
	if len(strings.Split(m.Target, "/")) < 2 {
		return fmt.Errorf("target should be in the form of organization/channel")
	}

	if m.Body == "" {
		return fmt.Errorf("body is required")
	}
	return nil
}
