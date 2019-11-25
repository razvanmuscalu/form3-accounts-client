package accounts

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/google/uuid"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	AccountsService = Service{URL: GetURL()}
	OrganisationID  = uuid.New().String()
	Type            = "accounts"
)

func GetURL() string {
	value := os.Getenv("ACCOUNTS_API_URL")
	if len(value) == 0 {
		return "http://localhost:8080"
	}
	return value
}

func TestCreateBareMinimumAccount(t *testing.T) {

	Convey("When I create an account with only required fields", t, func() {
		ID := uuid.New().String()
		DataRequest := NewDataRequest(ID, OrganisationID)
		resp, _ := AccountsService.Create(DataRequest)

		Convey("Then the required account fields should equal", func() {
			So(resp.Data.Attributes, ShouldResemble, DataRequest.Data.Attributes)
		})

		Convey("And the required data fields should equal", func() {
			So(resp.Data.ID, ShouldEqual, ID)
			So(resp.Data.OrganisationID, ShouldEqual, OrganisationID)
			So(resp.Data.Type, ShouldEqual, Type)
		})

	})

}

func TestCreateFullAccount(t *testing.T) {

	Convey("When I create an account with all fields", t, func() {
		ID := uuid.New().String()

		DataRequest := NewFullDataRequest(ID)

		resp, _ := AccountsService.Create(DataRequest)

		Convey("Then all the account fields should equal", func() {
			So(resp.Data.Attributes, ShouldResemble, DataRequest.Data.Attributes)
		})

	})

}

func TestCreateFailure(t *testing.T) {

	Convey("When I create an account on a non-existent server", t, func() {
		DataRequest := NewDataRequest(uuid.New().String(), OrganisationID)
		AccountsService := Service{URL: "http://unknown:9999"}
		_, err := AccountsService.Create(DataRequest)

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "An error has occured while creating account")
		})

	})

}

func TestCreateDuplicateAccount(t *testing.T) {

	Convey("Given I created an account", t, func() {
		ID := uuid.New().String()

		DataRequest := NewDataRequest(ID, OrganisationID)

		AccountsService.Create(DataRequest)

		Convey("When I create the account again with same ID", func() {
			_, err := AccountsService.Create(DataRequest)

			Convey("Then an appropriate error is propagated to the caller", func() {
				So(err.Error(), ShouldEqual, "Account cannot be created as it violates a duplicate constraint")
			})

		})

	})

}

func TestCreateAccountWithInvalidCountry(t *testing.T) {

	Convey("When I create an account with invalid Country", t, func() {
		ID := uuid.New().String()

		DataRequest := DataRequest{
			Data: Data{
				Attributes:     Account{Country: "GBR"},
				ID:             ID,
				Type:           Type,
				OrganisationID: OrganisationID,
			}}

		_, err := AccountsService.Create(DataRequest)

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid Country [GBR]")
		})

	})

}

func TestCreateAccountWithInvalidBaseCurrency(t *testing.T) {

	Convey("When I create an account with invalid BaseCurrency", t, func() {
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

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid BaseCurrency [GBPP]")
		})

	})

}

func TestCreateAccountWithInvalidBankID(t *testing.T) {

	Convey("When I create an account with invalid BankID", t, func() {
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

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid BankID [aStringLongerThanElevenCharacters]")
		})

	})

}

func TestCreateAccountWithInvalidBIC(t *testing.T) {

	Convey("When I create an account with invalid BIC", t, func() {
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

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid BIC [aStringLongerThanElevenCharacters]")
		})

	})

}

func TestCreateAccountWithInvalidAccountClassification(t *testing.T) {

	Convey("When I create an account with invalid AccountClassification", t, func() {
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

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid AccountClassification [unknown]")
		})

	})

}

func TestCreateAccountWithInvalidAlternativeBankAccountNames(t *testing.T) {

	Convey("When I create an account with invalid AlternativeBankAccountNames", t, func() {
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

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid AlternativeBankAccountNames [Peters Michaels Johns Bens]")
		})

	})

}

