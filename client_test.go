package accounts

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	HTTPClient = http.Client{
		Timeout: time.Second * 2,
	}
	AccountsService = NewClient().HTTPClient(HTTPClient).URL(GetURL()).Build()
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

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").Build()).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		resp, _ := AccountsService.Create(AccountData)

		Convey("Then the required account fields should equal", func() {
			So(resp.AccountData.Attributes, ShouldResemble, AccountData.Attributes)
		})

		Convey("And the required data fields should equal", func() {
			So(resp.AccountData.ID, ShouldEqual, ID)
			So(resp.AccountData.OrganisationID, ShouldEqual, OrganisationID)
			So(resp.AccountData.Type, ShouldEqual, Type)
		})

	})

}

func TestCreateFullAccount(t *testing.T) {

	Convey("When I create an account with all fields", t, func() {
		ID := uuid.New().String()

		Account := NewAccount().
			Country("GB").
			BaseCurrency("GBP").
			BankID("400302").
			BankIDCode("GBDSC").
			AccountNumber("10000004").
			BIC("NWBKGB42").
			IBAN("GB28NWBK40030212764204").
			CustomerID("234").
			Title("Sie").
			FirstName("Mary-Jane Doe").
			BankAccountName("Smith").
			AlternativeBankAccountNames([]string{"Peters"}).
			AccountClassification("Personal").
			JointAccount(false).
			AccountMatchingOptOut(false).
			SecondaryIdentification("44516").
			Build()

		AccountData := NewAccountData().
			Attributes(Account).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		resp, _ := AccountsService.Create(AccountData)

		Convey("Then all the account fields should equal", func() {
			So(resp.AccountData.Attributes, ShouldResemble, AccountData.Attributes)
		})

	})

}

func TestCreateFailure(t *testing.T) {

	Convey("When I create an account on a non-existent server", t, func() {
		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").Build()).
			ID(uuid.New().String()).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		AccountsService := NewClient().HTTPClient(HTTPClient).URL("http://unknown:9999").Build()
		_, err := AccountsService.Create(AccountData)

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "An error has occured while creating account")
		})

	})

}

func TestCreateDuplicateAccount(t *testing.T) {

	Convey("Given I created an account", t, func() {
		ID := uuid.New().String()

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").Build()).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		AccountsService.Create(AccountData)

		Convey("When I create the account again with same ID", func() {
			_, err := AccountsService.Create(AccountData)

			Convey("Then an appropriate error is propagated to the caller", func() {
				So(err.Error(), ShouldEqual, "Account cannot be created as it violates a duplicate constraint")
			})

		})

	})

}

func TestCreateAccountWithInvalidCountry(t *testing.T) {

	Convey("When I create an account with invalid Country", t, func() {
		ID := uuid.New().String()

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GBR").Build()).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		_, err := AccountsService.Create(AccountData)

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid Country [GBR]")
		})

	})

}

func TestCreateAccountWithInvalidBaseCurrency(t *testing.T) {

	Convey("When I create an account with invalid BaseCurrency", t, func() {
		ID := uuid.New().String()

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").BaseCurrency("GBPP").Build()).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		_, err := AccountsService.Create(AccountData)

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid BaseCurrency [GBPP]")
		})

	})

}

func TestCreateAccountWithInvalidBankID(t *testing.T) {

	Convey("When I create an account with invalid BankID", t, func() {
		ID := uuid.New().String()

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").BankID("aStringLongerThanElevenCharacters").Build()).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		_, err := AccountsService.Create(AccountData)

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid BankID [aStringLongerThanElevenCharacters]")
		})

	})

}

func TestCreateAccountWithInvalidBIC(t *testing.T) {

	Convey("When I create an account with invalid BIC", t, func() {
		ID := uuid.New().String()

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").BIC("aStringLongerThanElevenCharacters").Build()).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		_, err := AccountsService.Create(AccountData)

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid BIC [aStringLongerThanElevenCharacters]")
		})

	})

}

