package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
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
//
// The request is pre-validated to avoid unnecessary Bad Request
func (s Service) Create(request DataRequest) (Single, error) {

	if err := validate(request.Data.Attributes); err != nil {
		return Single{}, err
	}

	req := new(bytes.Buffer)
	json.NewEncoder(req).Encode(request)
	resp, err := httpClient.Post(s.URL+"/v1/organisation/accounts", "application/json", req)

	if err != nil {
		println("ERROR: ")
	}

	if resp.StatusCode != 201 {
		var result ErrorResponse
		json.NewDecoder(resp.Body).Decode(&result)

		return Single{}, errors.New(result.ErrorMessage)
	}

	var result Single
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil

}

// Fetch an account
func (s Service) Fetch(id string) (Single, error) {
	accountNumber := "123"
	firstName := "razvan"
	account := Account{AccountNumber: &accountNumber, FirstName: &firstName}
	data := Data{Attributes: account}
	return Single{Data: data}, nil
}

// List accounts
func (s Service) List() (List, error) {
	return List{}, nil
}

// Delete an account
func (s Service) Delete(id string) (bool, error) {
	return true, nil
}

func validate(account Account) error {
	var validBIC = regexp.MustCompile(`^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$`)
	if account.BIC != nil && validBIC.MatchString(*account.BIC) == false {
		return fmt.Errorf("Invalid BIC [%s]", *account.BIC)
	}

	var validAccountClassification = regexp.MustCompile(`^(Personal|Business)$`)
	if account.AccountClassification != nil && validAccountClassification.MatchString(*account.AccountClassification) == false {
		return fmt.Errorf("Invalid AccountClassification [%s]", *account.AccountClassification)
	}

	var validBankID = regexp.MustCompile(`^[A-Z0-9]{0,16}$`)
	if account.BankID != nil && validBankID.MatchString(*account.BankID) == false {
		return fmt.Errorf("Invalid BankID [%s]", *account.BankID)
	}

	var validBaseCurrency = regexp.MustCompile(`^[A-Z]{3}$`)
	if account.BaseCurrency != nil && validBaseCurrency.MatchString(*account.BaseCurrency) == false {
		return fmt.Errorf("Invalid BaseCurrency [%s]", *account.BaseCurrency)
	}

	var validCountry = regexp.MustCompile(`^[A-Z]{2}$`)
	if validCountry.MatchString(account.Country) == false {
		return fmt.Errorf("Invalid Country [%s]", account.Country)
	}

	if account.AlternativeBankAccountNames != nil && len(*account.AlternativeBankAccountNames) > 3 {
		return fmt.Errorf("Invalid AlternativeBankAccountNames %s", *account.AlternativeBankAccountNames)
	}

	return nil
}
