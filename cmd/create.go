package cmd

import (
	"github.com/spf13/cobra"
)

var (
	amount      int
	description string
	email       string
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().IntVarP(&amount, "amount", "a", 0, "amount you're requesting for")
	createCmd.Flags().StringVarP(&email, "email", "e", "", "customer's email")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "[optional] a description of the request")
	createCmd.MarkFlagRequired("amount")
	createCmd.MarkFlagRequired("email")
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a paystack payment request",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		// fetch customer_id (from email)
		// 	true: (200) get customer_code
		//	false: (400/404) create new customer and get code

		// create payment request (amount + customer_id + description)
		// 	true: (200) fetch payment_link(paystack.com/pay/{request_code}); display success message
		// false: display error
	},
}