func TestCreateAccountWithInvalidAccountClassification(t *testing.T) {

	Convey("When I create an account with invalid AccountClassification", t, func() {
		ID := uuid.New().String()

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").AccountClassification("unknown").Build()).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		_, err := AccountsService.Create(AccountData)

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid AccountClassification [unknown]")
		})

	})

}

func TestCreateAccountWithInvalidAlternativeBankAccountNames(t *testing.T) {

	Convey("When I create an account with invalid AlternativeBankAccountNames", t, func() {
		ID := uuid.New().String()

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").AlternativeBankAccountNames([]string{"Peters", "Michaels", "Johns", "Bens"}).Build()).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		_, err := AccountsService.Create(AccountData)

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "Invalid AlternativeBankAccountNames [Peters Michaels Johns Bens]")
		})

	})

}

func TestFetchBareMinimumAccount(t *testing.T) {

	Convey("Given I created an account with only required fields", t, func() {
		ID := uuid.New()

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").Build()).
			ID(ID.String()).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		AccountsService.Create(AccountData)

		Convey("When I fetch the account by ID", func() {

			resp, _ := AccountsService.Fetch(ID)

			Convey("Then the required account fields should equal", func() {
				So(resp.AccountData.Attributes, ShouldResemble, NewAccount().Country("GB").Build())
			})

		})

	})

}

func TestFetchFailure(t *testing.T) {

	Convey("When I fetch an account by ID on a non-existent server", t, func() {

		AccountsService := NewClient().HTTPClient(HTTPClient).URL("http://unknown:9999").Build()

		_, err := AccountsService.Fetch(uuid.New())

		Convey("Then an appropriate error is propagated to the caller", func() {
			So(err.Error(), ShouldEqual, "An error has occured while fetching account")
		})

	})

}

func TestFetchFullAccount(t *testing.T) {

	Convey("Given I created an account with all fields", t, func() {
		ID := uuid.New()

		Account := NewAccount().
			Country("GB").
			BaseCurrency("GBP").
			BankID("400302").
			BankIDCode("GBDSC").
			AccountNumber("10000004").
			BIC("NWBKGB42").
			IBAN("GB28NWBK40030212764204").
			CustomerID("234").
			Title("Sie").
			FirstName("Mary-Jane Doe").
			BankAccountName("Smith").
			AlternativeBankAccountNames([]string{"Peters"}).
			AccountClassification("Personal").
			JointAccount(false).
			AccountMatchingOptOut(false).
			SecondaryIdentification("44516").
			Build()

		AccountData := NewAccountData().
			Attributes(Account).
			ID(ID.String()).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		AccountsService.Create(AccountData)

		Convey("When I fetch the account by ID", func() {

			resp, _ := AccountsService.Fetch(ID)

			Convey("Then the required account fields should equal", func() {
				So(resp.AccountData.Attributes, ShouldResemble, AccountData.Attributes)
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

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").Build()).
			ID(ID.String()).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		AccountsService.Create(AccountData)

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

		AccountsService := NewClient().HTTPClient(HTTPClient).URL("http://unknown:9999").Build()

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

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").Build()).
			ID(ID.String()).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		AccountsService.Create(AccountData)

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

		AccountData := NewAccountData().
			Attributes(NewAccount().Country("GB").Build()).
			ID(ID).
			Type(Type).
			OrganisationID(OrganisationID).
			Build()

		AccountsService.Create(AccountData)

		Convey("When I list the accounts of the organisation", func() {

			resp, _ := AccountsService.List(nil, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain the account", func() {
				So(len(*resp.AccountData), ShouldEqual, 1)
			})

			Convey("And the account should match", func() {
				accountDataArr := *resp.AccountData
				So(accountDataArr[0].Attributes, ShouldResemble, NewAccount().Country("GB").Build())
			})

		})

	})

}

func TestListFailure(t *testing.T) {

	Convey("When I list the accounts of the organisation", t, func() {
		AccountsService := NewClient().HTTPClient(HTTPClient).URL("http://unknown:9999").Build()
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
			So(resp.AccountData, ShouldEqual, nil)
		})

	})

}

func TestListMultipleAccountsWithoutPaging(t *testing.T) {

	Convey("Given I created two accounts on an organisation", t, func() {

		OrganisationID = uuid.New().String()

		for i := 0; i < 2; i++ {
			ID := uuid.New().String()

			AccountData := NewAccountData().
				Attributes(NewAccount().Country("GB").Build()).
				ID(ID).
				Type(Type).
				OrganisationID(OrganisationID).
				Build()

			AccountsService.Create(AccountData)
		}

		Convey("When I list the accounts of the organisation without using pagination", func() {

			resp, _ := AccountsService.List(nil, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain the two accounts", func() {
				So(len(*resp.AccountData), ShouldEqual, 2)
			})

		})

	})

}

