package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

// Client is the client interface to access the Accounts API
type Client interface {
	Create(request AccountData) (Single, error)
	Fetch(id uuid.UUID) (Single, error)
	List(page *Page, filter *Filter) (List, error)
	Delete(id uuid.UUID, version int) (bool, error)
}

// Service implements Client
type service struct {
	URL        string
	HTTPClient http.Client
}

// NewClientWithURL return a new API client
func NewClientWithURL(url string, httpClient http.Client) Client {
	return service{
		URL:        url,
		HTTPClient: httpClient,
	}
}

// NewClient return a new API client
func NewClient(httpClient http.Client) Client {
	return service{HTTPClient: httpClient}
}

func (s *service) GetURL() string {
	if s.URL != "" {
		return s.URL
	}
	return "http://localhost:8080"
}

const path = "/v1/organisation/accounts"

// Create an account
//
// The request is pre-validated to avoid unnecessary Bad Request
func (s service) Create(request AccountData) (Single, error) {

	if err := validateAccount(request.Attributes); err != nil {
		return Single{}, err
	}

	req := new(bytes.Buffer)
	json.NewEncoder(req).Encode(AccountDataRequest{AccountData: request})
	resp, err := s.HTTPClient.Post(fmt.Sprintf("%s%s", s.GetURL(), path), "application/json", req)
	if err != nil {
		return Single{}, fmt.Errorf("An error has occured while creating account")
	}
	defer resp.Body.Close()

	if err := decodeErrorResponse(resp); err != nil {
		return Single{}, err
	}

	var result Single
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil

}

// Fetch an account
func (s service) Fetch(id uuid.UUID) (Single, error) {

	resp, err := s.HTTPClient.Get(fmt.Sprintf("%s%s/%s", s.GetURL(), path, id))
	if err != nil {
		return Single{}, fmt.Errorf("An error has occured while fetching account")
	}
	defer resp.Body.Close()

	if err := decodeErrorResponse(resp); err != nil {
		return Single{}, err
	}

	var result Single
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil

}

// List accounts
func (s service) List(page *Page, filter *Filter) (List, error) {

	resp, err := s.HTTPClient.Get(buildListURL(s.GetURL(), page, filter))
	if err != nil {
		return List{}, fmt.Errorf("An error has occured while listing accounts")
	}
	defer resp.Body.Close()

	if err := decodeErrorResponse(resp); err != nil {
		return List{}, err
	}

	var result List
	json.NewDecoder(resp.Body).Decode(&result)

	return result, nil
}

// Delete an account
func (s service) Delete(id uuid.UUID, version int) (bool, error) {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/%s?version=%s", s.GetURL(), path, id, strconv.Itoa(version)), new(bytes.Buffer))
	if err != nil {
		return false, fmt.Errorf("An error has occured while constructing delete request")
	}

	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("An error has occured while deleting account")
	}
	defer resp.Body.Close()

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
