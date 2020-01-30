package accounts

type accountData struct {
	ID             string  `json:"id"`
	OrganisationID string  `json:"organisation_id"`
	Type           string  `json:"type"`
	CreatedOn      *string `json:"created_on,omitempty"`
	ModifiedOn     *string `json:"modified_on,omitempty"`
	Version        *int    `json:"version,omitempty"`
	Attributes     account `json:"attributes"`
}

// AccountDataBuilder returns a builder for accountData struct
type AccountDataBuilder interface {
	ID(string) AccountDataBuilder
	OrganisationID(string) AccountDataBuilder
	Type(string) AccountDataBuilder
	CreatedOn(string) AccountDataBuilder
	ModifiedOn(string) AccountDataBuilder
	Version(int) AccountDataBuilder
	Attributes(account) AccountDataBuilder
	Build() accountData
}

type accountDataBuilder struct {
	id              string
	organisationID  string
	accountDataType string
	createdOn       *string
	modifiedOn      *string
	version         *int
	attributes      account
}

func (ab *accountDataBuilder) ID(value string) AccountDataBuilder {
	ab.id = value
	return ab
}

func (ab *accountDataBuilder) OrganisationID(value string) AccountDataBuilder {
	ab.organisationID = value
	return ab
}

func (ab *accountDataBuilder) Type(value string) AccountDataBuilder {
	ab.accountDataType = value
	return ab
}

func (ab *accountDataBuilder) CreatedOn(value string) AccountDataBuilder {
	ab.createdOn = &value
	return ab
}

func (ab *accountDataBuilder) ModifiedOn(value string) AccountDataBuilder {
	ab.modifiedOn = &value
	return ab
}

func (ab *accountDataBuilder) Version(value int) AccountDataBuilder {
	ab.version = &value
	return ab
}

func (ab *accountDataBuilder) Attributes(value account) AccountDataBuilder {
	ab.attributes = value
	return ab
}

func (ab *accountDataBuilder) Build() accountData {
	return accountData{
		ID:             ab.id,
		OrganisationID: ab.organisationID,
		Type:           ab.accountDataType,
		CreatedOn:      ab.createdOn,
		ModifiedOn:     ab.modifiedOn,
		Version:        ab.version,
		Attributes:     ab.attributes,
	}
}

// NewAccountData is used to create an AccountDataBuilder
func NewAccountData() AccountDataBuilder {
	return &accountDataBuilder{}
}
