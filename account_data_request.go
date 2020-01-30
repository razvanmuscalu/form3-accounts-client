package accounts

type accountDataRequest struct {
	AccountData accountData `json:"data"`
}

// AccountDataRequestBuilder returns a builder for accountDataRequest struct
type AccountDataRequestBuilder interface {
	AccountData(accountData) AccountDataRequestBuilder
	Build() accountDataRequest
}

type accountDataRequestBuilder struct {
	accountData accountData
}

func (ab *accountDataRequestBuilder) AccountData(value accountData) AccountDataRequestBuilder {
	ab.accountData = value
	return ab
}

func (ab *accountDataRequestBuilder) Build() accountDataRequest {
	return accountDataRequest{AccountData: ab.accountData}
}

// NewAccountDataRequest is used to create an AccountDataRequestBuilder
func NewAccountDataRequest() AccountDataRequestBuilder {
	return &accountDataRequestBuilder{}
}
