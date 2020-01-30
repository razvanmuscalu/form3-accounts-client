package accounts

type account struct {
	Country                     string    `json:"country"`
	BaseCurrency                *string   `json:"base_currency,omitempty"`
	BankID                      *string   `json:"bank_id,omitempty"`
	BankIDCode                  *string   `json:"bank_id_code,omitempty"`
	AccountNumber               *string   `json:"account_number,omitempty"`
	BIC                         *string   `json:"bic,omitempty"`
	IBAN                        *string   `json:"iban,omitempty"`
	CustomerID                  *string   `json:"customer_id,omitempty"`
	Title                       *string   `json:"title,omitempty"`
	FirstName                   *string   `json:"first_name,omitempty"`
	BankAccountName             *string   `json:"bank_account_name,omitempty"`
	AlternativeBankAccountNames *[]string `json:"alternative_bank_account_names,omitempty"`
	AccountClassification       *string   `json:"account_classification,omitempty"`
	JointAccount                *bool     `json:"joint_account,omitempty"`
	AccountMatchingOptOut       *bool     `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification     *string   `json:"secondary_identification,omitempty"`
}

// AccountBuilder returns a builder for account struct
type AccountBuilder interface {
	Country(string) AccountBuilder
	BaseCurrency(string) AccountBuilder
	BankID(string) AccountBuilder
	BankIDCode(string) AccountBuilder
	AccountNumber(string) AccountBuilder
	BIC(string) AccountBuilder
	IBAN(string) AccountBuilder
	CustomerID(string) AccountBuilder
	Title(string) AccountBuilder
	FirstName(string) AccountBuilder
	BankAccountName(string) AccountBuilder
	AlternativeBankAccountNames([]string) AccountBuilder
	AccountClassification(string) AccountBuilder
	JointAccount(bool) AccountBuilder
	AccountMatchingOptOut(bool) AccountBuilder
	SecondaryIdentification(string) AccountBuilder
	Build() account
}

type accountBuilder struct {
	country                     string
	baseCurrency                *string
	bankID                      *string
	bankIDCode                  *string
	accountNumber               *string
	bic                         *string
	iban                        *string
	customerID                  *string
	title                       *string
	firstName                   *string
	bankAccountName             *string
	alternativeBankAccountNames *[]string
	accountClassification       *string
	jointAccount                *bool
	accountMatchingOptOut       *bool
	secondaryIdentification     *string
}

func (ab *accountBuilder) Country(value string) AccountBuilder {
	ab.country = value
	return ab
}

func (ab *accountBuilder) BaseCurrency(value string) AccountBuilder {
	ab.baseCurrency = &value
	return ab
}

func (ab *accountBuilder) BankID(value string) AccountBuilder {
	ab.bankID = &value
	return ab
}

func (ab *accountBuilder) BankIDCode(value string) AccountBuilder {
	ab.bankIDCode = &value
	return ab
}

func (ab *accountBuilder) AccountNumber(value string) AccountBuilder {
	ab.accountNumber = &value
	return ab
}

func (ab *accountBuilder) BIC(value string) AccountBuilder {
	ab.bic = &value
	return ab
}

func (ab *accountBuilder) IBAN(value string) AccountBuilder {
	ab.iban = &value
	return ab
}

func (ab *accountBuilder) CustomerID(value string) AccountBuilder {
	ab.customerID = &value
	return ab
}

func (ab *accountBuilder) Title(value string) AccountBuilder {
	ab.title = &value
	return ab
}

func (ab *accountBuilder) FirstName(value string) AccountBuilder {
	ab.firstName = &value
	return ab
}

func (ab *accountBuilder) BankAccountName(value string) AccountBuilder {
	ab.bankAccountName = &value
	return ab
}

func (ab *accountBuilder) AlternativeBankAccountNames(value []string) AccountBuilder {
	ab.alternativeBankAccountNames = &value
	return ab
}

func (ab *accountBuilder) AccountClassification(value string) AccountBuilder {
	ab.accountClassification = &value
	return ab
}

func (ab *accountBuilder) JointAccount(value bool) AccountBuilder {
	ab.jointAccount = &value
	return ab
}

func (ab *accountBuilder) AccountMatchingOptOut(value bool) AccountBuilder {
	ab.accountMatchingOptOut = &value
	return ab
}

func (ab *accountBuilder) SecondaryIdentification(value string) AccountBuilder {
	ab.secondaryIdentification = &value
	return ab
}

func (ab *accountBuilder) Build() account {
	return account{
		Country:                     ab.country,
		BaseCurrency:                ab.baseCurrency,
		BankID:                      ab.bankID,
		BankIDCode:                  ab.bankIDCode,
		AccountNumber:               ab.accountNumber,
		BIC:                         ab.bic,
		IBAN:                        ab.iban,
		CustomerID:                  ab.customerID,
		Title:                       ab.title,
		FirstName:                   ab.firstName,
		BankAccountName:             ab.bankAccountName,
		AlternativeBankAccountNames: ab.alternativeBankAccountNames,
		AccountClassification:       ab.accountClassification,
		JointAccount:                ab.jointAccount,
		AccountMatchingOptOut:       ab.accountMatchingOptOut,
		SecondaryIdentification:     ab.secondaryIdentification,
	}
}

// NewAccount is used to create an AccountBuilder
func NewAccount() AccountBuilder {
	return &accountBuilder{}
}