func TestFetchBareMinimumAccount(t *testing.T) {

	Convey("Given I created an account with only required fields", t, func() {
		ID := uuid.New()

		AccountsService.Create(NewDataRequest(ID.String(), OrganisationID))

		Convey("When I fetch the account by ID", func() {

			resp, _ := AccountsService.Fetch(ID)

			Convey("Then the required account fields should equal", func() {
				So(resp.Data.Attributes, ShouldResemble, Account{Country: "GB"})
			})

		})

	})

}

func TestFetchFailure(t *testing.T) {

	Convey("When I fetch an account by ID on a non-existent server", t, func() {

		AccountsService := Service{URL: "http://unknown:9999"}

		_, err := AccountsService.Fetch(uuid.New())

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "An error has occured while fetching account")
		})

	})

}

func TestFetchFullAccount(t *testing.T) {

	Convey("Given I created an account with all fields", t, func() {
		ID := uuid.New()

		DataRequest := NewFullDataRequest(ID.String())

		AccountsService.Create(DataRequest)

		Convey("When I fetch the account by ID", func() {

			resp, _ := AccountsService.Fetch(ID)

			Convey("Then the required account fields should equal", func() {
				So(resp.Data.Attributes, ShouldResemble, DataRequest.Data.Attributes)
			})

		})

	})

}

func TestFetchNonExistentAccount(t *testing.T) {

	Convey("When I fetch an account by a random ID", t, func() {
		ID := uuid.New()
		_, err := AccountsService.Fetch(ID)

		Convey("The an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, fmt.Sprintf("record %s does not exist", ID))
		})
	})

}

func TestDeleteAccount(t *testing.T) {

	Convey("Given I created an account", t, func() {
		ID := uuid.New()

		AccountsService.Create(NewDataRequest(ID.String(), OrganisationID))

		Convey("When I delete the account by ID", func() {

			resp, _ := AccountsService.Delete(ID, 0)

			Convey("Then the response is true", func() {
				So(resp, ShouldEqual, true)
			})

		})

	})

}

func TestDeleteFailure(t *testing.T) {

	Convey("When I delete an account by ID on a non-existent server", t, func() {

		AccountsService := Service{URL: "http://unknown:9999"}

		_, err := AccountsService.Delete(uuid.New(), 0)

		Convey("The an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "An error has occured while deleting account")
		})

	})

}

func TestDeleteNonExistentAccount(t *testing.T) {

	Convey("When I delete an account by a random ID", t, func() {
		ID := uuid.New()
		resp, _ := AccountsService.Delete(ID, 0)

		Convey("The the response is true", func() {
			So(resp, ShouldEqual, true)
		})
	})

}

func TestDeleteNonExistentAccountVersion(t *testing.T) {

	Convey("Given I created an account", t, func() {
		ID := uuid.New()

		AccountsService.Create(NewDataRequest(ID.String(), OrganisationID))

		Convey("When I delete the account by ID and non-existent version", func() {

			resp, _ := AccountsService.Delete(ID, 1)

			Convey("Then the response is false", func() {
				So(resp, ShouldEqual, false)
			})

		})

	})

}

func TestListOneAccount(t *testing.T) {

	Convey("Given I created an account of an organisation", t, func() {

		OrganisationID := uuid.New().String()
		ID := uuid.New().String()

		DataRequest := NewDataRequest(ID, OrganisationID)

		AccountsService.Create(DataRequest)

		Convey("When I list the accounts of the organisation", func() {

			resp, _ := AccountsService.List(nil, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain the account", func() {
				So(len(*resp.Data), ShouldEqual, 1)
			})

			Convey("And the account should match", func() {
				dataArr := *resp.Data
				So(dataArr[0].Attributes, ShouldResemble, Account{Country: "GB"})
			})

		})

	})

}

func TestListFailure(t *testing.T) {

	Convey("When I list the accounts of the organisation", t, func() {
		AccountsService := Service{URL: "http://unknown:9999"}
		_, err := AccountsService.List(nil, nil)

		Convey("The an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "An error has occured while listing accounts")
		})

	})

}

