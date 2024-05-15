package cmd

import (
	"fmt"
	"payme/internal/adapters"
	"strings"

	"github.com/spf13/cobra"
)

var recentFlag bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolVar(&recentFlag, "recent", false, "get the most recent payment request")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all paystack payment requests",
	Run: func(cmd *cobra.Command, args []string) {
		adapter := adapters.NewPaystackAPIAdapter()

		// list payment requests
		paymentRequests, err := adapter.ListPaymentRequest(recentFlag)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("-------------------------------------------")
		if !recentFlag {
			fmt.Println("Payment Request Lists")
		} else {
			fmt.Println("Payment Request Details")
		}
		fmt.Println("-------------------------------------------")
		meta := paymentRequests.Meta

		perPage := meta.PerPage
		for i, paymentRequest := range paymentRequests.Data {
			if !recentFlag {
				fmt.Printf("(%d)\n", (meta.Page*int(perPage.(float64)))-int(perPage.(float64))+i+1)
			}
			fmt.Printf("Request ID: %d\n", paymentRequest.ID)
			fmt.Println("Request Code:", paymentRequest.RequestCode)
			fmt.Println("Date:", paymentRequest.CreatedAt)
			fmt.Println("Amount:", paymentRequest.Currency, paymentRequest.Amount/100)
			if len(paymentRequest.Description) > 0 {
				fmt.Println("Description:", paymentRequest.Description)
			}
			fmt.Println("Status:", paymentRequest.Status)
			fmt.Println("Paid:", paymentRequest.Paid)
			link := fmt.Sprintf("Payment link: %s/%s", "https://paystack.com/pay", paymentRequest.RequestCode)
			fmt.Println(link)
			fmt.Println(strings.Repeat("-", len(link)+7))
		}
	},
}
