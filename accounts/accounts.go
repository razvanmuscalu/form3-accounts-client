package accounts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// Client is the client interface to access the Accounts API
type Client interface {
	Create(request DataRequest) (Single, error)
	Fetch(id string) (Single, error)
	List() (List, error)
	Delete(id string) (bool, error)
}

// Service implements Client
type Service struct {
	URL string
}

var (
	httpClient http.Client = http.Client{
		Timeout: time.Second * 2,
	}
)

// Create an account
func (as Service) Create(request DataRequest) (Single, error) {

	req := new(bytes.Buffer)
	json.NewEncoder(req).Encode(request)
	resp, err := httpClient.Post(as.URL+"/v1/organisation/accounts", "application/json", req)

	if err != nil {
		println("ERROR: ")
	}

	var result Single
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// Fetch an account
func (as Service) Fetch(id string) (Single, error) {
	accountNumber := "123"
	firstName := "razvan"
	account := Account{AccountNumber: &accountNumber, FirstName: &firstName}
	data := Data{Attributes: account}
	return Single{Data: data}, nil
}

// List accounts
func (as Service) List() (List, error) {
	return List{}, nil
}

// Delete an account
func (as Service) Delete(id string) (bool, error) {
	return true, nil
}
