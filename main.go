package main

import (
	"github.com/razvanmuscalu/form3-accounts-client/accounts"
)

func main() {

	AccountsClient := accounts.Service{URL: ""}

	resp, _ := AccountsClient.Fetch("c24e60c2-9dcb-466a-8fc4-fdf19c630b29")

	AccountsClient.List(nil, nil)

	println(*resp.Data.Attributes.AccountNumber)
}
