package client

import (
	"flag"
	"fmt"
	"testing"

	"github.com/andrewrobinson/accountapi/pkg/client/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	uuid "github.com/satori/go.uuid"
)

func TestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client Integration Suite")
}

type Books struct {
	Title  string
	Author string
	Pages  int
}

var _ = Describe("Client Integration", func() {

	endpointFlag := flag.String("endpoint", "http://localhost:8080/v1/organisation/accounts", "")
	flag.Parse()

	// fmt.Printf("Client Integration tests running against endpoint:%s\n", *endpointFlag)

	accountClient := NewAccountClient(*endpointFlag)

	id := uuid.FromStringOrNil("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	BeforeEach(func() {
		//cleardown
		err := accountClient.Delete(id, 0)
		Expect(err).To(BeNil())

		//fetch and expect to find nothing.
		_, err = accountClient.Fetch(id)
		Expect(err.Error()).To(Equal("NOT FOUND"))

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

			//Delete
			err = accountClient.Delete(id, 0)
			Expect(err).To(BeNil())

			//fetch and expect to find nothing.
			_, err = accountClient.Fetch(id)
			Expect(err.Error()).To(Equal("NOT FOUND"))

		})
	})

})

func fetch(accountClient AccountClient, id uuid.UUID) {

	fetchedAccountData, err := accountClient.Fetch(id)

	if err != nil {
		fmt.Printf("fetch failure, err: %+v\n", err)
	} else {
		fmt.Printf("fetch success returned attributes: %+v, links:%s with id:%s\n\n",
			fetchedAccountData.Data.Attributes, *fetchedAccountData.Links.Self, fetchedAccountData.Data.ID)
	}

}

func create(accountRestClient AccountClient, id uuid.UUID) {

	data := buildAccountDataForCreate(id)

	accountData, err := accountRestClient.Create(data)

	if err != nil {
		fmt.Printf("create failure, err: %+v\n", err)
	} else {
		fmt.Printf("create success returned attributes: %+v with id:%s\n\n",
			accountData.Data.Attributes, accountData.Data.ID)
	}

}

func delete(accountRestClient AccountClient, id uuid.UUID) {

	err := accountRestClient.Delete(id, 0)

	if err != nil {
		fmt.Printf("delete failure, err: %+v\n", err)
	} else {
		fmt.Println("delete success")
	}

}

func buildAccountDataForCreate(id uuid.UUID) model.AccountDataForCreate {

	country := "GB"
	accountClassification := "Personal"

	att := model.AccountAttributes{Name: []string{"Samantha Holder"},
		Country: &country, BaseCurrency: "GBP", BankID: "400302", BankIDCode: "GBDSC",
		AccountNumber: "10000004", Iban: "GB28NWBK40030212764204", Bic: "NWBKGB42", AccountClassification: &accountClassification,
	}

	// att := model.AccountAttributes{Name: []string{"Samantha Holder"},
	// 	Country: &country, BaseCurrency: "GBP", BankID: "400302", BankIDCode: "GBDSC",
	// 	AccountNumber: "10000004", CustomerID: "234", Iban: "GB28NWBK40030212764204", Bic: "NWBKGB42", AccountClassification: &accountClassification,
	// }

	m := model.Account{ID: id,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts", Attributes: &att}

	return model.AccountDataForCreate{Data: &m}
}
