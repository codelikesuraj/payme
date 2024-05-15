/*
Copyright © 2024 Abdulbaki Suraj <codelikesuraj@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	amount      float64
	description string
	email       string
)

var rootCmd = &cobra.Command{
	Use:   "payme",
	Short: "Manage your paystack payment request from the CLI",
	Long: `Payme is an easy-to-use CLI tool that allows 
you to manage your paystack payment requests.
It allows you to send, list, update and 
also verify the status of your payment requests`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
