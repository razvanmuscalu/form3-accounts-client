package accounts

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFetch(t *testing.T) {
	service := Service{URL: ""}

	Convey("When I Fetch accounts", t, func() {
		resp, _ := service.Fetch("c24e60c2-9dcb-466a-8fc4-fdf19c630b29")

		Convey("The AccountNumber should match", func() {
			So(*resp.Data.Attributes.AccountNumber, ShouldEqual, "123")
		})

		Convey("The FirstName should match", func() {
			So(*resp.Data.Attributes.FirstName, ShouldEqual, "razvan")
		})
	})
}

func TestCreate(t *testing.T) {
	service := Service{URL: "http://localhost:8080"}

	account := Account{Country: "GB"}
	data := Data{
		Attributes:     account,
		ID:             "c24e60c2-9dcb-466a-8fc4-fdf19c630b29",
		Type:           "accounts",
		OrganisationID: "ec88b87d-40d5-46b8-8133-98079904ce4b",
	}
	dataRequest := DataRequest{Data: data}

	Convey("When I Create an account", t, func() {
		resp, _ := service.Create(dataRequest)

		Convey("The Country should match", func() {
			So(resp.Data.Attributes.Country, ShouldEqual, "GB")
		})
	})
}
