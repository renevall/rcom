package services

import (
	"testing"

	"github.com/reneval/rcom/domain"
)

// Note: In reality I would use testity and mockery libraries to generate the mocks

type EmailNotifierMock struct {
	Called bool
}

func (e *EmailNotifierMock) Send(m domain.Message) (string, error) {
	e.Called = true
	return "", nil
}

type SMSNotifierMock struct {
	Called bool
}

func (s *SMSNotifierMock) Send(m domain.Message) (string, error) {
	s.Called = true
	return "", nil
}

type SlackNotifierMock struct {
	Called bool
}

func (s *SlackNotifierMock) Send(m domain.Message) (string, error) {
	s.Called = true
	return "", nil
}

func TestNotificationService_Send(t *testing.T) {
	type fields struct {
		Email Notifier
		SMS   Notifier
		Slack Notifier
	}
	type args struct {
		m domain.Message
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		validate func(t *testing.T, service *NotificationService)
		wantErr  bool
	}{
		{
			name: "when sending an email, the email notifier should be used",
			fields: fields{
				Email: &EmailNotifierMock{},
			},
			args: args{
				m: domain.Message{
					Channel: "email",
				},
			},
			validate: func(t *testing.T, service *NotificationService) {
				if !service.Email.(*EmailNotifierMock).Called {
					t.Errorf("expected email notifier to be called")
				}
			},
			wantErr: false,
		},
		{
			name: "when sending an sms, the sms notifier should be used",
			fields: fields{
				SMS: &SMSNotifierMock{},
			},
			args: args{
				m: domain.Message{
					Channel: "sms",
				},
			},
			validate: func(t *testing.T, service *NotificationService) {
				if !service.SMS.(*SMSNotifierMock).Called {
					t.Errorf("expected sms notifier to be called")
				}
			},
			wantErr: false,
		},
		{
			name: "when sending a slack message, the slack notifier should be used",
			fields: fields{
				Slack: &SlackNotifierMock{},
			},
			args: args{
				m: domain.Message{
					Channel: "slack",
				},
			},
			validate: func(t *testing.T, service *NotificationService) {
				if !service.Slack.(*SlackNotifierMock).Called {
					t.Errorf("expected slack notifier to be called")
				}
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NotificationService{
				Email: tt.fields.Email,
				SMS:   tt.fields.SMS,
				Slack: tt.fields.Slack,
			}
			err := n.Send(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotificationService.Send() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.validate != nil {
				tt.validate(t, n)
			}
		})
	}
}
