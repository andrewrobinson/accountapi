package main

import (
	"flag"
	"fmt"

	"github.com/andrewrobinson/accountapi/pkg/client"
	"github.com/andrewrobinson/accountapi/pkg/client/model"
	uuid "github.com/satori/go.uuid"
)

func main() {

	endpointFlag := flag.String("endpoint", "http://localhost:8080/v1/organisation/accounts", "")
	flag.Parse()

	fmt.Printf("main running against endpoint:%s\n", *endpointFlag)

	accountRestClient := client.NewAccountRestClient(*endpointFlag)

	id := uuid.FromStringOrNil("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")

	fetch(accountRestClient, id)
	create(accountRestClient, id)
	fetch(accountRestClient, id)
	delete(accountRestClient, id)

}

func fetch(accountRestClient *client.AccountRestClient, id uuid.UUID) {

	fetchedAccountData, err := accountRestClient.Fetch(id)

	if err != nil {
		fmt.Printf("fetch failure, err: %+v\n", err)
	} else {
		fmt.Printf("fetch success returned attributes: %+v, links:%s with id:%s\n\n",
			fetchedAccountData.Data.Attributes, *fetchedAccountData.Links.Self, fetchedAccountData.Data.ID)
	}

}

func create(accountRestClient *client.AccountRestClient, id uuid.UUID) {

	data := buildAccountDataForCreate(id)

	accountData, err := accountRestClient.Create(data)

	if err != nil {
		fmt.Printf("create failure, err: %+v\n", err)
	} else {
		fmt.Printf("create success returned attributes: %+v with id:%s\n\n",
			accountData.Data.Attributes, accountData.Data.ID)
	}

}

func delete(accountRestClient *client.AccountRestClient, id uuid.UUID) {

	err := accountRestClient.Delete(id, 0)

	if err != nil {
		fmt.Printf("delete failure, err: %+v\n", err)
	} else {
		fmt.Println("delete success\n\n")
	}

}

func buildAccountDataForCreate(id uuid.UUID) model.AccountDataForCreate {

	country := "GB"
	accountClassification := "Personal"

	att := model.AccountAttributes{Name: []string{"Samantha Holder"},
		Country: &country, BaseCurrency: "GBP", BankID: "400302", BankIDCode: "GBDSC",
		AccountNumber: "10000004", CustomerID: "234", Iban: "GB28NWBK40030212764204", Bic: "NWBKGB42", AccountClassification: &accountClassification,
	}

	m := model.Account{ID: id,
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		Type:           "accounts", Attributes: &att}

	return model.AccountDataForCreate{Data: &m}
}
