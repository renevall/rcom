package services

import (
	"errors"
	"log/slog"

	"github.com/reneval/rcom/domain"
)

// NotificationService is a service that will send a message to a target using a given channel
type NotificationService struct {
	Email Notifier
	SMS   Notifier
	Slack Notifier
}

//  Note: To make it even more extendandable we could use a map of Notifiers instead of the Email, SMS and Slack fields.

// Notifier wraps the Send method to a given Notification channel
type Notifier interface {
	Send(m domain.Message) (string, error)
}

// NewNotificationService will create a new notification service
func NewNotificationService(email, sms, slack Notifier) *NotificationService {
	return &NotificationService{
		Email: email,
		SMS:   sms,
		Slack: slack,
	}
}

// Send will send a message to the target using the desired channel
func (n *NotificationService) Send(m domain.Message) error {

	var err error
	var id string
	switch m.Channel {

	// Here is where we see the adapter pattern in action. We are using the same method to send a message
	// to different channels. The implementation of the Send method is different for each channel. For
	// example the authentication is different on each channel, the way we send the message is different.
	// For email we add a tracking id/pixel/unique link to the email, for sms we just send the message and
	case "email":
		id, err = n.Email.Send(m)
	case "sms":
		id, err = n.SMS.Send(m)
	case "slack":
		id, err = n.Slack.Send(m)
	default:
		err = errors.New("channel not supported")
	}
	if err != nil {
		slog.Error("error sending message", "error", err)
		return err
	}

	slog.Info("message sent", "id", id, "target", m.Target, "channel", m.Channel)

	return nil
}
