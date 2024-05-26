package cmd

import (
	"errors"
	"fmt"
	"payme/internal/adapters"
	"payme/internal/models"
	"strconv"

	"github.com/spf13/cobra"
)

var silent bool

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.PersistentFlags().BoolVarP(&silent, "silent", "s", false, "create request without notifying customer")
}

var createCmd = &cobra.Command{
	Use:   "create [amount] [customer email] [description]",
	Short: "Create a paystack payment request",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("'create' requires an amount, the customer's email, & a description")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		adapter := adapters.NewPaystackAPIAdapter()
		amount, _ = strconv.ParseFloat(args[0], 64)
		email = args[1]
		description = args[2]

		// validate customer
		// create new if not exists
		var customer_id string

		customer, err := adapter.FetchCustomer(email)
		if err != nil {
			new_customer, err := adapter.CreateCustomer(models.Customer{Email: email})
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			customer_id = new_customer.Data.CustomerCode
		} else {
			customer_id = customer.Data.CustomerCode
		}

		// create payment request
		paymentRequest := models.PaymentRequest{
			Amount:           amount * 100,
			Description:      description,
			Customer:         customer_id,
			SendNotification: !silent,
		}

		createdPaymentRequest, err := adapter.CreatePaymentRequest(paymentRequest)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		data := createdPaymentRequest.Data

		fmt.Println("Payment request created successfully!")
		fmt.Println("-------------------------------------------")
		fmt.Println("Payment Request Details")
		fmt.Println("-------------------------------------------")
		fmt.Println("Date:", data.CreatedAt)
		fmt.Println("Amount:", data.Currency, data.Amount/100)
		if len(data.Description) > 0 {
			fmt.Println("Description:", data.Description)
		}
		fmt.Println("Status:", data.Status)
		fmt.Println("Paid:", data.Paid)
		fmt.Println("Payment link:", fmt.Sprintf("%s/%s", "https://paystack.com/pay", data.RequestCode))
	},
}
