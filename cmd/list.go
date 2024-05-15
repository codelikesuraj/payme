package cmd

import (
	"fmt"
	"payme/internal/adapters"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all paystack payment requests",
	Run: func(cmd *cobra.Command, args []string) {
		adapter := adapters.NewPaystackAPIAdapter()

		// list payment requests
		paymentRequests, err := adapter.ListPaymentRequest()
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("-------------------------------------------")
		fmt.Println("Payment Request Lists")
		fmt.Println("-------------------------------------------")
		meta := paymentRequests.Meta
		for i, paymentRequest := range paymentRequests.Data {
			fmt.Printf("%d.\tRequest ID: %d\n", (meta.Page*meta.PerPage)-meta.PerPage+i+1, paymentRequest.ID)
			fmt.Println("\tRequest Code:", paymentRequest.RequestCode)
			fmt.Println("\tDate:", paymentRequest.CreatedAt)
			fmt.Println("\tAmount:", paymentRequest.Currency, paymentRequest.Amount/100)
			if len(paymentRequest.Description) > 0 {
				fmt.Println("\tDescription:", paymentRequest.Description)
			}
			fmt.Println("\tStatus:", paymentRequest.Status)
			fmt.Println("\tPaid:", paymentRequest.Paid)
			link := fmt.Sprintf("\tPayment link: %s/%s", "https://paystack.com/pay", paymentRequest.RequestCode)
			fmt.Println(link)
			fmt.Println(strings.Repeat("-", len(link) + 7))
		}
	},
}