func TestListZeroAccounts(t *testing.T) {

	Convey("When I list the accounts of an organisation", t, func() {

		OrganisationID := uuid.New().String()

		resp, _ := AccountsService.List(nil, &Filter{OrganisationID: &OrganisationID})

		Convey("Then the response should contain no accounts", func() {
			So(resp.Data, ShouldEqual, nil)
		})

	})

}

func TestListMultipleAccountsWithoutPaging(t *testing.T) {

	Convey("Given I created two accounts on an organisation", t, func() {

		OrganisationID = uuid.New().String()

		for i := 0; i < 2; i++ {
			ID := uuid.New().String()
			DataRequest := NewDataRequest(ID, OrganisationID)

			AccountsService.Create(DataRequest)
		}

		Convey("When I list the accounts of the organisation without using pagination", func() {

			resp, _ := AccountsService.List(nil, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain the two accounts", func() {
				So(len(*resp.Data), ShouldEqual, 2)
			})

		})

	})

}

func TestListMultipleAccountsWithPaging(t *testing.T) {

	Convey("Given I created 10 accounts on an organisation", t, func() {

		OrganisationID = uuid.New().String()

		for i := 0; i < 10; i++ {
			ID := uuid.New().String()
			BankID := strconv.Itoa(i)

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

			AccountsService.Create(DataRequest)
		}

		Convey("When I list the accounts of the organisation with page size 10", func() {

			resp, _ := AccountsService.List(&Page{Number: 0, Size: 10}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain all accounts", func() {
				dataArr := *resp.Data
				So(len(dataArr), ShouldEqual, 10)
				So(*dataArr[0].Attributes.BankID, ShouldEqual, "0")
				So(*dataArr[9].Attributes.BankID, ShouldEqual, "9")
			})

		})

		Convey("When I list the accounts of the organisation with page size 5", func() {

			resp, _ := AccountsService.List(&Page{Number: 0, Size: 5}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain first 5 accounts", func() {
				dataArr := *resp.Data
				So(len(dataArr), ShouldEqual, 5)
				So(*dataArr[0].Attributes.BankID, ShouldEqual, "0")
				So(*dataArr[4].Attributes.BankID, ShouldEqual, "4")
			})

		})

		Convey("When I list the accounts of the organisation with page size 15", func() {

			resp, _ := AccountsService.List(&Page{Number: 0, Size: 15}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain all accounts", func() {
				dataArr := *resp.Data
				So(len(dataArr), ShouldEqual, 10)
				So(*dataArr[0].Attributes.BankID, ShouldEqual, "0")
				So(*dataArr[9].Attributes.BankID, ShouldEqual, "9")
			})

		})

		Convey("When I list the accounts of the organisation with page number 1 and page size 5", func() {

			resp, _ := AccountsService.List(&Page{Number: 1, Size: 5}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain last 5 accounts", func() {
				dataArr := *resp.Data
				So(len(dataArr), ShouldEqual, 5)
				So(*dataArr[0].Attributes.BankID, ShouldEqual, "5")
				So(*dataArr[4].Attributes.BankID, ShouldEqual, "9")
			})

		})

		Convey("When I list the accounts of the organisation with page number 1 and page size 3", func() {

			resp, _ := AccountsService.List(&Page{Number: 1, Size: 3}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain the last 4-6th accounts", func() {
				dataArr := *resp.Data
				So(len(dataArr), ShouldEqual, 3)
				So(*dataArr[0].Attributes.BankID, ShouldEqual, "3")
				So(*dataArr[2].Attributes.BankID, ShouldEqual, "5")
			})

		})

	})

}

func NewDataRequest(ID string, OrganisationID string) DataRequest {

	return DataRequest{
		Data: Data{
			Attributes:     Account{Country: "GB"},
			ID:             ID,
			Type:           Type,
			OrganisationID: OrganisationID,
		}}

}

func NewFullDataRequest(ID string) DataRequest {

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

	return DataRequest{
		Data: Data{
			Attributes:     Account,
			ID:             ID,
			Type:           Type,
			OrganisationID: OrganisationID,
		}}

}