func TestListMultipleAccountsWithPaging(t *testing.T) {

	Convey("Given I created 10 accounts on an organisation", t, func() {

		OrganisationID = uuid.New().String()

		for i := 0; i < 10; i++ {
			ID := uuid.New().String()

			AccountData := NewAccountData().
				Attributes(NewAccount().Country("GB").BankID(strconv.Itoa(i)).Build()).
				ID(ID).
				Type(Type).
				OrganisationID(OrganisationID).
				Build()

			AccountsService.Create(AccountData)
		}

		Convey("When I list the accounts of the organisation with page size 10", func() {

			resp, _ := AccountsService.List(&Page{Number: 0, Size: 10}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain all accounts", func() {
				accountDataArr := *resp.AccountData
				So(len(accountDataArr), ShouldEqual, 10)
				So(*accountDataArr[0].Attributes.BankID, ShouldEqual, "0")
				So(*accountDataArr[9].Attributes.BankID, ShouldEqual, "9")
			})

		})

		Convey("When I list the accounts of the organisation with page size 5", func() {

			resp, _ := AccountsService.List(&Page{Number: 0, Size: 5}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain first 5 accounts", func() {
				accountDataArr := *resp.AccountData
				So(len(accountDataArr), ShouldEqual, 5)
				So(*accountDataArr[0].Attributes.BankID, ShouldEqual, "0")
				So(*accountDataArr[4].Attributes.BankID, ShouldEqual, "4")
			})

		})

		Convey("When I list the accounts of the organisation with page size 15", func() {

			resp, _ := AccountsService.List(&Page{Number: 0, Size: 15}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain all accounts", func() {
				accountDataArr := *resp.AccountData
				So(len(accountDataArr), ShouldEqual, 10)
				So(*accountDataArr[0].Attributes.BankID, ShouldEqual, "0")
				So(*accountDataArr[9].Attributes.BankID, ShouldEqual, "9")
			})

		})

		Convey("When I list the accounts of the organisation with page number 1 and page size 5", func() {

			resp, _ := AccountsService.List(&Page{Number: 1, Size: 5}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain last 5 accounts", func() {
				accountDataArr := *resp.AccountData
				So(len(accountDataArr), ShouldEqual, 5)
				So(*accountDataArr[0].Attributes.BankID, ShouldEqual, "5")
				So(*accountDataArr[4].Attributes.BankID, ShouldEqual, "9")
			})

		})

		Convey("When I list the accounts of the organisation with page number 1 and page size 3", func() {

			resp, _ := AccountsService.List(&Page{Number: 1, Size: 3}, &Filter{OrganisationID: &OrganisationID})

			Convey("Then the response should contain the last 4-6th accounts", func() {
				accountDataArr := *resp.AccountData
				So(len(accountDataArr), ShouldEqual, 3)
				So(*accountDataArr[0].Attributes.BankID, ShouldEqual, "3")
				So(*accountDataArr[2].Attributes.BankID, ShouldEqual, "5")
			})

		})

	})

}
