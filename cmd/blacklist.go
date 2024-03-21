/*
Copyright © 2024 NAME HERE <reneval@gmail.com>
*/
package cmd

import (
	"github.com/reneval/rcom/postgres"
	"github.com/reneval/rcom/services"
	"github.com/spf13/cobra"
)

var (
	path      *string
	mode      *string
	batchSize *int
	startLine *int
)

// blacklistCmd represents the blacklist command
var blacklistCmd = &cobra.Command{
	Use:   "blacklist",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		postgresClient := postgres.NewClient()
		cleanService := services.NewEmailCleanerService(postgresClient)

		cleanService.CleanByFile(*path, *mode, *batchSize, *startLine)
	},
}

func init() {
	rootCmd.AddCommand(blacklistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// blacklistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blacklistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	path = blacklistCmd.Flags().StringP("path", "p", "./data/blacklist.txt", "The path of the file to clean")
	mode = blacklistCmd.Flags().StringP("mode", "m", "sequencial", "The mode of the blacklist command")
	batchSize = blacklistCmd.Flags().IntP("batch", "b", 2000, "The batch size of the blacklist command")
	startLine = blacklistCmd.Flags().IntP("line", "l", 0, "The start line of the blacklist command")
}
