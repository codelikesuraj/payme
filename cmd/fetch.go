package cmd

import (
	"errors"
	"fmt"
	"payme/internal/adapters"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fetchCmd)
}

var fetchCmd = &cobra.Command{
	Use:   "fetch [request code/ID]",
	Short: "Fetch a paystack payment request",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("'fetch' requires a request code/ID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		paymentRequest, err := adapters.NewPaystackAPIAdapter().FetchPaymentRequest(args[0])
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		data := paymentRequest.Data

		fmt.Println("-------------------------------------------")
		fmt.Println("Payment Request Details")
		fmt.Println("-------------------------------------------")
		fmt.Printf("Request ID: %d\n", data.ID)
		fmt.Println("Request Code:", data.RequestCode)
		fmt.Println("Date:", data.CreatedAt)
		fmt.Println("Amount:", data.Currency, data.Amount/100)
		if len(data.Description) > 0 {
			fmt.Println("Description:", data.Description)
		}
		fmt.Printf("Customer: %s (%s)\n", data.Customer.CustomerCode, data.Customer.Email)
		fmt.Println("Status:", data.Status)
		fmt.Println("Paid:", data.Paid)
		link := fmt.Sprintf("Payment link: %s/%s", "https://paystack.com/pay", data.RequestCode)
		fmt.Println(link)
		fmt.Println(strings.Repeat("-", len(link)+7))
	},
}
