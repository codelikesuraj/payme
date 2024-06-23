# Payme: Manage your paystack payment request from the CLI
Payme is an easy-to-use CLI tool written in Go that allows you to manage your paystack payment requests.
It allows you to send, list, update and also verify the status of your payment requests.

## Setup
- Go to your [Paystack dashboard](https:dashboard.paystack.com)
- Navigate to the settings page
- Click on the API & Webhook tab 
- Then copy your Secret API key
- Add this to your machine ENV as "PAYSTACK_SK"
- You "SHOULD" be up and running 👍

## Usage
payme [command] [args...] [flags...]

## Available commands
|Command|Description|Args|Flags [optional]|
|-|-|-|-|
|create|create a paystack payment request|{amount} {customer email} {description}|-s,--silent: do not send email notification|
|list|list all paystack payment requests|none|-c, --count {count}: no. of payment requests per list/page (max. 50)|
||||-l, --last: only the last payment request|
||||-p, --page {page}: page for payment requests list (default 1)|
|fetch|fetch a paystack payment request|{request code/request ID}|-|
