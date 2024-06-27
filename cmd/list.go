package cmd

import (
	"fmt"
	"os"
	"payme/internal/adapters/paystack"
	"payme/internal/domain"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var flag domain.ListFlag

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().IntVarP(&flag.Count, "count", "c", 50, "amount per list (max. 50)")
	listCmd.PersistentFlags().BoolVarP(&flag.Last, "last", "l", false, "fetch only the last payment request")
	listCmd.PersistentFlags().IntVarP(&flag.Page, "page", "p", 1, "page for list payment request")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all paystack payment requests",
	Run: func(cmd *cobra.Command, args []string) {
		paymentRequests, err := paystack.NewPaymentRequestAdapter(os.Getenv("PAYSTACK_SK")).ListPaymentRequest(flag)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("-------------------------------------------")
		if !flag.Last {
			fmt.Println("Payment Request Lists")
		} else {
			fmt.Println("Payment Request Details")
		}
		fmt.Println("-------------------------------------------")
		meta := paymentRequests.Meta

		rawPerPage := meta.PerPage
		var perPage int
		switch reflect.TypeOf(rawPerPage).String() {
		case "string":
			perPage, _ = strconv.Atoi(rawPerPage.(string))
		case "float64":
			perPage = int(rawPerPage.(float64))
		}

		if len(paymentRequests.Data) < 1 {
			fmt.Printf("No payment requests found!\n\n")
			return
		}

		for i, paymentRequest := range paymentRequests.Data {
			if !flag.Last {
				fmt.Printf("(%d)\n", (meta.Page*perPage)-perPage+i+1)
			}
			fmt.Printf("Request ID: %d\n", paymentRequest.ID)
			fmt.Println("Request Code:", paymentRequest.RequestCode)
			fmt.Println("Date:", paymentRequest.CreatedAt)
			fmt.Println("Amount:", paymentRequest.Currency, paymentRequest.Amount/100)
			if len(paymentRequest.Description) > 0 {
				fmt.Println("Description:", paymentRequest.Description)
			}
			fmt.Printf("Customer: %s (%s)\n", paymentRequest.Customer.CustomerCode, paymentRequest.Customer.Email)
			fmt.Println("Status:", paymentRequest.Status)
			fmt.Println("Paid:", paymentRequest.Paid)
			link := fmt.Sprintf("Payment link: %s/%s", "https://paystack.com/pay", paymentRequest.RequestCode)
			if !paymentRequest.Paid {
				fmt.Println(link)
			}
			fmt.Println(strings.Repeat("-", len(link)+7))
		}
	},
}
