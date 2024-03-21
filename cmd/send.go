/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/reneval/rcom/domain"
	"github.com/reneval/rcom/notification"
	"github.com/reneval/rcom/services"
	"github.com/spf13/cobra"
)

var (
	target  *string
	message *string
	channel *string
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Will send a message to the desired target using desired channel",
	Long: `This command will send a message to the desired target using the desired channel.
Example: rcom send -t foo@bar.com -m "Hello World" -c email
The supported channels are 
- email
- sms
- slack.`,
	Run: func(cmd *cobra.Command, args []string) {

		emailSender := notification.NewEmailSender()
		smsSender := notification.NewSMSSender()
		slackSender := notification.NewSlackSender()

		notificationService := services.NewNotificationService(emailSender, smsSender, slackSender)

		message := domain.Message{
			Target:  *target,
			Body: *message,
			Channel: *channel,
		}

		notificationService.Send(message)

	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	target = sendCmd.Flags().StringP("target", "t", "", "The target of the message")
	message = sendCmd.Flags().StringP("message", "m", "", "The message to send")
	channel = sendCmd.Flags().StringP("channel", "c", "", "The channel to send the message to")

}
