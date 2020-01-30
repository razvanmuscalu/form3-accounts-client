package accounts

// Single is the response payload when fetching an individual account
type Single struct {
	AccountData accountData `json:"data"`
	Links       Links       `json:"links"`
}

// List is the response payload when requesting a list of accounts
type List struct {
	AccountData *[]accountData `json:"data,omitempty"`
	Links       Links          `json:"links"`
}

// Links holds pointers to the first, last and self objects returned in the response
type Links struct {
	First *string `json:"first,omitempty"`
	Last  *string `json:"last,omitempty"`
	Next  *string `json:"next,omitempty"`
	Prev  *string `json:"prev,omitempty"`
	Self  string  `json:"self"`
}

// ErrorResponse represents a failed Form3 Accounts API response
type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}
