package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Client is the client interface to access the Accounts API
type Client interface {
	Create(request DataRequest) (Single, error)
	Fetch(id uuid.UUID) (Single, error)
	List(page *Page, filter *Filter) (List, error)
	Delete(id uuid.UUID, version int) (bool, error)
}

// Service implements Client
type Service struct {
	URL string
}

var httpClient http.Client = http.Client{
	Timeout: time.Second * 2,
}

const path = "/v1/organisation/accounts"

// Create an account
//
// The request is pre-validated to avoid unnecessary Bad Request
func (s Service) Create(request DataRequest) (Single, error) {

	if err := validateAccount(request.Data.Attributes); err != nil {
		return Single{}, err
	}

	req := new(bytes.Buffer)
	json.NewEncoder(req).Encode(request)
	resp, err := httpClient.Post(fmt.Sprintf("%s%s", s.URL, path), "application/json", req)
	if err != nil {
		return Single{}, fmt.Errorf("An error has occured while creating account")
	}

	if err := decodeErrorResponse(resp); err != nil {
		return Single{}, err
	}

	var result Single
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil

}

// Fetch an account
func (s Service) Fetch(id uuid.UUID) (Single, error) {

	resp, err := httpClient.Get(fmt.Sprintf("%s%s/%s", s.URL, path, id))
	if err != nil {
		return Single{}, fmt.Errorf("An error has occured while fetching account")
	}

	if err := decodeErrorResponse(resp); err != nil {
		return Single{}, err
	}

	var result Single
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil

}

// List accounts
func (s Service) List(page *Page, filter *Filter) (List, error) {

	resp, err := httpClient.Get(buildListURL(s.URL, page, filter))
	if err != nil {
		return List{}, fmt.Errorf("An error has occured while listing accounts")
	}

	if err := decodeErrorResponse(resp); err != nil {
		return List{}, err
	}

	var result List
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// Delete an account
func (s Service) Delete(id uuid.UUID, version int) (bool, error) {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/%s?version=%s", s.URL, path, id, strconv.Itoa(version)), new(bytes.Buffer))
	if err != nil {
		return false, fmt.Errorf("An error has occured while constructing delete request")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("An error has occured while deleting account")
	}

	if err := decodeErrorResponse(resp); err != nil {
		return false, err
	}

	return true, nil
}

func decodeErrorResponse(resp *http.Response) error {
	if !(resp.StatusCode == 200 || resp.StatusCode == 201 || resp.StatusCode == 204) {
		var result ErrorResponse
		json.NewDecoder(resp.Body).Decode(&result)

		return errors.New(result.ErrorMessage)
	}

	return nil
}

func validateAccount(account Account) error {
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

func buildListURL(baseURL string, page *Page, filter *Filter) string {

	var params []string
	if page != nil {
		params = append(params, fmt.Sprintf("page[number]=%s", strconv.Itoa(page.Number)))
		params = append(params, fmt.Sprintf("page[size]=%s", strconv.Itoa(page.Size)))
	}
	if filter != nil {
		if filter.OrganisationID != nil {
			params = append(params, fmt.Sprintf("filter[organisation_id]=%s", *filter.OrganisationID))
		}
	}

	var URL bytes.Buffer
	URL.WriteString(fmt.Sprintf("%s%s", baseURL, path))
	for i, param := range params {
		if i == 0 {
			URL.WriteString(fmt.Sprintf("?%s", param))
		}
		URL.WriteString(fmt.Sprintf("&%s", param))
	}

	return URL.String()
}
