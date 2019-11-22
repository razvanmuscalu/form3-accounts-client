package main

func main() {

	AccountsClient := AccountsService{URL: ""}

	resp, _ := AccountsClient.Fetch("c24e60c2-9dcb-466a-8fc4-fdf19c630b29")

	println(*resp.Data.Attributes.AccountNumber)
}
