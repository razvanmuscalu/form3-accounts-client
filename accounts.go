package main

import (
	"net/http"
	"time"
)

// AccountsClient is the client interface to access the Accounts API
type AccountsClient interface {
	Create(request DataRequest) (Single, error)
	Fetch(id string) (Single, error)
	List() (List, error)
	Delete(id string) (bool, error)
}

// AccountsService implements AccountsClient
type AccountsService struct {
	URL string
}

var _ AccountsClient = (*AccountsService)(nil)

var httpClient = http.Client{
	Timeout: time.Second * 2,
}

// Create an account
func (as AccountsService) Create(request DataRequest) (Single, error) {
	return Single{}, nil
}

// Fetch an account
func (as AccountsService) Fetch(id string) (Single, error) {
	accountNumber := "123"
	account := Account{AccountNumber: &accountNumber}
	data := Data{Attributes: account}
	return Single{Data: data}, nil
}

// List accounts
func (as AccountsService) List() (List, error) {
	return List{}, nil
}

// Delete an account
func (as AccountsService) Delete(id string) (bool, error) {
	return true, nil
}
