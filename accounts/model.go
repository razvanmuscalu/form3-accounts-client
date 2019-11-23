package accounts

// DataRequest is the request payload used when creating an account
type DataRequest struct {
	Data Data `json:"data"`
}

// Single is the response payload when fetching an individual account
type Single struct {
	Data  Data  `json:"data"`
	Links Links `json:"links"`
}

// List is the response payload when requesting a list of accounts
type List struct {
	Data  []Data `json:"data"`
	Links Links  `json:"links"`
}

// Data is a wrapper object around the Attributes object
type Data struct {
	ID             string  `json:"id"`
	OrganisationID string  `json:"organisation_id"`
	Type           string  `json:"type"`
	CreatedOn      *string `json:"created_on,omitempty"`
	ModifiedOn     *string `json:"modified_on,omitempty"`
	Version        *int    `json:"version,omitempty"`
	Attributes     Account `json:"attributes"`
}

// Links holds pointers to the first, last and self objects returned in the response
type Links struct {
	First *string `json:"first,omitempty"`
	Last  *string `json:"last,omitempty"`
	Self  string  `json:"self"`
}

// Account is the main domain model representing an account
type Account struct {
	Country                     string     `json:"country"`
	BaseCurrency                *string    `json:"base_currency,omitempty"`
	BankID                      *string    `json:"bank_id,omitempty"`
	BankIDCode                  *string    `json:"bank_id_code,omitempty"`
	AccountNumber               *string    `json:"account_number,omitempty"`
	BIC                         *string    `json:"bic,omitempty"`
	IBAN                        *string    `json:"iban,omitempty"`
	CustomerID                  *string    `json:"customer_id,omitempty"`
	Title                       *string    `json:"title,omitempty"`
	FirstName                   *string    `json:"first_name,omitempty"`
	BankAccountName             *string    `json:"bank_account_name,omitempty"`
	AlternativeBankAccountNames *[3]string `json:"alternative_bank_account_names,omitempty"`
	AccountClassification       *string    `json:"account_classification,omitempty"`
	JointAccount                *bool      `json:"joint_account,omitempty"`
	AccountMatchingOptOut       *bool      `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification     *string    `json:"secondary_identification,omitempty"`
}
