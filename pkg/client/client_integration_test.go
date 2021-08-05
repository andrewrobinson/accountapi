package client

//can use client_test here, but lose ability to get to internals for delete

import (
	"flag"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/andrewrobinson/accountapi/pkg/client/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	uuid "github.com/satori/go.uuid"
)

var endpointFlag = flag.String("endpoint", "http://localhost:8080/v1/organisation/accounts", "")

func TestClientIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Integration Suite")
}

func initHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			MaxIdleConns:        5,
			MaxIdleConnsPerHost: 1,
		},
	}
}

var _ = Describe("Client Integration", func() {

	fmt.Printf("Client Integration tests running against endpoint:%s\n", *endpointFlag)

	accountClient := NewAccountRestClient(*endpointFlag, initHTTPClient())

	id := uuid.FromStringOrNil("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	BeforeEach(func() {

		//cleardown via internal method that doesn't care about statusCode
		_, _, err := accountClient.deleteInternal(id, 0)
		Expect(err).To(BeNil())

	})

	Context("CRUD for a single insert", func() {
		It("Behaves as expected", func() {

			//insert and assert about what is returned
			dataToInsert := buildAccountDataForCreate(id)
			accountData, err := accountClient.Create(dataToInsert)
			Expect(err).To(BeNil())
			Expect(accountData.Data.ID).To(Equal(id))
			Expect(*accountData.Data.Version).To(Equal(int64(0)))
			Expect(accountData.Data.Attributes).To(Equal(dataToInsert.Data.Attributes))

			//fetch and expect back what was inserted, along with an extra Link section
			fetchedAccountData, err := accountClient.Fetch(id)
			Expect(err).To(BeNil())
			Expect(fetchedAccountData.Data.ID).To(Equal(id))
			Expect(*fetchedAccountData.Data.Version).To(Equal(int64(0)))
			Expect(fetchedAccountData.Data.Attributes).To(Equal(dataToInsert.Data.Attributes))
			expectedLink := fmt.Sprintf("/v1/organisation/accounts/%s", id)
			Expect(*fetchedAccountData.Links.Self).To(Equal(expectedLink))

			//Delete and get no error (due to a 204)
			err = accountClient.Delete(id, 0)
			Expect(err).To(BeNil())

			//fetch and it is not found
			_, err = accountClient.Fetch(id)
			Expect(err.Error()).To(Equal("NOT FOUND"))

			//Delete and get an error (due to a 404)
			err = accountClient.Delete(id, 0)
			Expect(err.Error()).To(Equal("Delete statusCode:404 instead of 204"))

		})
	})

})

func buildAccountDataForCreate(id uuid.UUID) model.AccountDataForCreate {

	country := "GB"
	accountClassification := "Personal"

	att := model.AccountAttributes{Name: []string{"Samantha Holder"},
		Country: &country, BaseCurrency: "GBP", BankID: "400302", BankIDCode: "GBDSC",
		AccountNumber: "10000004", Iban: "GB28NWBK40030212764204", Bic: "NWBKGB42", AccountClassification: &accountClassification,
	}

	m := model.Account{ID: id,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts", Attributes: &att}

	return model.AccountDataForCreate{Data: &m}
}
