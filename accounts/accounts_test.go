package accounts

import (
	"testing"

	"github.com/google/uuid"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	AccountsService = Service{URL: "http://localhost:8080"}
	OrganisationID  = uuid.New().String()
	Type            = "accounts"
)

func TestCreateBareMinimumAccount(t *testing.T) {
	Convey("When I Create an account with only required fields", t, func() {
		ID := uuid.New().String()

		DataRequest := DataRequest{
			Data: Data{
				Attributes:     Account{Country: "GB"},
				ID:             ID,
				Type:           Type,
				OrganisationID: OrganisationID,
			}}

		resp, _ := AccountsService.Create(DataRequest)

		Convey("The required account fields shout equal", func() {
			So(resp.Data.Attributes, ShouldResemble, Account{Country: "GB"})
		})

		Convey("And the required data fields shout equal", func() {
			So(resp.Data.ID, ShouldEqual, ID)
			So(resp.Data.OrganisationID, ShouldEqual, OrganisationID)
			So(resp.Data.Type, ShouldEqual, Type)
		})
	})
}

func TestCreateFullAccount(t *testing.T) {
	Convey("When I Create an account with all fields", t, func() {
		ID := uuid.New().String()

		BaseCurrency := "GBP"
		BankID := "400302"
		BankIDCode := "GBDSC"
		AccountNumber := "10000004"
		BIC := "NWBKGB42"
		IBAN := "GB28NWBK40030212764204"
		CustomerID := "234"
		Title := "Sie"
		FirstName := "Mary-Jane Doe"
		BankAccountName := "Smith"
		AlternativeBankAccountNames := []string{"Peters"}
		AccountClassification := "Personal"
		JointAccount := false
		AccountMatchingOptOut := false
		SecondaryIdentification := "44516"

		Account := Account{
			Country:                     "GB",
			BaseCurrency:                &BaseCurrency,
			BankID:                      &BankID,
			BankIDCode:                  &BankIDCode,
			AccountNumber:               &AccountNumber,
			BIC:                         &BIC,
			IBAN:                        &IBAN,
			CustomerID:                  &CustomerID,
			Title:                       &Title,
			FirstName:                   &FirstName,
			BankAccountName:             &BankAccountName,
			AlternativeBankAccountNames: &AlternativeBankAccountNames,
			AccountClassification:       &AccountClassification,
			JointAccount:                &JointAccount,
			AccountMatchingOptOut:       &AccountMatchingOptOut,
			SecondaryIdentification:     &SecondaryIdentification,
		}
		DataRequest := DataRequest{
			Data: Data{
				Attributes:     Account,
				ID:             ID,
				Type:           Type,
				OrganisationID: OrganisationID,
			}}

		resp, _ := AccountsService.Create(DataRequest)

		Convey("All the account fields shout equal", func() {
			So(resp.Data.Attributes, ShouldResemble, Account)
		})
	})
}

func TestCreateAccountWithInvalidCountry(t *testing.T) {
	Convey("When I Create an account with invalid Country", t, func() {
		ID := uuid.New().String()

		DataRequest := DataRequest{
			Data: Data{
				Attributes:     Account{Country: "GBR"},
				ID:             ID,
				Type:           Type,
				OrganisationID: OrganisationID,
			}}

		_, err := AccountsService.Create(DataRequest)

		Convey("An appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldContainSubstring, "Invalid Country [GBR]")
		})
	})
}

func TestCreateAccountWithInvalidBaseCurrency(t *testing.T) {
	Convey("When I Create an account with invalid BaseCurrency", t, func() {
		ID := uuid.New().String()

		BaseCurrency := "GBPP"

		DataRequest := DataRequest{
			Data: Data{
				Attributes: Account{
					Country:      "GB",
					BaseCurrency: &BaseCurrency,
				},
				ID:             ID,
				Type:           Type,
				OrganisationID: OrganisationID,
			}}

		_, err := AccountsService.Create(DataRequest)

		Convey("An appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldContainSubstring, "Invalid BaseCurrency [GBPP]")
		})
	})
}

func TestCreateAccountWithInvalidBankID(t *testing.T) {
	Convey("When I Create an account with invalid BankID", t, func() {
		ID := uuid.New().String()

		BankID := "aStringLongerThanElevenCharacters"

		DataRequest := DataRequest{
			Data: Data{
				Attributes: Account{
					Country: "GB",
					BankID:  &BankID,
				},
				ID:             ID,
				Type:           Type,
				OrganisationID: OrganisationID,
			}}

		_, err := AccountsService.Create(DataRequest)

		Convey("An appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldContainSubstring, "Invalid BankID [aStringLongerThanElevenCharacters]")
		})
	})
}

func TestCreateAccountWithInvalidBIC(t *testing.T) {
	Convey("When I Create an account with invalid BIC", t, func() {
		ID := uuid.New().String()

		BIC := "aStringLongerThanElevenCharacters"

		DataRequest := DataRequest{
			Data: Data{
				Attributes: Account{
					Country: "GB",
					BIC:     &BIC,
				},
				ID:             ID,
				Type:           Type,
				OrganisationID: OrganisationID,
			}}

		_, err := AccountsService.Create(DataRequest)

		Convey("An appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldContainSubstring, "Invalid BIC [aStringLongerThanElevenCharacters]")
		})
	})
}

func TestCreateAccountWithInvalidAccountClassification(t *testing.T) {
	Convey("When I Create an account with invalid AccountClassification", t, func() {
		ID := uuid.New().String()

		AccountClassification := "unknown"

		DataRequest := DataRequest{
			Data: Data{
				Attributes: Account{
					Country:               "GB",
					AccountClassification: &AccountClassification,
				},
				ID:             ID,
				Type:           Type,
				OrganisationID: OrganisationID,
			}}

		_, err := AccountsService.Create(DataRequest)

		Convey("An appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldContainSubstring, "Invalid AccountClassification [unknown]")
		})
	})
}

func TestCreateAccountWithInvalidAlternativeBankAccountNames(t *testing.T) {
	Convey("When I Create an account with invalid AlternativeBankAccountNames", t, func() {
		ID := uuid.New().String()

		AlternativeBankAccountNames := []string{"Peters", "Michaels", "Johns", "Bens"}

		DataRequest := DataRequest{
			Data: Data{
				Attributes: Account{
					Country:                     "GB",
					AlternativeBankAccountNames: &AlternativeBankAccountNames,
				},
				ID:             ID,
				Type:           Type,
				OrganisationID: OrganisationID,
			}}

		_, err := AccountsService.Create(DataRequest)

		Convey("An appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldContainSubstring, "Invalid AlternativeBankAccountNames [Peters Michaels Johns Bens]")
		})
	})
}
