# Payme: Manage your paystack payment request from the CLI
Payme is an easy-to-use CLI tool written in Go that allows you to manage your paystack payment requests.
It allows you to send, list, update and also verify the status of your payment requests.

## Usage
go run main.go [command] [args...] [flags...]

## Setup
- Go to your [Paystack dashboard](https:dashboard.paystack.com)
- Navigate to the settings page
- Click on the API & Webhook tab 
- Then copy your Secret API key
- Add this to your machine ENV as "PAYSTACK_SK"
- You "SHOULD" be up and running 👍