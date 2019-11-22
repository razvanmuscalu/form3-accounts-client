package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	AccountsClient := AccountsService{URL: ""}

	resp, _ := AccountsClient.Fetch("c24e60c2-9dcb-466a-8fc4-fdf19c630b29")
	assert.Equal(t, "123", *resp.Data.Attributes.AccountNumber)
}
