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

// Page holds the requested page number and size on the List function
type Page struct {
	Number int
	Size   int
}

// Filter holds the requested filter on the List function
type Filter struct {
	OrganisationID *string
}

// Client is the client interface to access the Accounts API
type Client interface {
	Create(request accountData) (Single, error)
	Fetch(id uuid.UUID) (Single, error)
	List(page *Page, filter *Filter) (List, error)
	Delete(id uuid.UUID, version int) (bool, error)
}

type client struct {
	url        string
	httpClient http.Client
}

// ClientBuilder is used to create a Client
type ClientBuilder interface {
	URL(string) ClientBuilder
	HTTPClient(http.Client) ClientBuilder
	Build() Client
}

type clientBuilder struct {
	url        string
	httpClient http.Client
}

func (cb *clientBuilder) URL(value string) ClientBuilder {
	cb.url = value
	return cb
}

func (cb *clientBuilder) HTTPClient(value http.Client) ClientBuilder {
	cb.httpClient = value
	return cb
}

func (cb *clientBuilder) Build() Client {
	return &client{
		url:        cb.url,
		httpClient: cb.httpClient,
	}
}

// NewClient is used to create a ClientBuilder
func NewClient() ClientBuilder {
	return &clientBuilder{}
}

const path = "/v1/organisation/accounts"

// Create an account
//
// The request is pre-validated to avoid unnecessary Bad Request
func (c client) Create(request accountData) (Single, error) {

	if err := validateAccount(request.Attributes); err != nil {
		return Single{}, err
	}

	req := new(bytes.Buffer)
	if err := json.NewEncoder(req).Encode(NewAccountDataRequest().AccountData(request).Build()); err != nil {
		return Single{}, fmt.Errorf("An error has occured while encoding request")
	}
	resp, err := c.httpClient.Post(fmt.Sprintf("%s%s", c.url, path), "application/json", req)
	if err != nil {
		return Single{}, fmt.Errorf("An error has occured while creating account")
	}
	defer resp.Body.Close()

	if err := decodeErrorResponse(resp); err != nil {
		return Single{}, err
	}

	var result Single
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Single{}, fmt.Errorf("An error has occured while decoding response")
	}

	return result, nil

}

// Fetch an account
func (c client) Fetch(id uuid.UUID) (Single, error) {

	resp, err := c.httpClient.Get(fmt.Sprintf("%s%s/%s", c.url, path, id))
	if err != nil {
		return Single{}, fmt.Errorf("An error has occured while fetching account")
	}
	defer resp.Body.Close()

	if err := decodeErrorResponse(resp); err != nil {
		return Single{}, err
	}

	var result Single
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return Single{}, fmt.Errorf("An error has occured while decoding response")
	}

	return result, nil

}

// List accounts
func (c client) List(page *Page, filter *Filter) (List, error) {

	resp, err := c.httpClient.Get(buildListURL(c.url, page, filter))
	if err != nil {
		return List{}, fmt.Errorf("An error has occured while listing accounts")
	}
	defer resp.Body.Close()

	if err := decodeErrorResponse(resp); err != nil {
		return List{}, err
	}

	var result List
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return List{}, fmt.Errorf("An error has occured while decoding response")
	}

	return result, nil
}

// Delete an account
func (c client) Delete(id uuid.UUID, version int) (bool, error) {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s/%s?version=%s", c.url, path, id, strconv.Itoa(version)), new(bytes.Buffer))
	if err != nil {
		return false, fmt.Errorf("An error has occured while constructing delete request")
	}

	resp, err := c.httpClient.Do(req)
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
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return fmt.Errorf("An error has occured while decoding error response")
		}

		return errors.New(result.ErrorMessage)
	}

	return nil
}

func validateAccount(account account) error {
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
